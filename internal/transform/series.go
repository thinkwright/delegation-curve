package transform

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// scoreTolerance absorbs float noise when comparing one-decimal scores.
const scoreTolerance = 0.05

// Series is the public curve described by explicit analysis runs: the
// public-series runs in curve order, the current run, and the public run
// immediately before it.
type Series struct {
	Public   []ingest.AnalysisRunJSON
	Current  *ingest.AnalysisRunJSON
	Previous *ingest.AnalysisRunJSON

	domainScores map[string]map[string]float64 // run ID -> domain ID -> score
}

// PublicSeries derives the public curve from seed.AnalysisRuns. It returns
// ok=false when the seed has no explicit runs (legacy trend-array seeds).
func PublicSeries(seed *ingest.Seed) (Series, bool, error) {
	if len(seed.AnalysisRuns) == 0 {
		return Series{}, false, nil
	}

	var public []ingest.AnalysisRunJSON
	for _, r := range seed.AnalysisRuns {
		if r.IsPublicSeries == nil || *r.IsPublicSeries {
			public = append(public, r)
		}
	}
	sortRunsForCurve(public)

	s := Series{Public: public, domainScores: map[string]map[string]float64{}}
	for _, d := range seed.DomainScores {
		if s.domainScores[d.RunID] == nil {
			s.domainScores[d.RunID] = map[string]float64{}
		}
		s.domainScores[d.RunID][d.DomainID] = d.Score
	}

	currentIdx := -1
	for i := range public {
		if public[i].IsCurrent {
			if currentIdx >= 0 {
				return Series{}, true, fmt.Errorf("runs %q and %q are both marked current", public[currentIdx].RunID, public[i].RunID)
			}
			currentIdx = i
		}
	}
	if currentIdx < 0 {
		return Series{}, true, fmt.Errorf("no public-series run is marked current")
	}
	s.Current = &s.Public[currentIdx]
	if currentIdx > 0 {
		s.Previous = &s.Public[currentIdx-1]
	}
	// The curve ends at the current run; later public runs (drafts) are not shown.
	s.Public = s.Public[:currentIdx+1]
	return s, true, nil
}

// DomainScore returns a domain's score in the given run.
func (s Series) DomainScore(runID, domainID string) (float64, bool) {
	v, ok := s.domainScores[runID][domainID]
	return v, ok
}

// DeriveHistory rewrites the seed's history-shaped fields (composite
// previous/delta/trend and each domain's previous_score/trend) from the
// explicit public series, so every consumer of seed.json sees one history.
// Seeds without explicit runs are left untouched.
func DeriveHistory(seed *ingest.Seed) error {
	s, ok, err := PublicSeries(seed)
	if err != nil || !ok {
		return err
	}

	c := &seed.Delegation.Composite
	c.Trend = make([]float64, 0, len(s.Public))
	for _, r := range s.Public {
		c.Trend = append(c.Trend, r.CompositeScore)
	}
	c.Previous = s.Current.CompositeScore
	if s.Previous != nil {
		c.Previous = s.Previous.CompositeScore
	}
	c.Delta = round1(s.Current.CompositeScore - c.Previous)

	for i := range seed.Delegation.Domains {
		d := &seed.Delegation.Domains[i]
		d.Trend = make([]float64, 0, len(s.Public))
		for _, r := range s.Public {
			if v, ok := s.DomainScore(r.RunID, d.ID); ok {
				d.Trend = append(d.Trend, v)
			}
		}
		d.PreviousScore = d.Score
		if s.Previous != nil {
			if v, ok := s.DomainScore(s.Previous.RunID, d.ID); ok {
				d.PreviousScore = v
			}
		}
	}
	return nil
}

// CheckSnapshotted reports an error when the seed's live scores differ from
// the current run's snapshot, i.e. scores were recomputed (by collect) but no
// run was recorded for them. Publishing such a seed would show numbers that
// are not part of the public series.
func CheckSnapshotted(seed *ingest.Seed) error {
	s, ok, err := PublicSeries(seed)
	if err != nil || !ok {
		return err
	}

	var diffs []string
	if !close1(seed.Delegation.Composite.Current, s.Current.CompositeScore) {
		diffs = append(diffs, fmt.Sprintf("composite %.1f (run %s: %.1f)",
			seed.Delegation.Composite.Current, s.Current.RunID, s.Current.CompositeScore))
	}
	for _, d := range seed.Delegation.Domains {
		v, ok := s.DomainScore(s.Current.RunID, d.ID)
		switch {
		case !ok:
			diffs = append(diffs, fmt.Sprintf("%s has no score in run %s", d.ID, s.Current.RunID))
		case !close1(d.Score, v):
			diffs = append(diffs, fmt.Sprintf("%s %.1f (run %s: %.1f)", d.ID, d.Score, s.Current.RunID, v))
		}
	}
	if len(diffs) > 0 {
		return fmt.Errorf("seed scores are not snapshotted in current run %s: %s",
			s.Current.RunID, strings.Join(diffs, "; "))
	}
	return nil
}

// CheckDerived reports an error when the seed's stored history fields differ
// from what DeriveHistory would write.
func CheckDerived(seed *ingest.Seed) error {
	var derived ingest.Seed
	if err := ingest.Clone(seed, &derived); err != nil {
		return err
	}
	if err := DeriveHistory(&derived); err != nil {
		return err
	}

	var diffs []string
	got, want := seed.Delegation.Composite, derived.Delegation.Composite
	if !close1(got.Previous, want.Previous) || !close1(got.Delta, want.Delta) || !sameSeries(got.Trend, want.Trend) {
		diffs = append(diffs, fmt.Sprintf("composite previous/delta/trend %.1f/%.1f/%v, want %.1f/%.1f/%v",
			got.Previous, got.Delta, got.Trend, want.Previous, want.Delta, want.Trend))
	}
	for i, d := range seed.Delegation.Domains {
		w := derived.Delegation.Domains[i]
		if !close1(d.PreviousScore, w.PreviousScore) || !sameSeries(d.Trend, w.Trend) {
			diffs = append(diffs, fmt.Sprintf("%s previous/trend %.1f/%v, want %.1f/%v",
				d.ID, d.PreviousScore, d.Trend, w.PreviousScore, w.Trend))
		}
	}
	if len(diffs) > 0 {
		return fmt.Errorf("seed history fields are stale: %s", strings.Join(diffs, "; "))
	}
	return nil
}

// sortRunsForCurve orders runs by measurement year, then publication date,
// then run ID — the same order used for every exported history table.
func sortRunsForCurve(runs []ingest.AnalysisRunJSON) {
	sort.SliceStable(runs, func(i, j int) bool {
		if runs[i].MeasurementYear != runs[j].MeasurementYear {
			return runs[i].MeasurementYear < runs[j].MeasurementYear
		}
		if runs[i].PublishedAt != runs[j].PublishedAt {
			return runs[i].PublishedAt < runs[j].PublishedAt
		}
		return runs[i].RunID < runs[j].RunID
	})
}

func round1(v float64) float64 { return math.Round(v*10) / 10 }

func close1(a, b float64) bool { return math.Abs(a-b) < scoreTolerance }

func sameSeries(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !close1(a[i], b[i]) {
			return false
		}
	}
	return true
}
