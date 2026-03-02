package collect

import (
	"math"
	"testing"
	"time"
)

func TestComputeDomainScore_AllPresent(t *testing.T) {
	cfg := DomainConfig{
		Indicators: []IndicatorConfig{
			{Name: "A", Weight: 0.5},
			{Name: "B", Weight: 0.5},
		},
	}
	values := map[string]float64{"A": 80, "B": 60}
	got := ComputeDomainScore(values, cfg)
	if got != 70 {
		t.Errorf("got %v, want 70", got)
	}
}

func TestComputeDomainScore_MissingExcluded(t *testing.T) {
	cfg := DomainConfig{
		Indicators: []IndicatorConfig{
			{Name: "A", Weight: 0.5},
			{Name: "B", Weight: 0.5},
		},
	}
	// Only A is present — B is excluded, so score = A's value.
	values := map[string]float64{"A": 80}
	got := ComputeDomainScore(values, cfg)
	if got != 80 {
		t.Errorf("got %v, want 80 (B excluded, A re-weighted)", got)
	}
}

func TestComputeDomainScore_NonePresent(t *testing.T) {
	cfg := DomainConfig{
		Indicators: []IndicatorConfig{
			{Name: "A", Weight: 0.5},
		},
	}
	values := map[string]float64{}
	got := ComputeDomainScore(values, cfg)
	if got != 0 {
		t.Errorf("got %v, want 0", got)
	}
}

func TestComputeDomainScore_CodeGen(t *testing.T) {
	cfg := CodeGenConfig()
	values := map[string]float64{
		"Copilot Code Acceptance":  48.2,
		"Developer AI Tool Usage":  82.0,
		"AI-Assisted Commits (OSS)": 31.5,
		"IDE AI Extension Installs": 80.7,
	}
	got := ComputeDomainScore(values, cfg)
	// (48.2*0.35 + 82.0*0.15 + 31.5*0.40 + 80.7*0.10) / 1.0
	expected := 48.2*0.35 + 82.0*0.15 + 31.5*0.40 + 80.7*0.10
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("got %v, want ~%v", got, expected)
	}
}

func TestAllDomainWeightsSumToOne(t *testing.T) {
	configs := AllDomainConfigs()
	if len(configs) != 9 {
		t.Fatalf("expected 9 domains, got %d", len(configs))
	}
	var sum float64
	for _, c := range configs {
		sum += c.Weight
	}
	if math.Abs(sum-1.0) > 0.001 {
		t.Errorf("domain weights sum to %v, want 1.0", sum)
	}
}

func TestAllDomainIndicatorWeightsSumToOne(t *testing.T) {
	for _, cfg := range AllDomainConfigs() {
		var sum float64
		for _, ind := range cfg.Indicators {
			sum += ind.Weight
		}
		if math.Abs(sum-1.0) > 0.001 {
			t.Errorf("%s indicator weights sum to %v, want 1.0", cfg.DomainID, sum)
		}
	}
}

