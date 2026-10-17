// Package diagnostics describes what a run's composite is made of and how
// sensitive it is: composition by construct, evidence grade and data age;
// alternative aggregations; a Monte Carlo uncertainty interval; and the
// methodology v3 depth/reach shadow index. Nothing here changes v2 scores.
package diagnostics

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"strings"
	"time"

	"github.com/thinkwright/delegation-curve/internal/collect"
	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// Input is one scored indicator in a run.
type Input struct {
	DomainID       string
	Indicator      string
	Construct      collect.Construct
	Grade          string
	Confidence     string
	MeasurementEnd string
	DomainWeight   float64
	Weight         float64 // within-domain weight
	Normalized     float64
}

// EffectiveWeight is the indicator's share of the composite.
func (in Input) EffectiveWeight() float64 { return in.DomainWeight * in.Weight }

// Inputs collects the scored observations of a run, attaching construct tags
// and domain weights from the current configuration.
func Inputs(seed *ingest.Seed, runID string) ([]Input, error) {
	constructs := map[string]collect.Construct{}
	domainWeights := map[string]float64{}
	for _, d := range collect.AllDomainConfigs() {
		domainWeights[d.DomainID] = d.Weight
		for _, ind := range d.Indicators {
			constructs[d.DomainID+"/"+ind.Name] = ind.Construct
		}
	}
	for _, ds := range seed.DomainScores {
		if ds.RunID == runID {
			domainWeights[ds.DomainID] = ds.Weight
		}
	}

	var inputs []Input
	for _, o := range seed.IndicatorObservations {
		if o.RunID != runID || !o.IncludedInScore {
			continue
		}
		inputs = append(inputs, Input{
			DomainID:       o.DomainID,
			Indicator:      o.IndicatorName,
			Construct:      constructs[o.DomainID+"/"+o.IndicatorName],
			Grade:          o.EvidenceGrade,
			Confidence:     o.Confidence,
			MeasurementEnd: o.MeasurementEnd,
			DomainWeight:   domainWeights[o.DomainID],
			Weight:         o.Weight,
			Normalized:     o.NormalizedValue,
		})
	}
	if len(inputs) == 0 {
		return nil, fmt.Errorf("run %q has no scored indicator observations", runID)
	}
	return inputs, nil
}

// Composite recomputes the weighted composite from inputs, keeping only
// inputs accepted by keep. Domains with no kept input drop out and the
// remaining domain weights are renormalized; coverage reports the share of
// domain weight that remains.
func Composite(inputs []Input, keep func(Input) bool) (score, coverage float64) {
	type acc struct{ num, den, w float64 }
	byDomain := map[string]*acc{}
	for _, in := range inputs {
		if keep != nil && !keep(in) {
			continue
		}
		a := byDomain[in.DomainID]
		if a == nil {
			a = &acc{w: in.DomainWeight}
			byDomain[in.DomainID] = a
		}
		a.num += in.Weight * in.Normalized
		a.den += in.Weight
	}
	var num, den float64
	for _, a := range byDomain {
		if a.den == 0 {
			continue
		}
		num += a.w * a.num / a.den
		den += a.w
	}
	if den == 0 {
		return 0, 0
	}
	return num / den, den
}

// DomainScores returns each domain's score from the kept inputs.
func DomainScores(inputs []Input, keep func(Input) bool) map[string]float64 {
	num, den := map[string]float64{}, map[string]float64{}
	for _, in := range inputs {
		if keep == nil || keep(in) {
			num[in.DomainID] += in.Weight * in.Normalized
			den[in.DomainID] += in.Weight
		}
	}
	out := map[string]float64{}
	for d, w := range den {
		if w > 0 {
			out[d] = num[d] / w
		}
	}
	return out
}

// ShadowDepthWeight is the depth share of the v3 shadow composite: each
// domain scores 0.7 x depth (share-of-work indicators) + 0.3 x reach (usage
// and adoption). Capacity metrics are excluded from the shadow.
const ShadowDepthWeight = 0.7

