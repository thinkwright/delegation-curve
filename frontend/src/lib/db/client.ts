import { query } from './boot';
import { Q } from './queries';
import type { DelegationDomain, DataSource, ScorePoint } from '$lib/data/types';

// Headline data (meta.json) is loaded by the root layout and prerendered;
// this client only serves the Parquet-backed domain detail.
const domainDetailPromises = new Map<string, Promise<DelegationDomain | null>>();

export async function getDomainDetail(id: string): Promise<DelegationDomain | null> {
	let promise = domainDetailPromises.get(id);
	if (!promise) {
		promise = loadDomainDetail(id).catch((e) => {
			domainDetailPromises.delete(id);
			throw e;
		});
		domainDetailPromises.set(id, promise);
	}
	return promise;
}

async function loadDomainDetail(id: string): Promise<DelegationDomain | null> {
	const [domainRows, subRows, sourceRows, historyRows] = await Promise.all([
		query(Q.delegationById(id), ['delegation']),
		query(Q.subIndicators(id), ['sub_indicators']),
		query(Q.dataSources(id), ['data_sources']),
		query(Q.domainRunHistory(id), ['domain_scores', 'analysis_runs'])
	]);
	if (!domainRows.length) return null;
	const r = domainRows[0];
	return {
		id: r.id as string,
		name: r.name as string,
		fullName: r.full_name as string,
		score: r.score as number,
		previousScore: r.previous_score as number,
		trend: JSON.parse(r.trend as string),
		status: r.status as 'nominal' | 'elevated' | 'autonomous',
		weight: r.weight as number,
		tier: r.tier as 1 | 2 | 3,
		runHistory: historyRows.map(scorePointFromRow),
		subIndicators: subRows.map((s) => ({
			name: s.name as string,
			value: s.value as number,
			unit: s.unit as string,
			source: s.source as string,
			freshness: s.freshness as string,
			sourceUrl: (s.source_url as string) || undefined,
			publishedAt: (s.published_at as string) || undefined,
			measurementStart: (s.measurement_start as string) || undefined,
			measurementEnd: (s.measurement_end as string) || undefined,
			evidenceGrade: (s.evidence_grade as string) || undefined,
			confidence: (s.confidence as string) || undefined,
			transform: (s.transform as string) || undefined
		})),
		dataSources: sourceRows.map((s) => ({
			name: s.name as string,
			cadence: s.cadence as string,
			type: s.type as DataSource['type']
		})),
		description: r.description as string
	};
}

function scorePointFromRow(r: Record<string, unknown>): ScorePoint {
	return {
		runId: r.run_id as string,
		label: r.label as string,
		publishedAt: r.published_at as string,
		measurementPeriod: r.measurement_period as string,
		measurementYear: r.measurement_year as number,
		methodologyVersion: r.methodology_version as string,
		score: r.score as number,
		notes: r.notes as string,
		isCurrent: r.is_current as boolean,
		periodMidpoint: (r.period_midpoint as string) || undefined
	};
}
