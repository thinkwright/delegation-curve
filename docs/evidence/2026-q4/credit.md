# Credit And Lending: 2026 Q4 Evidence Notes

Run `2026-q4`, method `delegation-curve-v2`. Prior notes: [2026 Q2](../2026-q2/credit.md).

Domain score 30.0 → **34.5**, entirely from correcting the bank credit-decisioning input.

## Lock Decisions

| Indicator | Action | Value | Source | Window | Grade / confidence |
| --- | --- | ---: | --- | --- | --- |
| AI-Underwritten Personal Loan Proxy | hold (numerator refreshed) | 38.2% | [Upstart Form 10-Q, Q2 2026](https://www.sec.gov/Archives/edgar/data/1647639/000164763926000063/upst-20260630.htm) × [TransUnion Q4 2025 CIIR](https://newsroom.transunion.com/q4-2025-ciir/) | Q3 2025 (denominator) to Q2 2026 (numerator) | A / medium |
| AI Credit Decisioning (Banks) | correct | 30% | [ECB Banking Supervision newsletter, 2025-11-20](https://www.bankingsupervision.europa.eu/press/supervisory-newsletters/newsletter/2025/html/ssm.nl251120_1.en.html) | 2024 supervisory data | B / medium-low |

## Notes

- **Personal-loan proxy.** Upstart reports 91% of loans fully automated in both Q1 and Q2 2026, unchanged from FY2025. In Q2 2026 Upstart excluded its new revolving credit line product from the definition; the restated 2025 comparisons match the originally reported figures, so the change appears immaterial. TransUnion's Q2 2026 release does not state the fintech share of unsecured personal-loan originations, and the report PDF could not be retrieved, so the 42% Q3 2025 share is held. The product is unchanged at 38.2%.
- **Bank credit decisioning.** The 2026 Q2 input (20%) was the share of 24 North American credit-risk chiefs reporting at least one generative-AI use case at an October 2023 roundtable, which is not deployment in credit decisions. The replacement is the ECB's supervisory figure: 30% of euro-area significant institutions used AI for credit scoring in 2024 (14% in 2023). Caveats: it covers large euro-area banks rather than U.S. banks or a share of decisions; the ECB notes that conventional statistical models may be reported as AI and that humans still take most credit decisions outside small retail loans.

## Candidates For Methodology v3

- Upstart's 91% fully automated share is a direct, single-platform measure of delegated decisions.
- The [CSBS 2026 annual survey](https://www.csbs.org/2026-csbs-annual-survey) of U.S. community banks (330 banks) reports 13% offering automated loan underwriting; it is not AI-specific and fits as a guardrail.
