package collect

import "time"

// Cadence represents the native data collection frequency of an indicator.
type Cadence int

const (
	Annual     Cadence = iota // default; surveys, annual reports
	Quarterly                 // transparency reports, quarterly filings
	Continuous                // live databases, marketplaces
)

// StalenessThreshold returns the age threshold beyond which data at this
// cadence should be considered stale.
func (c Cadence) StalenessThreshold() time.Duration {
	switch c {
	case Continuous:
		return 30 * 24 * time.Hour // 30 days
	case Quarterly:
		return 120 * 24 * time.Hour // 4 months
	default:
		return 400 * 24 * time.Hour // ~13 months
	}
}

// DomainConfig defines scoring configuration for one domain.
type DomainConfig struct {
	DomainID    string
	DomainName  string
	FullName    string
	Description string
	Weight      float64 // domain weight in composite score
	Tier        int
	Indicators  []IndicatorConfig
}

// IndicatorConfig defines one sub-indicator's normalization and weight.
type IndicatorConfig struct {
	Name       string
	Weight     float64 // weight within this domain (should sum to 1.0)
	NormConfig NormConfig
	SourceName string
	Cadence    Cadence // native data collection frequency
}

// CodeGenConfig returns the configuration for the code-gen domain.
func CodeGenConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "code-gen",
		DomainName:  "CODE-GEN",
		FullName:    "Software Development",
		Description: "Percentage of production code substantially generated or influenced by AI coding assistants.",
		Weight:      0.15,
		Tier:        1,
		Indicators: []IndicatorConfig{
			{
				Name:       "Copilot Code Acceptance",
				Weight:     0.35,
				NormConfig: NormConfig{Method: DirectPercent},
				SourceName: "GitHub Octoverse",
			},
			{
				Name:       "Developer AI Tool Usage",
				Weight:     0.15,
				NormConfig: NormConfig{Method: DirectPercent},
				SourceName: "Stack Overflow Survey",
			},
			{
				Name:       "AI-Assisted Commits (OSS)",
				Weight:     0.40,
				NormConfig: NormConfig{Method: DirectPercent},
				SourceName: "GitClear Analysis",
			},
			{
				Name:       "IDE AI Extension Installs",
				Weight:     0.10,
				NormConfig: NormConfig{Method: LogScale, LogMin: 0, LogMax: 2.3}, // value stored in millions: 1M→0, 200M→2.3
				SourceName: "VS Code Marketplace",
				Cadence:    Continuous,
			},
		},
	}
}

// ContentModConfig returns the configuration for the content-mod domain.
func ContentModConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "content-mod",
		DomainName:  "CONTENT-MOD",
		FullName:    "Content Moderation",
		Description: "Percentage of content moderation decisions made or assisted by AI systems.",
		Weight:      0.10,
		Tier:        1,
		Indicators: []IndicatorConfig{
			{Name: "Meta Automated Detection", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Meta Transparency Report", Cadence: Quarterly},
			{Name: "Google Automated Removal", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Google Transparency Report", Cadence: Quarterly},
			{Name: "TikTok Automated Detection", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "TikTok Transparency Report", Cadence: Quarterly},
			{Name: "X/Twitter Automated Action", Weight: 0.20, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "X Transparency Report"},
		},
	}
}

// AlgoTradeConfig returns the configuration for the algo-trade domain.
func AlgoTradeConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "algo-trade",
		DomainName:  "ALGO-TRADE",
		FullName:    "Algorithmic Trading",
		Description: "Percentage of financial trading volume executed by AI/algorithmic systems.",
		Weight:      0.15,
		Tier:        1,
		Indicators: []IndicatorConfig{
			{Name: "US Equities Algo Volume", Weight: 0.35, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "SEC Market Structure"},
			{Name: "FX Algo Trading", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "BIS Triennial Survey"},
			{Name: "Options Algo Volume", Weight: 0.20, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "CBOE Data"},
			{Name: "Institutional AI Adoption", Weight: 0.20, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Greenwich Associates"},
		},
	}
}

// SupportConfig returns the configuration for the support domain.
func SupportConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "support",
		DomainName:  "SUPPORT",
		FullName:    "Customer Support",
		Description: "Percentage of customer support interactions handled or assisted by AI.",
		Weight:      0.15,
		Tier:        1,
		Indicators: []IndicatorConfig{
			{Name: "AI Resolution Rate", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Zendesk CX Trends"},
			{Name: "Bot Deflection Rate", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Intercom Trends"},
			{Name: "Orgs Using AI Support", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Salesforce State of Service"},
			{Name: "AI Copilot Adoption (Agents)", Weight: 0.20, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Zendesk/Salesforce Reports"},
		},
	}
}

// CreditConfig returns the configuration for the credit domain.
func CreditConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "credit",
		DomainName:  "CREDIT",
		FullName:    "Credit & Lending",
		Description: "Degree of AI involvement in credit underwriting and lending decisions.",
		Weight:      0.10,
		Tier:        2,
		Indicators: []IndicatorConfig{
			{Name: "AI-Underwritten Loan Volume", Weight: 0.40, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Fintech Filings", Cadence: Quarterly},
			{Name: "Fintech Lending Market Share", Weight: 0.15, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Industry Reports", Cadence: Quarterly},
			{Name: "AI Credit Decisioning (Banks)", Weight: 0.45, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "OCC Survey"},
		},
	}
}

// MedicalDxConfig returns the configuration for the medical-dx domain.
func MedicalDxConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "medical-dx",
		DomainName:  "MEDICAL-DX",
		FullName:    "Medical Diagnosis",
		Description: "Degree of AI involvement in clinical diagnosis and medical imaging interpretation.",
		Weight:      0.12,
		Tier:        2,
		Indicators: []IndicatorConfig{
			{Name: "FDA-Cleared Diagnostic AI Devices", Weight: 0.10, NormConfig: NormConfig{Method: LinearClamp, Min: 0, Max: 1200}, SourceName: "FDA AI/ML Database", Cadence: Continuous},
			{Name: "Radiology AI Adoption", Weight: 0.35, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "RSNA Survey"},
			{Name: "AI-Assisted Diagnosis Rate", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "AMA Survey"},
			{Name: "Pathology AI Adoption", Weight: 0.25, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "CAP Survey"},
		},
	}
}

