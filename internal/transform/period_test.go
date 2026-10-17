package transform

import "testing"

func TestPeriodMidpoint(t *testing.T) {
	tests := []struct {
		period string
		year   int
		want   string
	}{
		{"2024", 2024, "2024-07-01"},
		{"2026 Q2", 2026, "2026-05-16"},
		{"Q4 2026", 2026, "2026-11-15"},
		{"2026 H2", 2026, "2026-09-30"},
		{"2025 recalculated baseline", 2025, "2025-07-02"},
	}
	for _, tt := range tests {
		if got := PeriodMidpoint(tt.period, tt.year); got != tt.want {
			t.Errorf("PeriodMidpoint(%q, %d) = %s, want %s", tt.period, tt.year, got, tt.want)
		}
	}
}
