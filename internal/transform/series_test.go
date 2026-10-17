package transform

import (
	"strings"
	"testing"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// seriesSeed mirrors the published shape: an archived legacy run, two public
// restated runs, and a current run whose stored history fields are stale.
func seriesSeed() *ingest.Seed {
	return &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{Current: 45.8, Previous: 99, Delta: 99, Trend: []float64{1}},
			Domains: []ingest.DomainJSON{
				{ID: "a", Score: 50, PreviousScore: 99, Trend: []float64{1}},
				{ID: "b", Score: 40},
			},
		},
		AnalysisRuns: []ingest.AnalysisRunJSON{
			{RunID: "2026-q2", MeasurementYear: 2026, PublishedAt: "2026-05-16", CompositeScore: 45.8, IsCurrent: true},
			{RunID: "legacy-2025", MeasurementYear: 2025, PublishedAt: "2026-03-02", CompositeScore: 46, IsPublicSeries: boolPtr(false)},
			{RunID: "current-method-2025", MeasurementYear: 2025, PublishedAt: "2026-05-16", CompositeScore: 37.7},
			{RunID: "current-method-2024", MeasurementYear: 2024, PublishedAt: "2026-05-16", CompositeScore: 31.5},
		},
		DomainScores: []ingest.DomainScoreJSON{
			{RunID: "current-method-2024", DomainID: "a", Score: 30},
			{RunID: "current-method-2025", DomainID: "a", Score: 35},
			{RunID: "2026-q2", DomainID: "a", Score: 50},
			{RunID: "legacy-2025", DomainID: "a", Score: 60},
			{RunID: "current-method-2025", DomainID: "b", Score: 38},
			{RunID: "2026-q2", DomainID: "b", Score: 40},
		},
	}
}

func TestPublicSeriesOrdersAndSkipsArchivedRuns(t *testing.T) {
	s, ok, err := PublicSeries(seriesSeed())
	if err != nil || !ok {
		t.Fatalf("PublicSeries: ok=%v err=%v", ok, err)
	}
	var ids []string
	for _, r := range s.Public {
		ids = append(ids, r.RunID)
	}
	if got := strings.Join(ids, ","); got != "current-method-2024,current-method-2025,2026-q2" {
		t.Fatalf("public order = %s", got)
	}
	if s.Current.RunID != "2026-q2" || s.Previous.RunID != "current-method-2025" {
		t.Fatalf("current/previous = %s/%s", s.Current.RunID, s.Previous.RunID)
	}
}

func TestDeriveHistoryFromRuns(t *testing.T) {
	seed := seriesSeed()
	if err := DeriveHistory(seed); err != nil {
		t.Fatal(err)
	}
	c := seed.Delegation.Composite
	if c.Previous != 37.7 || c.Delta != 8.1 || !sameSeries(c.Trend, []float64{31.5, 37.7, 45.8}) || c.LastUpdated != "2026-05-16" {
		t.Fatalf("composite = %+v", c)
	}
	a, b := seed.Delegation.Domains[0], seed.Delegation.Domains[1]
	if a.PreviousScore != 35 || !sameSeries(a.Trend, []float64{30, 35, 50}) {
		t.Fatalf("domain a = prev %v trend %v", a.PreviousScore, a.Trend)
	}
	// Domain b has no 2024 point; its trend skips the gap.
	if b.PreviousScore != 38 || !sameSeries(b.Trend, []float64{38, 40}) {
		t.Fatalf("domain b = prev %v trend %v", b.PreviousScore, b.Trend)
	}
	if err := CheckDerived(seed); err != nil {
		t.Fatalf("CheckDerived after DeriveHistory: %v", err)
	}
}

func TestDeriveHistoryAdvancesWithNewRun(t *testing.T) {
	seed := seriesSeed()
	for i := range seed.AnalysisRuns {
		seed.AnalysisRuns[i].IsCurrent = false
	}
	seed.AnalysisRuns = append(seed.AnalysisRuns, ingest.AnalysisRunJSON{
		RunID: "2026-q4", MeasurementYear: 2026, PublishedAt: "2026-11-15", CompositeScore: 47.0, IsCurrent: true,
	})
	seed.DomainScores = append(seed.DomainScores,
		ingest.DomainScoreJSON{RunID: "2026-q4", DomainID: "a", Score: 52},
		ingest.DomainScoreJSON{RunID: "2026-q4", DomainID: "b", Score: 41},
	)
	if err := DeriveHistory(seed); err != nil {
		t.Fatal(err)
	}
	c := seed.Delegation.Composite
	if c.Previous != 45.8 || c.Delta != 1.2 || len(c.Trend) != 4 {
		t.Fatalf("composite after Q4 = %+v", c)
	}
	if seed.Delegation.Domains[0].PreviousScore != 50 {
		t.Fatalf("domain a previous = %v, want the 2026-q2 score", seed.Delegation.Domains[0].PreviousScore)
	}
}

func TestCheckDerivedDetectsStaleFields(t *testing.T) {
	if err := CheckDerived(seriesSeed()); err == nil {
		t.Fatal("expected stale history fields to be reported")
	}
}

func TestCheckSnapshotted(t *testing.T) {
	seed := seriesSeed()
	if err := CheckSnapshotted(seed); err != nil {
		t.Fatalf("snapshotted seed reported: %v", err)
	}
	seed.Delegation.Domains[0].Score = 51
	seed.Delegation.Composite.Current = 46.1
	err := CheckSnapshotted(seed)
	if err == nil || !strings.Contains(err.Error(), "a 51.0") || !strings.Contains(err.Error(), "composite 46.1") {
		t.Fatalf("unsnapshotted scores not reported: %v", err)
	}
}

func TestPublicSeriesRequiresOneCurrentRun(t *testing.T) {
	seed := seriesSeed()
	seed.AnalysisRuns[0].IsCurrent = false
	if _, _, err := PublicSeries(seed); err == nil {
		t.Fatal("expected error without a current run")
	}
	seed.AnalysisRuns[0].IsCurrent = true
	seed.AnalysisRuns[2].IsCurrent = true
	if _, _, err := PublicSeries(seed); err == nil {
		t.Fatal("expected error with two current runs")
	}
}

func TestLegacySeedIsUntouched(t *testing.T) {
	seed := &ingest.Seed{Delegation: ingest.DelegationSeed{Composite: ingest.CompositeDelegation{Previous: 1, Delta: 2, Trend: []float64{3}}}}
	if err := DeriveHistory(seed); err != nil {
		t.Fatal(err)
	}
	if c := seed.Delegation.Composite; c.Previous != 1 || c.Delta != 2 || len(c.Trend) != 1 {
		t.Fatalf("legacy seed changed: %+v", c)
	}
}
