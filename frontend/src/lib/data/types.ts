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
	type: string;
}

export interface SourceDisclosure {
	name: string;
	role: string;
	cadence: string;
	type: string;
	note: string;
	url?: string;
}

export interface ScorePoint {
	runId: string;
	label: string;
	publishedAt: string;
	measurementPeriod: string;
	measurementYear: number;
	methodologyVersion: string;
	score: number;
	notes?: string;
	isCurrent?: boolean;
}

export interface DelegationDomain {
	id: string;
	name: string;
	fullName: string;
	score: number;
	previousScore: number;
	trend: number[];
	runHistory?: ScorePoint[];
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
		runHistory?: ScorePoint[];
		lastUpdated: string;
		dataYear: number;
	};
	domainsTracked: number;
	highestDomain: { name: string; score: number };
	dataFreshness: string;
}
