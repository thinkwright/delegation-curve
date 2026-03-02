package collect

import (
	"math"
	"testing"
)

func TestNormalize_DirectPercent(t *testing.T) {
	cfg := NormConfig{Method: DirectPercent}
	tests := []struct {
		raw  float64
		want float64
	}{
		{50, 50},
		{0, 0},
		{100, 100},
		{-5, 0},   // clamped
		{120, 100}, // clamped
		{48.2, 48.2},
	}
	for _, tt := range tests {
		got := Normalize(tt.raw, cfg)
		if got != tt.want {
			t.Errorf("DirectPercent(%v) = %v, want %v", tt.raw, got, tt.want)
		}
	}
}

func TestNormalize_LinearClamp(t *testing.T) {
	cfg := NormConfig{Method: LinearClamp, Min: 10, Max: 110}
	tests := []struct {
		raw  float64
		want float64
	}{
		{10, 0},
		{110, 100},
		{60, 50},
		{0, 0},    // below min → clamped to 0
		{200, 100}, // above max → clamped to 100
	}
	for _, tt := range tests {
		got := Normalize(tt.raw, cfg)
		if got != tt.want {
			t.Errorf("LinearClamp(%v) = %v, want %v", tt.raw, got, tt.want)
		}
	}
}

func TestNormalize_LinearClamp_EqualMinMax(t *testing.T) {
	cfg := NormConfig{Method: LinearClamp, Min: 50, Max: 50}
	got := Normalize(50, cfg)
	if got != 0 {
		t.Errorf("LinearClamp equal min/max = %v, want 0", got)
	}
}

func TestNormalize_LogScale(t *testing.T) {
	cfg := NormConfig{Method: LogScale, LogMin: 0, LogMax: 2.3}
	tests := []struct {
		raw     float64
		wantMin float64
		wantMax float64
	}{
		{1, 0, 0},           // log10(1) = 0 → 0%
		{200, 99, 101},      // log10(200) ≈ 2.3 → ~100%
		{71.8, 75, 85},      // log10(71.8) ≈ 1.856 → ~80.7%
		{0, 0, 0},           // zero → 0
		{-5, 0, 0},          // negative → 0
	}
	for _, tt := range tests {
		got := Normalize(tt.raw, cfg)
		if got < tt.wantMin || got > tt.wantMax {
			t.Errorf("LogScale(%v) = %v, want [%v, %v]", tt.raw, got, tt.wantMin, tt.wantMax)
		}
	}
}

func TestNormalize_LinearClamp_DiagnosticFDA(t *testing.T) {
	cfg := NormConfig{Method: LinearClamp, Min: 0, Max: 1200}
	got := Normalize(950, cfg)
	expected := 950.0 / 1200.0 * 100
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("FDA diagnostic 950/1200 = %v, want ~%v", got, expected)
	}
}

func TestNormalize_LinearClamp_AssessmentPlatformReach(t *testing.T) {
	cfg := NormConfig{Method: LinearClamp, Min: 0, Max: 300}
	got := Normalize(70, cfg)
	expected := 70.0 / 300.0 * 100
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("Assessment platform reach 70/300 = %v, want ~%v", got, expected)
	}
}

func TestNormalize_NaN(t *testing.T) {
	cfg := NormConfig{Method: DirectPercent}
	if got := Normalize(math.NaN(), cfg); got != 0 {
		t.Errorf("NaN input = %v, want 0", got)
	}
}

func TestNormalize_Inf(t *testing.T) {
	cfg := NormConfig{Method: DirectPercent}
	if got := Normalize(math.Inf(1), cfg); got != 0 {
		t.Errorf("+Inf input = %v, want 0", got)
	}
	if got := Normalize(math.Inf(-1), cfg); got != 0 {
		t.Errorf("-Inf input = %v, want 0", got)
	}
}