// Shadow is the methodology v3 depth/reach shadow computation.
type Shadow struct {
	Depth, DepthCoverage float64
	Reach, ReachCoverage float64
	// Combined uses only domains that have both depth and reach evidence.
	Combined, CombinedCoverage float64
	DomainDepth, DomainReach   map[string]float64
}

func ComputeShadow(inputs []Input) Shadow {
	isDepth := func(in Input) bool { return in.Construct == collect.Work }
	isReach := func(in Input) bool { return in.Construct == collect.Usage || in.Construct == collect.Adoption }
	s := Shadow{DomainDepth: DomainScores(inputs, isDepth), DomainReach: DomainScores(inputs, isReach)}
	s.Depth, s.DepthCoverage = Composite(inputs, isDepth)
	s.Reach, s.ReachCoverage = Composite(inputs, isReach)

	weights := map[string]float64{}
	for _, in := range inputs {
		weights[in.DomainID] = in.DomainWeight
	}
	var num, den float64
	for d, w := range weights {
		depth, okD := s.DomainDepth[d]
		reach, okR := s.DomainReach[d]
		if okD && okR {
			num += w * (ShadowDepthWeight*depth + (1-ShadowDepthWeight)*reach)
			den += w
		}
	}
	if den > 0 {
		s.Combined, s.CombinedCoverage = num/den, den
	}
	return s
}

// ConfidenceHalfWidth maps a confidence label to the half-width, in
// normalized points, of a 90% plausible range around the indicator value.
// These are explicit modelling assumptions, documented with the output.
var ConfidenceHalfWidth = map[string]float64{
	"high":        5,
	"medium-high": 7.5,
	"medium":      10,
	"medium-low":  12.5,
	"low-medium":  12.5,
	"low":         15,
}

const (
	defaultHalfWidth = 15   // unlabelled inputs are treated as low confidence
	weightLogSD      = 0.25 // lognormal domain-weight perturbation
	z90              = 1.6449
)

// Interval is a Monte Carlo uncertainty summary.
type Interval struct {
	P5, P50, P95 float64
	Draws        int
}

// Uncertainty perturbs every indicator by its confidence range (normal,
// 90% half-width per ConfidenceHalfWidth, clamped to 0-100) and every
// domain weight by lognormal noise, then reports composite percentiles.
// The seed makes results reproducible.
func Uncertainty(inputs []Input, draws int, seed uint64) Interval {
	rng := rand.New(rand.NewPCG(seed, seed^0x9e3779b97f4a7c15))
	domains := map[string]float64{}
	for _, in := range inputs {
		domains[in.DomainID] = in.DomainWeight
	}
	ids := make([]string, 0, len(domains))
	for d := range domains {
		ids = append(ids, d)
	}
	sort.Strings(ids)

	results := make([]float64, draws)
	perturbed := make([]Input, len(inputs))
	for k := 0; k < draws; k++ {
		w := map[string]float64{}
		var total float64
		for _, d := range ids {
			w[d] = domains[d] * math.Exp(rng.NormFloat64()*weightLogSD)
			total += w[d]
		}
		for i, in := range inputs {
			hw, ok := ConfidenceHalfWidth[in.Confidence]
			if !ok {
				hw = defaultHalfWidth
			}
			in.Normalized = math.Max(0, math.Min(100, in.Normalized+rng.NormFloat64()*hw/z90))
			in.DomainWeight = w[in.DomainID] / total
			perturbed[i] = in
		}
		results[k], _ = Composite(perturbed, nil)
	}
	sort.Float64s(results)
	pct := func(p float64) float64 { return results[int(p*float64(draws-1))] }
	return Interval{P5: pct(0.05), P50: pct(0.5), P95: pct(0.95), Draws: draws}
}

// AgeOf returns the data age at asOf, reading reduced-precision dates from
// the start of the period they name. ok is false when the date is missing.
func AgeOf(measurementEnd string, asOf time.Time) (time.Duration, bool) {
	for _, layout := range []string{"2006-01-02", "2006-01", "2006"} {
		if t, err := time.Parse(layout, measurementEnd); err == nil {
			return asOf.Sub(t), true
		}
	}
	return 0, false
}

