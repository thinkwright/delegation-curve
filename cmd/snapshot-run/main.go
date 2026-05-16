package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/thinkwright/delegation-curve/internal/collect"
	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/schema"
	"github.com/thinkwright/delegation-curve/internal/transform"
)

func main() {
	seedPath := flag.String("seed", "seed/seed.json", "Path to seed.json")
	runID := flag.String("run-id", "", "Run identifier to append from the current seed scores")
	label := flag.String("label", "", "Human-readable run label")
	publishedAt := flag.String("published-at", "", "Publication date for the run")
	measurementPeriod := flag.String("measurement-period", "", "Measurement period label")
	measurementYear := flag.Int("measurement-year", 0, "Measurement year")
	methodologyVersion := flag.String("methodology-version", "delegation-curve-v1", "Methodology version")
	notes := flag.String("notes", "", "Run notes")
	dataFreshness := flag.String("data-freshness", "", "Composite data freshness label to write")
	current := flag.Bool("current", true, "Mark appended run as the current run")
	bootstrap := flag.Bool("bootstrap", true, "Materialize legacy trend history first when no explicit runs exist")
	bootstrapOnly := flag.Bool("bootstrap-only", false, "Only materialize legacy trend history")
	replace := flag.Bool("replace", false, "Replace an existing run with the same run ID")
	flag.Parse()

	seed, err := ingest.ReadSeed(*seedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading seed: %v\n", err)
		os.Exit(1)
	}

	changed := false
	if *dataFreshness != "" {
		seed.Delegation.Composite.DataFreshness = *dataFreshness
		changed = true
	}

	if *bootstrap && len(seed.AnalysisRuns) == 0 {
		materializeLegacyHistory(seed)
		changed = true
		fmt.Fprintf(os.Stderr, "Materialized %d legacy analysis runs.\n", len(seed.AnalysisRuns))
	}

	if !*bootstrapOnly {
		if *runID == "" {
			fmt.Fprintln(os.Stderr, "Error: -run-id is required unless -bootstrap-only is set")
			os.Exit(1)
		}
		if hasRun(seed, *runID) && !*replace {
			fmt.Fprintf(os.Stderr, "Error: run %q already exists; use -replace to update it\n", *runID)
			os.Exit(1)
		}
		if *replace {
			removeRun(seed, *runID)
		}

		appendCurrentSnapshot(seed, snapshotOptions{
			runID:              *runID,
			label:              *label,
			publishedAt:        *publishedAt,
			measurementPeriod:  *measurementPeriod,
			measurementYear:    *measurementYear,
			methodologyVersion: *methodologyVersion,
			notes:              *notes,
			current:            *current,
		})
		changed = true
	}

	if !changed {
		fmt.Fprintln(os.Stderr, "No changes.")
		return
	}

	sortExplicitHistory(seed)
	if err := ingest.WriteSeed(*seedPath, seed); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing seed: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Updated %s.\n", *seedPath)
}

type snapshotOptions struct {
	runID              string
	label              string
	publishedAt        string
	measurementPeriod  string
	measurementYear    int
	methodologyVersion string
	notes              string
	current            bool
}

func materializeLegacyHistory(seed *ingest.Seed) {
	runs, domains, observations := transform.History(seed)
	seed.AnalysisRuns = toAnalysisRunJSON(runs)
	seed.DomainScores = toDomainScoreJSON(domains)
	seed.IndicatorObservations = toIndicatorObservationJSON(observations)
}

func appendCurrentSnapshot(seed *ingest.Seed, opts snapshotOptions) {
	composite := seed.Delegation.Composite
	if opts.measurementYear == 0 {
		opts.measurementYear = composite.DataYear
	}
	if opts.measurementPeriod == "" {
		opts.measurementPeriod = fmt.Sprintf("%d", opts.measurementYear)
	}
	if opts.label == "" {
		opts.label = opts.measurementPeriod
	}
	if opts.publishedAt == "" {
		opts.publishedAt = composite.LastUpdated
	}

	if opts.current {
		for i := range seed.AnalysisRuns {
			seed.AnalysisRuns[i].IsCurrent = false
		}
	}

	seed.AnalysisRuns = append(seed.AnalysisRuns, ingest.AnalysisRunJSON{
		RunID:              opts.runID,
		Label:              opts.label,
		PublishedAt:        opts.publishedAt,
		MeasurementPeriod:  opts.measurementPeriod,
		MeasurementYear:    opts.measurementYear,
		MethodologyVersion: opts.methodologyVersion,
		CompositeScore:     composite.Current,
		Notes:              opts.notes,
		IsCurrent:          opts.current,
	})

	for _, d := range seed.Delegation.Domains {
		status := d.Status
		if status == "" {
			status = collect.ClassifyStatus(d.Score)
		}
		seed.DomainScores = append(seed.DomainScores, ingest.DomainScoreJSON{
			RunID:    opts.runID,
			DomainID: d.ID,
			Score:    d.Score,
			Weight:   d.Weight,
			Status:   status,
		})
	}

	observations := transform.CurrentIndicatorObservations(seed, opts.runID)
	seed.IndicatorObservations = append(seed.IndicatorObservations, toIndicatorObservationJSON(observations)...)
}

