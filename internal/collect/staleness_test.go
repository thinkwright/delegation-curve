package collect

import (
	"testing"
	"time"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

func TestParseDateFloor(t *testing.T) {
	tests := map[string]string{
		"2025-09-30":   "2025-09-30",
		"2025-09":      "2025-09-01",
		"2025":         "2025-01-01",
		"Q3 2025":      "2025-07-01",
		"H2 2025":      "2025-07-01",
		"Jul-Sep 2025": "2025-07-01",
		"April 2025":   "2025-04-01",
		"Mar 2026":     "2026-03-01",
	}
	for in, want := range tests {
		got, ok := parseDateFloor(in)
		if !ok || got.Format("2006-01-02") != want {
			t.Errorf("parseDateFloor(%q) = %v %v, want %s", in, got, ok, want)
		}
	}
	if _, ok := parseDateFloor("2025/2026"); ok {
		t.Error("ambiguous label should not parse")
	}
}

func TestOverrideAgeUsesMeasurementWindow(t *testing.T) {
	now := time.Date(2026, 10, 17, 0, 0, 0, 0, time.UTC)
	e := OverrideEntry{
		EnteredAt:  "2026-05-16",
		Freshness:  "2025",
		SourceMeta: ingest.SourceMeta{MeasurementEnd: "2024-09-30"},
	}
	if days := int(overrideAge(e, now).Hours() / 24); days != 747 {
		t.Fatalf("age = %d days, want 747 (since measurement end, not entry)", days)
	}
}
