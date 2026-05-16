import { query } from './boot';
import { Q } from './queries';
import { base } from '$app/paths';
import type {
	CompositeData,
	DelegationDomain,
	DataSource,
	ScorePoint,
} from '$lib/data/types';

let metaPromise: Promise<CompositeData> | null = null;
let domainsPromise: Promise<DelegationDomain[]> | null = null;
const domainDetailPromises = new Map<string, Promise<DelegationDomain | null>>();

export async function getMeta(): Promise<CompositeData> {
	if (!metaPromise) {
		metaPromise = fetch(`${base || ''}/data/meta.json`)
			.then((resp) => {
				if (!resp.ok) throw new Error(`Failed to load meta.json: ${resp.status}`);
				return resp.json();
			})
			.catch((e) => {
				metaPromise = null;
				throw e;
			});
	}
	return metaPromise;
}

export async function getDomains(): Promise<DelegationDomain[]> {
	if (!domainsPromise) {
		domainsPromise = loadDomains().catch((e) => {
			domainsPromise = null;
			throw e;
		});
	}
	return domainsPromise;
}

async function loadDomains(): Promise<DelegationDomain[]> {
	const rows = await query(Q.delegationAll, ['delegation']);
	return rows.map((r) => ({
		id: r.id as string,
		name: r.name as string,
		fullName: r.full_name as string,
		score: r.score as number,
		previousScore: r.previous_score as number,
		trend: JSON.parse(r.trend as string),
		status: r.status as 'nominal' | 'elevated' | 'autonomous',
		weight: r.weight as number,
		tier: r.tier as 1 | 2 | 3,
		subIndicators: [],
		dataSources: [],
		description: r.description as string
	}));
}

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
			freshness: s.freshness as string
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
		isCurrent: r.is_current as boolean
	};
}
