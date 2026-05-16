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
		"AI-Generated Code Output Share": 39.2,
		"Technical Work AI Value Share":  50.0,
		"AI Workflow Reliance":           67.7,
		"Agentic Task Delegation":        17.7,
		"Tool Ecosystem Reach":           84.4,
	}
	got := ComputeDomainScore(values, cfg)
	// Uses the current CodeGenConfig weights.
	expected := 39.2*0.30 + 50.0*0.25 + 67.7*0.25 + 17.7*0.15 + 84.4*0.05
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
		"Meta Automated Detection":     95.2,
		"YouTube Automated Flagging":   99.5,
		"TikTok Automated Enforcement": 93.8,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 95.2*0.375 + 99.5*0.3125 + 93.8*0.3125
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("content-mod: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_AlgoTrade(t *testing.T) {
	cfg := AlgoTradeConfig()
	values := map[string]float64{
		"FX Electronic Trading Share":          59.0,
		"Buy-Side AI Trade Execution Adoption": 15.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 59.0*0.55 + 15.0*0.45
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("algo-trade: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Support(t *testing.T) {
	cfg := SupportConfig()
	values := map[string]float64{
		"Cases Handled by AI":                          30.0,
		"Bot Deflection Rate":                          52.3,
		"Production AI Customer Communications Agents": 62.0,
		"Mature AI Support Deployment":                 10.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 30.0*0.45 + 52.3*0.20 + 62.0*0.25 + 10.0*0.10
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("support: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Credit(t *testing.T) {
	cfg := CreditConfig()
	values := map[string]float64{
		"AI-Underwritten Personal Loan Proxy": 38.2,
		"AI Credit Decisioning (Banks)":       32.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 38.2*0.55 + 32.0*0.45
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("credit: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_MedicalDx(t *testing.T) {
	cfg := MedicalDxConfig()
	// FDA: 1,430 devices -> LinearClamp(0,2000) -> 71.5
	values := map[string]float64{
		"FDA AI-Enabled Medical Devices":   71.5,
		"Radiology or Imaging AI Adoption": 50.0,
		"AI-Assisted Diagnosis Rate":       17.0,
		"Pathology AI Adoption":            10.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 71.5*0.10 + 50.0*0.35 + 17.0*0.30 + 10.0*0.25
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("medical-dx: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_LegalAI(t *testing.T) {
	cfg := LegalAIConfig()
	values := map[string]float64{
		"Legal Organization GenAI Adoption":            40.0,
		"Solo and Small Firms Using AI for Legal Work": 73.0,
		"AI-Assisted Document Review":                  28.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 40.0*0.40 + 73.0*0.30 + 28.0*0.30
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("legal-ai: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Hire(t *testing.T) {
	cfg := HireConfig()
	// AI Assessment Platform Reach: 80 M/yr -> LinearClamp(0,300) -> 26.67
	values := map[string]float64{
		"Orgs Using AI in Talent Acquisition": 69.0,
		"AI Screening Use Case Adoption":      58.0,
		"Broad AI Across Hiring Processes":    18.0,
		"AI Assessment Platform Reach":        26.67,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 69.0*0.30 + 58.0*0.35 + 18.0*0.15 + 26.67*0.20
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("hire: got %v, want ~%v", got, expected)
	}
}

func TestComputeDomainScore_Education(t *testing.T) {
	cfg := EducationConfig()
	values := map[string]float64{
		"Students Using AI for Schoolwork": 54.0,
		"AI-Graded Assessments":            8.0,
		"Teachers Using AI for Work":       37.0,
		"Student Papers 80%+ AI-Written":   15.0,
	}
	got := ComputeDomainScore(values, cfg)
	expected := 54.0*0.40 + 8.0*0.20 + 37.0*0.30 + 15.0*0.10
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

func TestShouldAppendTrend(t *testing.T) {
	if !ShouldAppendTrend(time.Time{}) {
		t.Fatal("zero last run should append")
	}
	if ShouldAppendTrend(time.Now().Add(-24 * time.Hour)) {
		t.Fatal("recent last run should update current point")
	}
	if !ShouldAppendTrend(time.Now().Add(-30 * 24 * time.Hour)) {
		t.Fatal("older last run should append")
	}
}