func hasRun(seed *ingest.Seed, runID string) bool {
	for _, r := range seed.AnalysisRuns {
		if r.RunID == runID {
			return true
		}
	}
	return false
}

func removeRun(seed *ingest.Seed, runID string) {
	filteredRuns := seed.AnalysisRuns[:0]
	for _, r := range seed.AnalysisRuns {
		if r.RunID != runID {
			filteredRuns = append(filteredRuns, r)
		}
	}
	seed.AnalysisRuns = filteredRuns

	filteredDomains := seed.DomainScores[:0]
	for _, d := range seed.DomainScores {
		if d.RunID != runID {
			filteredDomains = append(filteredDomains, d)
		}
	}
	seed.DomainScores = filteredDomains

	filteredObservations := seed.IndicatorObservations[:0]
	for _, o := range seed.IndicatorObservations {
		if o.RunID != runID {
			filteredObservations = append(filteredObservations, o)
		}
	}
	seed.IndicatorObservations = filteredObservations
}

func sortExplicitHistory(seed *ingest.Seed) {
	sort.SliceStable(seed.AnalysisRuns, func(i, j int) bool {
		if seed.AnalysisRuns[i].MeasurementYear != seed.AnalysisRuns[j].MeasurementYear {
			return seed.AnalysisRuns[i].MeasurementYear < seed.AnalysisRuns[j].MeasurementYear
		}
		if seed.AnalysisRuns[i].PublishedAt != seed.AnalysisRuns[j].PublishedAt {
			return seed.AnalysisRuns[i].PublishedAt < seed.AnalysisRuns[j].PublishedAt
		}
		return seed.AnalysisRuns[i].RunID < seed.AnalysisRuns[j].RunID
	})
	sort.SliceStable(seed.DomainScores, func(i, j int) bool {
		if seed.DomainScores[i].RunID != seed.DomainScores[j].RunID {
			return seed.DomainScores[i].RunID < seed.DomainScores[j].RunID
		}
		return seed.DomainScores[i].DomainID < seed.DomainScores[j].DomainID
	})
	sort.SliceStable(seed.IndicatorObservations, func(i, j int) bool {
		if seed.IndicatorObservations[i].RunID != seed.IndicatorObservations[j].RunID {
			return seed.IndicatorObservations[i].RunID < seed.IndicatorObservations[j].RunID
		}
		if seed.IndicatorObservations[i].DomainID != seed.IndicatorObservations[j].DomainID {
			return seed.IndicatorObservations[i].DomainID < seed.IndicatorObservations[j].DomainID
		}
		return seed.IndicatorObservations[i].IndicatorName < seed.IndicatorObservations[j].IndicatorName
	})
}

func toAnalysisRunJSON(rows []schema.AnalysisRunRow) []ingest.AnalysisRunJSON {
	out := make([]ingest.AnalysisRunJSON, 0, len(rows))
	for _, r := range rows {
		out = append(out, ingest.AnalysisRunJSON{
			RunID:              r.RunID,
			Label:              r.Label,
			PublishedAt:        r.PublishedAt,
			MeasurementPeriod:  r.MeasurementPeriod,
			MeasurementYear:    int(r.MeasurementYear),
			MethodologyVersion: r.MethodologyVersion,
			CompositeScore:     r.CompositeScore,
			Notes:              r.Notes,
			IsCurrent:          r.IsCurrent,
		})
	}
	return out
}

func toDomainScoreJSON(rows []schema.DomainScoreRow) []ingest.DomainScoreJSON {
	out := make([]ingest.DomainScoreJSON, 0, len(rows))
	for _, r := range rows {
		out = append(out, ingest.DomainScoreJSON{
			RunID:    r.RunID,
			DomainID: r.DomainID,
			Score:    r.Score,
			Weight:   r.Weight,
			Status:   r.Status,
		})
	}
	return out
}

func toIndicatorObservationJSON(rows []schema.IndicatorObservationRow) []ingest.IndicatorObservationJSON {
	out := make([]ingest.IndicatorObservationJSON, 0, len(rows))
	for _, r := range rows {
		out = append(out, ingest.IndicatorObservationJSON{
			RunID:           r.RunID,
			DomainID:        r.DomainID,
			IndicatorName:   r.IndicatorName,
			RawValue:        r.RawValue,
			NormalizedValue: r.NormalizedValue,
			Unit:            r.Unit,
			Weight:          r.Weight,
			IncludedInScore: r.IncludedInScore,
			Source:          r.Source,
			SourceURL:       r.SourceURL,
			Freshness:       r.Freshness,
			EvidenceGrade:   r.EvidenceGrade,
			Confidence:      r.Confidence,
		})
	}
	return out
}