// Report is the full diagnostics for one run.
type Report struct {
	RunID       string
	AsOf        time.Time
	Published   float64
	Inputs      []Input
	ByConstruct map[collect.Construct][2]float64 // weight share, composite points
	ByGrade     map[string]float64
	// Evidence age, weighted by effective weight.
	MedianAgeMonths   float64
	ShareOlder12m     float64
	ShareOlder24m     float64
	ShareUndated      float64
	EqualWeights      float64
	GradeAB, GradeABC float64
	GradeABCoverage   float64
	NoCapacity        float64
	Shadow            Shadow
	Interval          Interval
}

// Analyze builds the report. asOf is the date ages are measured at (the
// run's publication date).
func Analyze(runID string, inputs []Input, published float64, asOf time.Time) Report {
	r := Report{
		RunID: runID, AsOf: asOf, Published: published, Inputs: inputs,
		ByConstruct: map[collect.Construct][2]float64{}, ByGrade: map[string]float64{},
	}

	type aged struct{ months, w float64 }
	var ages []aged
	var undated float64
	for _, in := range inputs {
		ew := in.EffectiveWeight()
		c := r.ByConstruct[in.Construct]
		c[0] += ew
		c[1] += ew * in.Normalized
		r.ByConstruct[in.Construct] = c
		grade := in.Grade
		if grade == "" {
			grade = "ungraded"
		}
		r.ByGrade[grade] += ew
		if age, ok := AgeOf(in.MeasurementEnd, asOf); ok {
			months := age.Hours() / 24 / 30.4375
			ages = append(ages, aged{months, ew})
			if months > 12 {
				r.ShareOlder12m += ew
			}
			if months > 24 {
				r.ShareOlder24m += ew
			}
		} else {
			undated += ew
		}
	}
	r.ShareUndated = undated
	sort.Slice(ages, func(i, j int) bool { return ages[i].months < ages[j].months })
	var dated, cum float64
	for _, a := range ages {
		dated += a.w
	}
	for _, a := range ages {
		cum += a.w
		if cum >= dated/2 {
			r.MedianAgeMonths = a.months
			break
		}
	}

	equal := make([]Input, len(inputs))
	domains := map[string]bool{}
	for _, in := range inputs {
		domains[in.DomainID] = true
	}
	for i, in := range inputs {
		in.DomainWeight = 1 / float64(len(domains))
		equal[i] = in
	}
	r.EqualWeights, _ = Composite(equal, nil)
	r.GradeAB, r.GradeABCoverage = Composite(inputs, func(in Input) bool { return in.Grade == "A" || in.Grade == "B" })
	r.GradeABC, _ = Composite(inputs, func(in Input) bool { return in.Grade != "D" })
	r.NoCapacity, _ = Composite(inputs, func(in Input) bool { return in.Construct != collect.Capacity })
	r.Shadow = ComputeShadow(inputs)
	r.Interval = Uncertainty(inputs, 20000, 20261017)
	return r
}

