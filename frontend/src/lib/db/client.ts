import { query } from './boot';
import { Q } from './queries';
import { base } from '$app/paths';
import type {
	CompositeData,
	DelegationDomain,
	DataSource,
} from '$lib/data/types';

export async function getMeta(): Promise<CompositeData> {
	const resp = await fetch(`${base || ''}/data/meta.json`);
	return resp.json();
}

export async function getDomains(): Promise<DelegationDomain[]> {
	const rows = await query(Q.delegationAll);
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
	const [domainRows, subRows, sourceRows] = await Promise.all([
		query(Q.delegationById(id)),
		query(Q.subIndicators(id)),
		query(Q.dataSources(id))
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
