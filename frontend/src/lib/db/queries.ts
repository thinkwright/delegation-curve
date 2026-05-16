// DuckDB-WASM doesn't support parameterized queries, so we allowlist identifiers
// to prevent injection. Only lowercase alphanumeric + hyphens are permitted.
const VALID_ID = /^[a-z0-9-]+$/;

function safe(id: string): string {
	if (!VALID_ID.test(id)) throw new Error(`Invalid identifier: ${id}`);
	return id;
}

export const Q = {
	// Dashboard
	meta: `SELECT * FROM meta LIMIT 1`,

	// Delegation index (all domains, sorted by score)
	delegationAll: `
		SELECT id, name, full_name, score, previous_score,
		       trend, status, weight, tier, description
		FROM delegation
		ORDER BY score DESC
	`,

	// Single domain
	delegationById: (id: string) => `
		SELECT * FROM delegation WHERE id = '${safe(id)}'
	`,

	// Sub-indicators for a domain
	subIndicators: (domainId: string) => `
		SELECT name, value, unit, source, freshness
		FROM sub_indicators
		WHERE domain_id = '${safe(domainId)}'
	`,

	// Data sources for a domain
	dataSources: (domainId: string) => `
		SELECT name, cadence, type
		FROM data_sources
		WHERE domain_id = '${safe(domainId)}'
	`,

	// Composite run history, one row per analysis run.
	analysisRuns: `
		SELECT run_id, label, published_at, measurement_period, measurement_year,
		       methodology_version, composite_score, notes, is_current
		FROM analysis_runs
		WHERE is_public_series = true
		ORDER BY measurement_year, published_at, run_id
	`,

	// Domain score history joined to run metadata.
	domainRunHistory: (domainId: string) => `
		SELECT r.run_id, r.label, r.published_at, r.measurement_period,
		       r.measurement_year, r.methodology_version, d.score, r.notes,
		       r.is_current
		FROM domain_scores d
		JOIN analysis_runs r ON d.run_id = r.run_id
		WHERE d.domain_id = '${safe(domainId)}'
		  AND r.is_public_series = true
		ORDER BY r.measurement_year, r.published_at, r.run_id
	`,
};
