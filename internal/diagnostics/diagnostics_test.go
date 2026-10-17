package diagnostics

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

func fixture() []Input {
	return []Input{
		{DomainID: "a", Indicator: "work", Construct: collect.Work, Grade: "A", Confidence: "high", MeasurementEnd: "2026-01-31", DomainWeight: 0.5, Weight: 0.5, Normalized: 80},
		{DomainID: "a", Indicator: "adopt", Construct: collect.Adoption, Grade: "B", Confidence: "medium", MeasurementEnd: "2025", DomainWeight: 0.5, Weight: 0.5, Normalized: 40},
		{DomainID: "b", Indicator: "usage", Construct: collect.Usage, Grade: "C", Confidence: "low", MeasurementEnd: "2024-06", DomainWeight: 0.5, Weight: 0.75, Normalized: 20},
		{DomainID: "b", Indicator: "count", Construct: collect.Capacity, Grade: "D", Confidence: "medium", DomainWeight: 0.5, Weight: 0.25, Normalized: 60},
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestCompositeMatchesWeightedAverage(t *testing.T) {
	// a = 60, b = 0.75*20 + 0.25*60 = 30; composite = 45.
	if got, cov := Composite(fixture(), nil); !near(got, 45) || !near(cov, 1) {
		t.Fatalf("Composite = %v (coverage %v), want 45 (1)", got, cov)
	}
	// Only grade A/B: domain b drops out, a = 60.
	got, cov := Composite(fixture(), func(in Input) bool { return in.Grade == "A" || in.Grade == "B" })
	if !near(got, 60) || !near(cov, 0.5) {
		t.Fatalf("A/B composite = %v (coverage %v), want 60 (0.5)", got, cov)
	}
}

func TestShadowSplitsDepthAndReach(t *testing.T) {
	s := ComputeShadow(fixture())
	if !near(s.Depth, 80) || !near(s.DepthCoverage, 0.5) {
		t.Fatalf("depth = %v (%v)", s.Depth, s.DepthCoverage)
	}
	if !near(s.Reach, 30) || !near(s.ReachCoverage, 1) { // a reach 40, b reach 20
		t.Fatalf("reach = %v (%v)", s.Reach, s.ReachCoverage)
	}
	// Only domain a has both: 0.7*80 + 0.3*40 = 68.
	if !near(s.Combined, 68) || !near(s.CombinedCoverage, 0.5) {
		t.Fatalf("combined = %v (%v)", s.Combined, s.CombinedCoverage)
	}
}

func TestUncertaintyIsReproducibleAndBracketsPoint(t *testing.T) {
	a := Uncertainty(fixture(), 4000, 7)
	b := Uncertainty(fixture(), 4000, 7)
	if a != b {
		t.Fatalf("same seed gave %+v and %+v", a, b)
	}
	if !(a.P5 < 45 && 45 < a.P95) || a.P5 >= a.P50 || a.P50 >= a.P95 {
		t.Fatalf("interval %+v does not bracket 45", a)
	}
}

func TestAnalyzeAgesAndMarkdown(t *testing.T) {
	r := Analyze("test", fixture(), 45, time.Date(2026, 5, 16, 0, 0, 0, 0, time.UTC))
	// Effective weights: work .25 (3.5 mo), adopt .25 (2025 -> 16.5 mo),
	// usage .375 (2024-06 -> 23.5 mo), count .125 undated.
	if !near(r.ShareOlder12m, 0.625) || !near(r.ShareUndated, 0.125) {
		t.Fatalf("age shares: >12m %v undated %v", r.ShareOlder12m, r.ShareUndated)
	}
	if r.MedianAgeMonths < 16 || r.MedianAgeMonths > 17 {
		t.Fatalf("median age = %.1f months", r.MedianAgeMonths)
	}
	md := r.Markdown()
	for _, want := range []string{"Diagnostics: run `test`", "Equal domain weights", "90% interval", "Delegation depth"} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q", want)
		}
	}
}