func TestComputeDomainScore_ContentMod(t *testing.T) {
	cfg := ContentModConfig()
	values := map[string]float64{
		"Meta Automated Detection":  95.2,
		"Google Automated Removal":  92.8,
		"TikTok Automated Detection": 94.1,
		"X/Twitter Automated Action": 88.6,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 95.2*0.30 + 92.8*0.25 + 94.1*0.25 + 88.6*0.20
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("content-mod: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_AlgoTrade(t *testing.T) {
	cfg := AlgoTradeConfig()
	values := map[string]float64{
		"US Equities Algo Volume":  73.2,
		"FX Algo Trading":         61.4,
		"Options Algo Volume":     58.9,
		"Institutional AI Adoption": 78.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 73.2*0.35 + 61.4*0.25 + 58.9*0.20 + 78.0*0.20
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("algo-trade: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Support(t *testing.T) {
	cfg := SupportConfig()
	values := map[string]float64{
		"AI Resolution Rate":           41.0,
		"Bot Deflection Rate":          52.3,
		"Orgs Using AI Support":        63.0,
		"AI Copilot Adoption (Agents)": 45.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 41.0*0.30 + 52.3*0.25 + 63.0*0.25 + 45.0*0.20
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("support: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Credit(t *testing.T) {
	cfg := CreditConfig()
	values := map[string]float64{
		"AI-Underwritten Loan Volume":   34.2,
		"Fintech Lending Market Share":  38.0,
		"AI Credit Decisioning (Banks)": 32.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 34.2*0.40 + 38.0*0.15 + 32.0*0.45
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("credit: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_MedicalDx(t *testing.T) {
	cfg := MedicalDxConfig()
	// FDA: 950 devices → LinearClamp(0,1200) → 79.17
	values := map[string]float64{
		"FDA-Cleared Diagnostic AI Devices": 79.17,
		"Radiology AI Adoption":             30.0,
		"AI-Assisted Diagnosis Rate":        12.0,
		"Pathology AI Adoption":             17.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 79.17*0.10 + 30.0*0.35 + 12.0*0.30 + 17.0*0.25
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("medical-dx: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_LegalAI(t *testing.T) {
	cfg := LegalAIConfig()
	values := map[string]float64{
		"AI Tool Adoption (BigLaw)":    42.0,
		"AI Tool Adoption (Solo/Small)": 18.5,
		"AI-Assisted Document Review":  52.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 42.0*0.40 + 18.5*0.30 + 52.0*0.30
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("legal-ai: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Hire(t *testing.T) {
	cfg := HireConfig()
	// AI Assessment Platform Reach: 70 M/yr → LinearClamp(0,300) → 23.33
	values := map[string]float64{
		"Orgs Using AI Screening":      28.0,
		"AI-Screened Applications":     55.0,
		"AI Assessment Platform Reach": 23.33,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 28.0*0.40 + 55.0*0.35 + 23.33*0.25
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("hire: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Education(t *testing.T) {
	cfg := EducationConfig()
	values := map[string]float64{
		"Students Using AI Tutors":     18.5,
		"AI-Graded Assessments":        8.0,
		"Faculty Using AI in Teaching": 26.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 18.5*0.35 + 8.0*0.30 + 26.0*0.35
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("education: got %v, want ~%v", got, expected)
	}
}

func TestClassifyStatus(t *testing.T) {
	tests := []struct {
		score float64
		want  string
	}{
		{93, "autonomous"},
		{75, "autonomous"},
		{74.9, "elevated"},
		{40, "elevated"},
		{39.9, "nominal"},
		{0, "nominal"},
	}
	for _, tt := range tests {
		if got := ClassifyStatus(tt.score); got != tt.want {
			t.Errorf("ClassifyStatus(%v) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestUpdateTrend_Append(t *testing.T) {
	existing := []float64{1, 2, 3}
	// Zero time → always append (old run).
	got := UpdateTrend(existing, 4, time.Time{}, 84)
	if len(got) != 4 || got[3] != 4 {
		t.Errorf("append: got %v", got)
	}
}

func TestUpdateTrend_UpdateLastPoint(t *testing.T) {
	existing := []float64{1, 2, 3}
	// Recent run → update last point.
	got := UpdateTrend(existing, 4, time.Now().Add(-24*time.Hour), 84)
	if len(got) != 3 || got[2] != 4 {
		t.Errorf("update last: got %v", got)
	}
}

func TestUpdateTrend_SlidingWindow(t *testing.T) {
	existing := make([]float64, 84)
	for i := range existing {
		existing[i] = float64(i)
	}
	got := UpdateTrend(existing, 100, time.Time{}, 84)
	if len(got) != 84 {
		t.Errorf("should cap at 84, got %d", len(got))
	}
	if got[83] != 100 {
		t.Errorf("last point should be 100, got %v", got[83])
	}
	if got[0] != 1 {
		t.Errorf("first point should be 1 (shifted), got %v", got[0])
	}
}

func TestUpdateTrend_EmptyExisting(t *testing.T) {
	got := UpdateTrend(nil, 42, time.Time{}, 84)
	if len(got) != 1 || got[0] != 42 {
		t.Errorf("empty existing: got %v", got)
	}
}

func TestUpdateTrend_DoesNotMutateOriginal(t *testing.T) {
	existing := []float64{1, 2, 3}
	UpdateTrend(existing, 4, time.Time{}, 84)
	if existing[2] != 3 {
		t.Errorf("original slice was mutated")
	}
}
