package collect

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// RunConfig holds paths and settings for a collection run.
type RunConfig struct {
	SeedPath      string
	OverridesPath string
	LogPath       string
	DomainConfigs []DomainConfig
	Collectors    []Collector
	Timeout       time.Duration
	// CompositeConfigs weights the composite; nil means AllDomainConfigs, so
	// partial runs (-domain) still average over every domain.
	CompositeConfigs []DomainConfig
	DryRun           bool // compute and report changes without writing the seed or log
	Check            bool // with DryRun: fail when any score would change
}

// ErrScoresChanged is returned by a checked dry run whose recomputed scores
// differ from the seed.
var ErrScoresChanged = errors.New("recomputed scores differ from seed")

// Run executes the collection pipeline for all configured domains.
//
// Collection only updates current values: sub-indicator values, domain
// scores and statuses, and the composite score. Run history (previous
// scores, deltas, trends) belongs to explicit analysis runs and is written
// by curve-snapshot-run, never here.
func Run(cfg RunConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	seed, err := ingest.ReadSeed(cfg.SeedPath)
	if err != nil {
		return fmt.Errorf("read seed: %w", err)
	}

	overrides, err := LoadOverrides(cfg.OverridesPath)
	if err != nil {
		return fmt.Errorf("load overrides: %w", err)
	}

	// C3: Check for stale override data before proceeding.
	if warnings := CheckOverrideStaleness(overrides); len(warnings) > 0 {
		fmt.Fprintln(os.Stderr, "\n  Staleness warnings:")
		for _, w := range warnings {
			fmt.Fprintf(os.Stderr, "    ⚠ %s\n", w)
		}
		fmt.Fprintln(os.Stderr)
	}

	var before ingest.Seed
	if err := ingest.Clone(seed, &before); err != nil {
		return err
	}

	// Run all collectors.
	//
	// Error convention (Q6): A collector returns (results, nil) with Err set on
	// individual CollectResult entries for per-indicator failures (e.g., manual
	// stubs). A collector returns (nil, err) only for infrastructure-level
	// failures that prevent any collection attempt (e.g., context canceled).
	var allResults []CollectResult
	for _, c := range cfg.Collectors {
		fmt.Fprintf(os.Stderr, "  Collecting from %s...\n", c.Name())
		results, err := c.Collect(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "    WARN: %s failed: %v\n", c.Name(), err)
			continue
		}
		now := time.Now()
		for i := range results {
			results[i].CollectedAt = now
			if results[i].Err != nil {
				fmt.Fprintf(os.Stderr, "    MANUAL: %s → %v\n", results[i].IndicatorName, results[i].Err)
			} else {
				fmt.Fprintf(os.Stderr, "    OK: %s = %.2f %s\n", results[i].IndicatorName, results[i].RawValue, results[i].Unit)
			}
		}
		allResults = append(allResults, results...)
	}

	// Update each configured domain.
	newLog := &CollectLog{
		RunAt:   time.Now(),
		Entries: map[string][]LogEntry{},
	}

	for _, dcfg := range cfg.DomainConfigs {
		if err := updateDomain(seed, dcfg, allResults, overrides, newLog); err != nil {
			return fmt.Errorf("update domain %s: %w", dcfg.DomainID, err)
		}
	}

	compositeConfigs := cfg.CompositeConfigs
	if compositeConfigs == nil {
		compositeConfigs = AllDomainConfigs()
	}
	updateComposite(seed, compositeConfigs)

	changes := DiffScores(&before, seed)
	reportChanges(changes)
	if cfg.DryRun {
		fmt.Fprintln(os.Stderr, "\n  Dry run: seed and collect log not written.")
		if cfg.Check && len(changes) > 0 {
			return ErrScoresChanged
		}
		return nil
	}

	// W5: Backup seed.json before overwriting.
	backupPath := cfg.SeedPath + ".bak"
	if data, err := os.ReadFile(cfg.SeedPath); err == nil {
		if err := os.WriteFile(backupPath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "  WARN: could not create backup at %s: %v\n", backupPath, err)
		}
	}

	if err := ingest.WriteSeed(cfg.SeedPath, seed); err != nil {
		return fmt.Errorf("write seed: %w", err)
	}

	if err := newLog.Save(cfg.LogPath); err != nil {
		return fmt.Errorf("write collect log: %w", err)
	}

	return nil
}

