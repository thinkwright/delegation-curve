# Credit 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `credit` score: 28.9.

Configured scoring formula:

- `AI-Underwritten Loan Volume`: 40% weight.
- `Fintech Lending Market Share`: 15% weight.
- `AI Credit Decisioning (Banks)`: 45% weight.

Current observations:

- `AI-Underwritten Loan Volume`: 34%, freshness 2025.
- `Fintech Lending Market Share`: 42%, freshness Q3 2025.
- `AI Credit Decisioning (Banks)`: 20%, freshness H2 2024.
- `UK Firms Using AI Credit Risk`: 16%, freshness 2024, display-only.

The current score calculates as:

```text
0.40 * 34 + 0.15 * 42 + 0.45 * 20 = 28.9
```

The main 2026 issue is formula clarity. A raw fintech platform automation rate is not a whole-market credit-decisioning rate. If we directly replace `AI-Underwritten Loan Volume` with Upstart's 91%, the score would jump to 51.7 before any bank update, which would overstate delegation across the full credit market.

## Extracted Candidate Sources

### Upstart 2025 Form 10-K

Source: https://www.sec.gov/Archives/edgar/data/1647639/000164763926000027/upst-20251231.htm
Published: 2026.
Measurement period: year ended 2025-12-31.
Evidence grade: A for platform automation.
Confidence: medium for market-level scoring.

Relevant values:

- 91% of loans on the Upstart platform were fully automated in 2025, with no human intervention by Upstart.
- Upstart-powered loans were also 91% fully automated in 2024.
- Transaction volume in 2025 was $11.004 billion.
- Transaction volume in 2025 was 1,497,149 loans.
- Transaction volume dollars increased 86% year over year.
- Transaction volume number of loans increased 115% year over year.
- The percentage fully automated metric was revised in Q4 2025 to include HELOCs; prior-period metrics were not recast because Upstart says the impact was immaterial.
- Upstart defines the metric as loans originated end-to-end with no human involvement required by the company, divided by transaction volume by number of loans.

Recommendation: use 91% only as a platform-level automation numerator. Do not score it as broad credit-market delegation without a market denominator.

### TransUnion Q4 2025 CIIR

Source: https://newsroom.transunion.com/q4-2025-ciir/
Published: 2026-02-19.
Measurement period: Q3/Q4 2025 with originations viewed one quarter in arrears.
Evidence grade: B for personal-loan denominator.
Confidence: medium.

Relevant values:

- Unsecured personal loan originations reached a record 7.2 million in Q3 2025.
- FinTech lenders held 42% share of unsecured personal loan originations.
- FinTech share was up from roughly one-third a year earlier.
- Total unsecured personal loan balances climbed to $276 billion in Q4 2025.
- 26.4 million consumers carried unsecured personal loan balances.
- Consumer-level 60+ days past due delinquency rose to 3.99% from 3.57% one year earlier.

Recommendation: use the 42% FinTech origination share as the denominator for a personal-loan fintech automation proxy. Do not apply it to credit cards, mortgages, auto, or commercial credit without an explicit product weighting rule.

### TransUnion Q1 2026 CIIR

Source: https://www.transunion.com/blog/q1-2026-consumer-credit-industry-insights-k-shaped-credit-market-reshaping-lending
Published: updated 2026-05-03.
Measurement period: Q1 2026 report.
Evidence grade: B for current credit-market context.
Confidence: medium.

Relevant values:

- U.S. credit outcomes are becoming more polarized, with super-prime consumers strengthening and non-prime borrowers facing rising pressure.
- Subprime and deep-subprime borrowers captured a larger share of bankcard and unsecured personal loan originations.
- Unsecured personal loan originations hit a new high, up more than 20% year over year.
- Lenders continue extending credit while using tighter risk controls.

Recommendation: use as market context and freshness signal. It does not refresh FinTech share directly from the public page.

### Bank of England / FCA AI in UK Financial Services 2024

Source: https://www.bankofengland.co.uk/report/2024/artificial-intelligence-in-uk-financial-services-2024
Published: 2024-11-21.
Measurement period: 2024 survey.
Evidence grade: B for financial-services AI governance and automation.
Confidence: medium.

Relevant values:

- Survey received 118 firm responses.
- 75% of firms were already using AI and 10% more planned to use AI over the next three years.
- 55% of all AI use cases had some degree of automated decision-making.
- 24% of automated-decision use cases were semi-autonomous.
- Only 2% of use cases were fully autonomous.
- 62% of AI use cases were rated low materiality and 16% high materiality.
- 46% of respondent firms reported only partial understanding of the AI technologies they use; 34% reported complete understanding.
- FCA's summary notes past Bank/FCA machine-learning surveys covered credit underwriting use cases.

Recommendation: use as bank/financial-services context and a guardrail against assuming full autonomy. It is not a credit-specific bank decisioning metric.

### Bank of England February 2026 AI Roundtables

Source: https://www.bankofengland.co.uk/minutes/2026/february/summary-of-ai-roundtables-feb-2026
Published: 2026-02-16.
Measurement period: late-2025 roundtables with regulated firms.
Evidence grade: B for regulated-firm deployment constraints.
Confidence: medium.

Relevant values:

- Roundtables included challenger banks and UK-focused larger banks, global systemically important banks, and insurers.
- Participants generally supported the PRA's principles-based regulatory framework for AI.
- Second-line risk functions continue to approach AI cautiously, which may delay deployment pipelines.
- Traditional model risk validation was described as hard to sustain as generative and agentic systems proliferate.
- The human-in-the-loop concept was challenged by the rise of agentic AI.
- Firms noted cross-jurisdiction regulatory fragmentation, procurement issues, data protection, and data quality as constraints.

Recommendation: use as 2026 deployment-friction context. It supports keeping bank AI credit decisioning conservative unless a direct bank credit underwriting source is found.

## Proposed Credit Source Lock

Proposed v2 scoring candidates:

| Indicator | Suggested role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| AI-Underwritten Personal Loan Proxy | score input | A/B | medium | Upstart 91% fully automated times TransUnion 42% FinTech personal-loan origination share equals 38.2% |
| FinTech Personal Loan Origination Share | denominator or context | B | medium | Strong product-specific market denominator; avoid double-counting if used in blended proxy |
| Bank AI Credit Decisioning | hold or low-confidence score input | B | low | BoE/FCA broad AI survey is not credit-specific; 2026 roundtables imply caution |
| UK Financial Services AI Credit Risk | display/context | B | medium | Useful governance and automation context; not part of current score |

Near-term decision:

- Do not use Upstart 91% raw as `AI-Underwritten Loan Volume`.
- Candidate blended proxy: `91% Upstart fully automated * 42% FinTech personal-loan origination share = 38.2%`.
- If the blended proxy is used, consider demoting `Fintech Lending Market Share` from scored input to denominator/context to avoid double counting.
- Keep `AI Credit Decisioning (Banks)` at hold until a credit-specific bank deployment metric is found.
- Document that the strongest available 2026 signal is for unsecured personal loans, not all consumer and commercial credit.