// LegalAIConfig returns the configuration for the legal-ai domain.
func LegalAIConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "legal-ai",
		DomainName:  "LEGAL-AI",
		FullName:    "Legal Research & Review",
		Description: "Degree of AI adoption in legal research, document review, and court proceedings.",
		Weight:      0.08,
		Tier:        2,
		Indicators: []IndicatorConfig{
			{Name: "AI Tool Adoption (BigLaw)", Weight: 0.40, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "ALM Survey"},
			{Name: "AI Tool Adoption (Solo/Small)", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Clio Legal Trends"},
			{Name: "AI-Assisted Document Review", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "EDRM/Relativity Survey"},
		},
	}
}

// HireConfig returns the configuration for the hire domain.
func HireConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "hire",
		DomainName:  "HIRE",
		FullName:    "Recruitment & Screening",
		Description: "Degree of AI involvement in hiring, screening, and employment decisions.",
		Weight:      0.08,
		Tier:        2,
		Indicators: []IndicatorConfig{
			{Name: "Orgs Using AI Screening", Weight: 0.40, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "SHRM Survey"},
			{Name: "AI-Screened Applications", Weight: 0.35, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "SHRM/LinkedIn Data"},
			{Name: "AI Assessment Platform Reach", Weight: 0.25, NormConfig: NormConfig{Method: LinearClamp, Min: 0, Max: 300}, SourceName: "Platform Aggregation"},
		},
	}
}

// EducationConfig returns the configuration for the education domain.
func EducationConfig() DomainConfig {
	return DomainConfig{
		DomainID:    "education",
		DomainName:  "EDUCATION",
		FullName:    "Education & Assessment",
		Description: "Degree of AI adoption in tutoring, assessment, and academic integrity monitoring.",
		Weight:      0.07,
		Tier:        3,
		Indicators: []IndicatorConfig{
			{Name: "Students Using AI Tutors", Weight: 0.35, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Educause Survey"},
			{Name: "AI-Graded Assessments", Weight: 0.30, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "Gradescope/EdTech Reports"},
			{Name: "Faculty Using AI in Teaching", Weight: 0.35, NormConfig: NormConfig{Method: DirectPercent}, SourceName: "AAUP/Educause Survey"},
		},
	}
}

// AllDomainConfigs returns all 9 domain configurations.
func AllDomainConfigs() []DomainConfig {
	return []DomainConfig{
		ContentModConfig(),
		AlgoTradeConfig(),
		CodeGenConfig(),
		SupportConfig(),
		CreditConfig(),
		MedicalDxConfig(),
		LegalAIConfig(),
		HireConfig(),
		EducationConfig(),
	}
}

// ComputeDomainScore computes a weighted average of normalized indicator values.
func ComputeDomainScore(values map[string]float64, cfg DomainConfig) float64 {
	var weightedSum, totalWeight float64
	for _, ind := range cfg.Indicators {
		if val, ok := values[ind.Name]; ok {
			weightedSum += val * ind.Weight
			totalWeight += ind.Weight
		}
	}
	if totalWeight == 0 {
		return 0
	}
	return weightedSum / totalWeight
}

// ClassifyStatus maps a domain score to a status label.
func ClassifyStatus(score float64) string {
	switch {
	case score >= 75:
		return "autonomous"
	case score >= 40:
		return "elevated"
	default:
		return "nominal"
	}
}

// UpdateTrend appends or updates the latest point in a trend array.
// If the last run was <25 days ago, updates the last point (same-month re-run).
// Otherwise appends. Caps at maxPoints (default 84 = 7 years monthly).
func UpdateTrend(existing []float64, newScore float64, lastRunAt time.Time, maxPoints int) []float64 {
	if maxPoints <= 0 {
		maxPoints = 84
	}

	result := make([]float64, len(existing))
	copy(result, existing)

	daysSinceLastRun := time.Since(lastRunAt).Hours() / 24

	if daysSinceLastRun < 25 && len(result) > 0 {
		result[len(result)-1] = newScore
		return result
	}

	result = append(result, newScore)
	if len(result) > maxPoints {
		result = result[len(result)-maxPoints:]
	}
	return result
}
