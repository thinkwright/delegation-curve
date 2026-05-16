package schema

// DelegationRow — 9 rows, one per domain.
type DelegationRow struct {
	ID            string  `parquet:"id,zstd"`
	Name          string  `parquet:"name,zstd"`
	FullName      string  `parquet:"full_name,zstd"`
	Score         float64 `parquet:"score"`
	PreviousScore float64 `parquet:"previous_score"`
	Trend         string  `parquet:"trend,zstd"`
	Status        string  `parquet:"status,zstd"`
	Weight        float64 `parquet:"weight"`
	Tier          int32   `parquet:"tier"`
	Description   string  `parquet:"description,zstd"`
}

// SubIndicatorRow — ~35 rows (3–5 per domain).
type SubIndicatorRow struct {
	DomainID  string  `parquet:"domain_id,zstd"`
	Name      string  `parquet:"name,zstd"`
	Value     float64 `parquet:"value"`
	Unit      string  `parquet:"unit,zstd"`
	Source    string  `parquet:"source,zstd"`
	Freshness string  `parquet:"freshness,zstd"`
}

// DataSourceRow — ~30 rows (3–4 per domain).
type DataSourceRow struct {
	DomainID string `parquet:"domain_id,zstd"`
	Name     string `parquet:"name,zstd"`
	Cadence  string `parquet:"cadence,zstd"`
	Type     string `parquet:"type,zstd"`
}

// MetaRow — 1 row with headline statistics.
type MetaRow struct {
	DelegationComposite float64 `parquet:"delegation_composite"`
	DelegationPrevious  float64 `parquet:"delegation_previous"`
	DelegationDelta     float64 `parquet:"delegation_delta"`
	DelegationTrend     string  `parquet:"delegation_trend,zstd"`
	DomainsTracked      int32   `parquet:"domains_tracked"`
	HighestDomainName   string  `parquet:"highest_domain_name,zstd"`
	HighestDomainScore  float64 `parquet:"highest_domain_score"`
	DataFreshness       string  `parquet:"data_freshness,zstd"`
	LastUpdated         string  `parquet:"last_updated,zstd"`
	DataYear            int32   `parquet:"data_year"`
}

// AnalysisRunRow — one row per published or derived analysis run.
type AnalysisRunRow struct {
	RunID              string  `parquet:"run_id,zstd"`
	Label              string  `parquet:"label,zstd"`
	PublishedAt        string  `parquet:"published_at,zstd"`
	MeasurementPeriod  string  `parquet:"measurement_period,zstd"`
	MeasurementYear    int32   `parquet:"measurement_year"`
	MethodologyVersion string  `parquet:"methodology_version,zstd"`
	CompositeScore     float64 `parquet:"composite_score"`
	Notes              string  `parquet:"notes,zstd"`
	IsCurrent          bool    `parquet:"is_current"`
	IsPublicSeries     bool    `parquet:"is_public_series"`
}

// DomainScoreRow — one row per domain per analysis run.
type DomainScoreRow struct {
	RunID    string  `parquet:"run_id,zstd"`
	DomainID string  `parquet:"domain_id,zstd"`
	Score    float64 `parquet:"score"`
	Weight   float64 `parquet:"weight"`
	Status   string  `parquet:"status,zstd"`
}

// IndicatorObservationRow — source-level observations for a run.
type IndicatorObservationRow struct {
	RunID           string  `parquet:"run_id,zstd"`
	DomainID        string  `parquet:"domain_id,zstd"`
	IndicatorName   string  `parquet:"indicator_name,zstd"`
	RawValue        float64 `parquet:"raw_value"`
	NormalizedValue float64 `parquet:"normalized_value"`
	Unit            string  `parquet:"unit,zstd"`
	Weight          float64 `parquet:"weight"`
	IncludedInScore bool    `parquet:"included_in_score"`
	Source          string  `parquet:"source,zstd"`
	SourceURL       string  `parquet:"source_url,zstd"`
	Freshness       string  `parquet:"freshness,zstd"`
	EvidenceGrade   string  `parquet:"evidence_grade,zstd"`
	Confidence      string  `parquet:"confidence,zstd"`
}
