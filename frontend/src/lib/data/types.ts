export interface SubIndicator {
	name: string;
	value: number;
	unit: string;
	source: string;
	freshness: string;
}

export interface DataSource {
	name: string;
	cadence: string;
	type: 'transparency-report' | 'survey' | 'api' | 'filing' | 'scrape';
}

export interface DelegationDomain {
	id: string;
	name: string;
	fullName: string;
	score: number;
	previousScore: number;
	trend: number[];
	status: 'nominal' | 'elevated' | 'autonomous';
	weight: number;
	tier: 1 | 2 | 3;
	subIndicators: SubIndicator[];
	dataSources: DataSource[];
	description: string;
}

export interface CompositeData {
	delegation: {
		current: number;
		previous: number;
		delta: number;
		trend: number[];
		lastUpdated: string;
		dataYear: number;
	};
	domainsTracked: number;
	highestDomain: { name: string; score: number };
	dataFreshness: string;
}
