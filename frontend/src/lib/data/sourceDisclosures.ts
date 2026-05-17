import type { SourceDisclosure } from './types';

export const sourceDisclosures: Record<string, SourceDisclosure[]> = {
	credit: [
		{
			name: 'Upstart 2025 Form 10-K',
			role: 'Scored automation numerator',
			cadence: 'Annual',
			type: 'filing',
			url: 'https://www.sec.gov/Archives/edgar/data/1647639/000164763926000027/upst-20251231.htm',
			note: 'Provides the 91% fully automated platform-loan numerator used inside the personal-loan proxy; not treated as broad credit-market share.'
		},
		{
			name: 'TransUnion Q4 2025 CIIR',
			role: 'Scored market denominator',
			cadence: 'Quarterly',
			type: 'report',
			url: 'https://newsroom.transunion.com/q4-2025-ciir/',
			note: 'Provides the 42% fintech unsecured personal-loan origination denominator used with the Upstart platform automation rate.'
		},
		{
			name: 'McKinsey Credit Risk Survey',
			role: 'Conservative bank-decisioning hold',
			cadence: 'Annual',
			type: 'survey',
			note: 'Retained as a cautious bank credit-decisioning signal while direct 2026 bank underwriting deployment data remains unavailable.'
		},
		{
			name: 'Bank of England/FCA AI in UK Financial Services 2024',
			role: 'Automation and governance context',
			cadence: 'Periodic',
			type: 'survey',
			url: 'https://www.bankofengland.co.uk/report/2024/artificial-intelligence-in-uk-financial-services-2024',
			note: 'Broad financial-services AI survey used as a guardrail; useful for automated-decision context but not credit-specific enough to raise the bank score.'
		},
		{
			name: 'Bank of England 2026 AI Roundtables',
			role: 'Deployment-friction context',
			cadence: 'Periodic',
			type: 'roundtable',
			url: 'https://www.bankofengland.co.uk/minutes/2026/february/summary-of-ai-roundtables-feb-2026',
			note: 'Supports conservative treatment of regulated-bank deployment because risk, model validation, and governance functions remain cautious.'
		}
	]
};
