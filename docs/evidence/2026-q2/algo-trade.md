# Algorithmic Trading 2026 Q2 Evidence Notes

Status: implemented for the 2026 Q2 source refresh.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `algo-trade` score: 39.2.

Configured scoring formula:

- `FX Electronic Trading Share`: 55% weight.
- `Buy-Side AI Trade Execution Adoption`: 45% weight.

Current observations:

- `FX Electronic Trading Share`: 59.0%, freshness April 2025.
- `Buy-Side AI Trade Execution Adoption`: 15.0%, freshness 2025.

The current score calculates as:

```text
0.55 * 59.0 + 0.45 * 15.0 = 39.2
```

The implemented 2026 Q2 contract deliberately narrows this domain away from broad market-automation proxies. BIS electronic FX execution remains an execution-automation anchor, while Coalition Greenwich current internal-AI trade-execution adoption keeps the score grounded in AI-specific workflow use.

## Extracted Candidate Sources

### SEC Market Structure Analytics and Algorithmic Trading Report

Sources:

- https://www.sec.gov/featured-topics/market-structure-analytics
- https://www.sec.gov/data-research/sec-markets-data/marketstructuredata-exchange
- https://www.sec.gov/files/marketstructure/research/algo_trading_report_2020.pdf
- https://www.cboe.com/insights/posts/2025-u-s-equities-year-in-review/

Published or updated: 2020-2026.
Measurement period: SEC MIDAS downloads through 2025 Q4; Cboe 2025 equity review.
Evidence grade: A for market-structure context.
Confidence: medium for automation context; low for direct algorithmic volume.

Relevant values and constraints:

- SEC staff describes U.S.-listed equity securities as having a primarily automated secondary-market structure.
- SEC staff also says nearly all major equity trading centers depend on automated systems and algorithms for market-structure functions.
- SEC Market Structure Analytics publishes MIDAS-derived datasets through 2025 Q4 for metrics such as trade-to-order volume, cancel-to-trade ratio, odd-lot activity, hidden rate, and hidden volume.
- Those SEC metrics are strong evidence that U.S. equity markets are automated and machine-paced, but they do not directly measure the share of volume executed by algorithmic trading strategies.
- Cboe's 2025 U.S. equities review reports 17.6 billion shares average daily volume, $1.1 trillion average daily notional value, and 50.6% TRF market share.
- Cboe also reports that 18.7% of TRF volume occurred on ATS platforms and 81.3% through principal dealers.
- TRF share and off-exchange share are market-structure signals, not direct algorithmic-execution shares.

Recommendation: keep SEC and Cboe as authoritative market-structure context. Do not refresh `US Equities Algo Volume` from them unless we define the indicator as an automation proxy rather than direct algorithmic volume.

### BIS 2025 Triennial Survey and FX Execution Analysis

Sources:

- https://www.bis.org/statistics/rpfx25_fx.htm
- https://www.bis.org/publ/qtrpdf/r_qt2512v.htm

Published: 2025-09-30 and 2025-12-08.
Measurement period: April 2025.
Evidence grade: A.
Confidence: high for electronic FX execution; medium for algorithmic or AI interpretation.

Relevant values and constraints:

- BIS reports global OTC FX turnover of $9.6 trillion per day in April 2025, up 28% from 2022.
- BIS execution analysis reports electronic trading accounted for 59% of global FX turnover in April 2025.
- BIS says the electronic share was virtually unchanged from the previous Triennial Survey.
- BIS execution categories measure how trades are executed, not whether AI makes trading decisions.

Recommendation: use BIS as the cleanest refreshable trading input. Rename `FX Algo Trading` to `FX Electronic Trading Share` if using the 59.0% value.

### Cboe Options Industry Reports

Sources:

- https://www.cboe.com/insights/posts/the-state-of-the-options-industry-2025/
- https://www.cboe.com/insights/posts/the-state-of-the-options-industry-q-1-2026/
- https://www.sec.gov/featured-topics/market-structure-analytics/research-analysis-market-structure

Published: 2026-01-22 and 2026-05-04.
Measurement period: 2025 and Q1 2026.
Evidence grade: B for this indicator.
Confidence: medium for options-market context; low for algorithmic or AI share.

Relevant values and constraints:

- Cboe reports 2025 U.S. listed options volume above 15.2 billion contracts, with 61 million contracts traded daily on average.
- Cboe reports Q1 2026 market-wide options average daily volume of 68.6 million contracts.
- Cboe reports Q1 2026 index options ADV of 6.1 million contracts and SPX ADV of 4.9 million contracts.
- Cboe reports Q1 2026 FLEX options volume of 1.9 million contracts daily.
- Cboe reports 2025 0DTE SPX options averaged 2.3 million contracts daily and made up 59% of SPX product volume.
- SEC has 2026 options-market-structure roundtable material, but that material supports market-structure context more than algorithmic share.

Recommendation: demote `Options Algo Volume` to context unless a source directly measures automated or algorithmic execution share in options. Current Cboe sources are timely but not semantically matched to the existing indicator.

### The TRADE 2025 Algorithmic Trading Survey

Source: https://www.thetradenews.com/wp-content/uploads/2025/04/Algo56-FINAL.pdf
Published: Q1 2025.
Measurement period: 2025 survey publication.
Evidence grade: B.
Confidence: medium for adoption context; low for direct score input.

Relevant values and constraints:

- The TRADE 2025 survey reports long-only algorithm users rated algorithmic trading providers at 6.00 on average, up from 5.81 in 2024.
- The survey describes algorithmic trading as moving beyond simple rules-based execution into more sophisticated models.
- Reported electronic-trading asset-class participation includes ETFs 58%, fixed income 36%, FX 35%, listed derivatives 39%, and crypto 3%.
- The survey is useful for qualitative momentum and provider-performance context, but not for a clean market-volume denominator.

Recommendation: keep as supporting evidence only unless the methodology and response base are converted into a clear adoption indicator.

### Coalition Greenwich Buy-Side AI in Equity Trading

Sources:

- https://www.greenwich.com/market-structure-technology/great-expectations-ai-equity-trading
- https://www.greenwich.com/node/156366
- https://www.greenwich.com/node/149831

Published: 2024-04-30 and 2025-07-08.
Measurement period: 2024 survey fieldwork for 2025 report.
Evidence grade: C for score input; B for directional context.
Confidence: medium.

Relevant values and constraints:

- Coalition Greenwich says 15% of North American buy-side equity traders already incorporated internal AI technologies into trade execution workflow.
- Another 24% planned to incorporate internal AI technologies in the next 12 months.
- The study explicitly excludes third-party tools such as algo wheels and broker or vendor analytics platforms.
- Nearly 80% of participants thought AI would have a significant effect on algorithm optimization.
- The linked report methodology says Coalition Greenwich interviewed 40 buy-side equity traders in North America from July through September 2024.
- A 2024 Coalition Greenwich blog using 90 buy-side traders across the U.S. and Europe found 10% already incorporating AI/ML in equity trading processes, 16% planning in the next year or two, and 74% not planning to do so.
- The 2024 blog cautions that respondents may interpret AI/ML broadly and may conflate their own AI use with sell-side or vendor tools.

Recommendation: replace the current `Institutional AI Adoption` value of 78.0. The defensible values are 15% current internal-AI use or 39% current-plus-planned internal-AI use. The 78% figure is closer to expected impact on algo optimization, not adoption.

## Implemented Algorithmic Trading Source Lock

Current v2 scoring inputs and retained candidates:

| Indicator | Role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| U.S. Equities Automated Market Structure | context or score only after redefinition | A | medium | SEC and Cboe are strong on automation context but not direct algo volume |
| FX Electronic Trading Share | score input | A | high | BIS 2025 supports 59.0% electronic trading share |
| Options Market Electronic Structure Context | context | B | medium | Cboe gives fresh options volume but no algorithmic share |
| Buy-Side AI Trade Execution Adoption | score input | C | medium | Use 15% current internal-AI trade-execution adoption, not 39% current-plus-planned or 78% expected impact |
| EU Equity Algo Volume | hold or drop | D | low | Display-only 2021 value remains stale |

Implementation decision:

- Do not treat 2026 market-structure volume sources as direct AI delegation evidence.
- Use BIS 59.0% only under the renamed `FX Electronic Trading Share` indicator.
- Use Coalition Greenwich 15% current internal-AI trade-execution adoption, not the 39% current-plus-planned value or 78% expected-impact figure.
- Retire broad U.S. equities, options, and stale EU equity proxies from the current score unless a direct automated-execution source is found.
- Keep a clear note that this domain remains partly market automation and partly AI-specific trade execution.