// Markdown renders the report for run records and evidence notes.
func (r Report) Markdown() string {
	var b strings.Builder
	pct := func(v float64) string { return fmt.Sprintf("%.1f%%", v*100) }
	fmt.Fprintf(&b, "## Diagnostics: run `%s`\n\n", r.RunID)
	fmt.Fprintf(&b, "Published composite: **%.1f**. Ages are measured at %s. Generated by `curve-diagnostics`; nothing here changes the published score.\n\n",
		r.Published, r.AsOf.Format("2006-01-02"))

	b.WriteString("### What the composite measures\n\n| Construct | Composite weight | Points |\n| --- | ---: | ---: |\n")
	labels := map[collect.Construct]string{
		collect.Work:     "Share of work or decisions handled by AI",
		collect.Usage:    "Share of people using AI",
		collect.Adoption: "Share of organizations using AI",
		collect.Capacity: "Capacity or count metrics",
	}
	for _, c := range []collect.Construct{collect.Work, collect.Usage, collect.Adoption, collect.Capacity} {
		v := r.ByConstruct[c]
		fmt.Fprintf(&b, "| %s | %s | %.1f |\n", labels[c], pct(v[0]), v[1])
	}

	b.WriteString("\n### Evidence grade\n\n| Grade | Composite weight |\n| --- | ---: |\n")
	var grades []string
	for g := range r.ByGrade {
		grades = append(grades, g)
	}
	sort.Strings(grades)
	for _, g := range grades {
		fmt.Fprintf(&b, "| %s | %s |\n", g, pct(r.ByGrade[g]))
	}

	b.WriteString("\n### Evidence age\n\n")
	fmt.Fprintf(&b, "- Weighted median age of the data at publication: **%.0f months** (from each input's measurement end).\n", r.MedianAgeMonths)
	fmt.Fprintf(&b, "- Weight resting on data more than 12 months old: %s; more than 24 months: %s.\n", pct(r.ShareOlder12m), pct(r.ShareOlder24m))
	if r.ShareUndated > 0 {
		fmt.Fprintf(&b, "- Weight with no recorded measurement window: %s.\n", pct(r.ShareUndated))
	}

	b.WriteString("\n### Sensitivity\n\n| Aggregation | Composite |\n| --- | ---: |\n")
	fmt.Fprintf(&b, "| Published (current weights) | %.1f |\n", r.Published)
	fmt.Fprintf(&b, "| Equal domain weights | %.1f |\n", r.EqualWeights)
	fmt.Fprintf(&b, "| Grade A/B evidence only (%s of domain weight covered) | %.1f |\n", pct(r.GradeABCoverage), r.GradeAB)
	fmt.Fprintf(&b, "| Excluding grade D | %.1f |\n", r.GradeABC)
	fmt.Fprintf(&b, "| Excluding capacity/count metrics | %.1f |\n", r.NoCapacity)

	fmt.Fprintf(&b, "\n### Uncertainty\n\n90%% interval: **%.1f to %.1f** (median %.1f, %d draws). ", r.Interval.P5, r.Interval.P95, r.Interval.P50, r.Interval.Draws)
	b.WriteString("Each indicator is perturbed within a confidence-based 90% range (high ±5, medium-high ±7.5, medium ±10, low-medium ±12.5, low ±15 normalized points) and domain weights by lognormal noise (σ = 0.25). The ranges are modelling assumptions, not sampling error.\n")

	s := r.Shadow
	b.WriteString("\n### Methodology v3 shadow: depth and reach\n\n")
	fmt.Fprintf(&b, "- **Delegation depth** (share-of-work indicators only): %.1f, covering %s of domain weight.\n", s.Depth, pct(s.DepthCoverage))
	fmt.Fprintf(&b, "- **Reach** (people and organizations using AI): %.1f, covering %s.\n", s.Reach, pct(s.ReachCoverage))
	fmt.Fprintf(&b, "- **Shadow composite** (%.0f%% depth / %.0f%% reach, domains with both): %.1f, covering %s.\n\n",
		ShadowDepthWeight*100, (1-ShadowDepthWeight)*100, s.Combined, pct(s.CombinedCoverage))
	b.WriteString("| Domain | Depth | Reach |\n| --- | ---: | ---: |\n")
	var ids []string
	seen := map[string]bool{}
	for _, in := range r.Inputs {
		if !seen[in.DomainID] {
			seen[in.DomainID] = true
			ids = append(ids, in.DomainID)
		}
	}
	cell := func(m map[string]float64, d string) string {
		if v, ok := m[d]; ok {
			return fmt.Sprintf("%.1f", v)
		}
		return "—"
	}
	for _, d := range ids {
		fmt.Fprintf(&b, "| %s | %s | %s |\n", d, cell(s.DomainDepth, d), cell(s.DomainReach, d))
	}
	return b.String()
}