func updateDomain(
	seed *ingest.Seed,
	dcfg DomainConfig,
	results []CollectResult,
	overrides OverrideFile,
	newLog *CollectLog,
) error {
	// Find the domain in seed.
	var domain *ingest.DomainJSON
	for i := range seed.Delegation.Domains {
		if seed.Delegation.Domains[i].ID == dcfg.DomainID {
			domain = &seed.Delegation.Domains[i]
			break
		}
	}
	if domain == nil {
		return fmt.Errorf("domain %s not found in seed", dcfg.DomainID)
	}

	// W2: Only include indicators that have a real value in the scoring map.
	// Missing indicators are excluded from the weighted average, not scored as zero.
	normalizedValues := map[string]float64{}

	for _, indCfg := range dcfg.Indicators {
		r := resolveValue(indCfg.Name, dcfg.DomainID, results, overrides, domain.SubIndicators)

		// W1: Loud warning for cached/missing fallbacks.
		switch r.Method {
		case "cached":
			fmt.Fprintf(os.Stderr, "    ⚠ CACHED: %s — using stale value %.2f %s (no collector or override)\n",
				indCfg.Name, r.Value, r.Unit)
		case "missing":
			fmt.Fprintf(os.Stderr, "    ✗ MISSING: %s — no value from any source, excluded from scoring\n",
				indCfg.Name)
		}

		// Only include in scoring if we actually have a value.
		if r.Method != "missing" {
			normalized := Normalize(r.Value, indCfg.NormConfig)
			normalizedValues[indCfg.Name] = normalized
			updateSubIndicator(domain, indCfg.Name, r)
			fmt.Fprintf(os.Stderr, "    %s: %.2f %s [%s] → normalized %.1f\n",
				indCfg.Name, r.Value, r.Unit, r.Method, normalized)
		}

		newLog.Entries[dcfg.DomainID] = append(newLog.Entries[dcfg.DomainID], LogEntry{
			IndicatorName: indCfg.Name,
			SourceName:    r.Source,
			Value:         r.Value,
			Unit:          r.Unit,
			Freshness:     r.Freshness,
			CollectedAt:   time.Now(),
			Method:        r.Method,
			Error:         r.Err,
			SourceURL:     r.Meta.SourceURL,
		})
	}

	newScore := math.Round(ComputeDomainScore(normalizedValues, dcfg)*10) / 10

	fmt.Fprintf(os.Stderr, "  %s score: %.1f → %.1f\n", dcfg.DomainID, domain.Score, newScore)

	filterSubIndicatorsToConfig(domain, dcfg.Indicators)

	domain.Score = newScore
	domain.Status = ClassifyStatus(newScore)

	return nil
}

// resolved is one indicator value and where it came from.
type resolved struct {
	Value     float64
	Unit      string
	Source    string
	Freshness string
	Meta      ingest.SourceMeta
	Method    string // "auto", "override", "cached", or "missing"
	Err       string
}

// resolveValue implements three-tier fallback: collected → override → cached.
// Returns method "missing" if no value is available from any source.
func resolveValue(
	name, domainID string,
	results []CollectResult,
	overrides OverrideFile,
	existingSubs []ingest.SubIndicatorJSON,
) resolved {
	// 1. Try collected results (auto).
	for _, r := range results {
		if r.IndicatorName == name && r.DomainID == domainID && r.Err == nil {
			return resolved{r.RawValue, r.Unit, r.SourceName, r.Freshness, r.Meta, "auto", ""}
		}
	}

	// 2. Try overrides (manual YAML).
	for _, e := range overrides[domainID] {
		if e.Name == name {
			return resolved{e.Value, e.Unit, e.Source, e.Freshness, e.SourceMeta, "override", ""}
		}
	}

	// 3. Fall back to existing seed value (cached — last known good).
	for _, si := range existingSubs {
		if si.Name == name {
			return resolved{si.Value, si.Unit, si.Source, si.Freshness, si.SourceMeta, "cached", "no collector or override available"}
		}
	}

	return resolved{Method: "missing", Err: "no value found from any source"}
}

// updateComposite recomputes the composite delegation score as the weighted
// average of domain scores. Domains in the seed but absent from configs are
// skipped, and the average is taken over the weight that is present.
func updateComposite(seed *ingest.Seed, configs []DomainConfig) {
	weightMap := make(map[string]float64, len(configs))
	for _, c := range configs {
		weightMap[c.DomainID] = c.Weight
	}

	var weighted, totalWeight float64
	for _, d := range seed.Delegation.Domains {
		w, ok := weightMap[d.ID]
		if !ok {
			continue
		}
		weighted += d.Score * w
		totalWeight += w
	}
	if totalWeight == 0 {
		return
	}
	composite := math.Round(weighted/totalWeight*10) / 10

	prev := seed.Delegation.Composite.Current
	seed.Delegation.Composite.Current = composite

	fmt.Fprintf(os.Stderr, "\n  Composite score: %.1f → %.1f\n", prev, composite)
}

// updateSubIndicator updates an existing sub-indicator in the domain, or appends it.
func updateSubIndicator(domain *ingest.DomainJSON, name string, r resolved) {
	si := ingest.SubIndicatorJSON{
		Name:       name,
		Value:      r.Value,
		Unit:       r.Unit,
		Source:     r.Source,
		Freshness:  r.Freshness,
		SourceMeta: r.Meta,
	}
	for i := range domain.SubIndicators {
		if domain.SubIndicators[i].Name == name {
			domain.SubIndicators[i] = si
			return
		}
	}
	domain.SubIndicators = append(domain.SubIndicators, si)
}

func filterSubIndicatorsToConfig(domain *ingest.DomainJSON, indicators []IndicatorConfig) {
	byName := make(map[string]ingest.SubIndicatorJSON, len(domain.SubIndicators))
	for _, si := range domain.SubIndicators {
		byName[si.Name] = si
	}

	filtered := make([]ingest.SubIndicatorJSON, 0, len(indicators))
	for _, ind := range indicators {
		if si, ok := byName[ind.Name]; ok {
			filtered = append(filtered, si)
		}
	}
	domain.SubIndicators = filtered
}
