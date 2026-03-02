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
