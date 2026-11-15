# Restatement And Correction Policy

The Delegation Curve is meant to be one continuous measurement. Changing how it is measured can make a real change look like noise, or noise look like a trend, so restatements are rare, run in shadow first, and are always disclosed.

## Kinds Of Change

| Change | Examples | Treatment |
| --- | --- | --- |
| **Source refresh** | A new edition of the same report, measuring the same construct and population | Ordinary data update in the next run |
| **Correction** | A value that does not match its cited source, a transcription error, a wrong reporting period | Fixed in the next run and logged below with its effect on the composite; earlier published runs are not edited |
| **Methodology change** | New or retired indicators, a different construct or population (for example an international sample replaced by a U.S. one), normalization, weights | Restatement, under the rules below |

## Rules For Methodology Changes

1. **At most once a year.** Methodology changes are introduced in the first run of a calendar year.
2. **Shadow first.** A proposed method is computed alongside the published one for at least one run, with results in the run record, before it becomes the headline method.
3. **Restate the whole public series.** Every public prior point is recomputed under the new method, with indicator-level observations for each restated run, so a restated point can be audited the same way as a current one.
4. **Keep the originals.** Previously published runs stay in the dataset as archived runs (`is_public_series: false`) and are listed in the log with their original and restated values.
5. **Say what moved.** The run that introduces a change reports how much of the headline movement comes from the method rather than the data.

## Log

With all 2026 Q2 corrections applied, the 2026 Q2 composite would have been 46.2 instead of the published 45.8. The published run is kept as released.

| Date | Change | Kind | Runs affected | Published | Restated or corrected | Notes |
| --- | --- | --- | --- | ---: | ---: | --- |
| 2026-05-16 | `delegation-curve-v2`: revised source contracts, algorithmic trading construct, five-input code-gen method | Methodology | 2024, 2025 | 39.7, 46.0 | 31.5, 37.7 | 2025 was restated from prior-period values under the v2 contracts; 2024 is a continuity backcast. Both restatements predate rule 3 and carry domain-level scores only. See [current-series.md](current-series.md). |
| 2026-10-17 | YouTube Automated Flagging: the 2026 Q2 input (99.5%) was the Jul-Sep 2023 comment-removal share; the Jul-Sep 2025 video share was 97.9% | Correction | 2026 Q2 content moderation | 96.1 | 95.3 | Together with the Meta correction below; composite effect -0.08. Fixed in 2026 Q4. |
| 2026-10-17 | Meta Automated Detection: the 2026 Q2 input (95.2%) could not be reproduced from Meta's data; the documented method gives 94.5% for Q3 2025 | Correction | 2026 Q2 content moderation | (above) | (above) | Method documented in [docs/evidence/2026-q4/content-mod.md](../evidence/2026-q4/content-mod.md). |
| 2026-10-17 | AI Credit Decisioning (Banks): the 2026 Q2 input (20%) measured generative-AI use cases among 24 North American credit-risk chiefs (October 2023), not credit-decisioning deployment; replaced by the ECB 2024 share of euro-area significant banks using AI for credit scoring (30%) | Correction | 2026 Q2 credit | 30.0 | 34.5 | Composite effect +0.45. Fixed in 2026 Q4. |
