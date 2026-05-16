# 2026 Q2 Run Summary

Status: implemented as the current run in `seed/seed.json`.
Run ID: `2026-q2`.
Methodology version: `delegation-curve-v2`.
Published: 2026-05-16.
Data freshness label: `2026 Q2`.

## Result

The 2026 Q2 refresh yields a composite Delegation Curve score of **45.8**, up 8.1 index points from the 2025 comparable estimate of 37.7.

The March 2026 published score of 46.0 is retained as an archived run, but it is no longer used as the public prior because it was produced under the original scoring series. The public curve now compares `2026-q2` against `current-method-2025`, a recalculated estimate under the 2026 scoring frame. See `research/2026-current-method-series.md`.

The largest construct correction is in `algo-trade`: the archived 2025 score mixed market-automation proxies with AI-specific adoption. The v2 run narrows the scored inputs to:

```text
0.55 * 59.0 FX Electronic Trading Share
+ 0.45 * 15.0 Buy-Side AI Trade Execution Adoption
= 39.2
```

The 15.0 value is current internal-AI trade-execution adoption from Coalition Greenwich. The 39.0 current-plus-planned value is documented but not used for current scoring.

## Domain Scores

| Domain | 2025 comparable | 2026 Q2 | Change | Notes |
| --- | ---: | ---: | ---: | --- |
| content-mod | 94.1 | 96.1 | +2.0 | X demoted; YouTube automated flagging and TikTok DSA automated enforcement used |
| algo-trade | 39.2 | 39.2 | +0.0 | Adjusted from broad market automation toward defensible current AI execution adoption |
| code-gen | 32.6 | 48.2 | +15.6 | Adds Sonar committed-code estimate and METR value-share estimate |
| support | 37.0 | 40.5 | +3.5 | Renames resolution to cases handled by AI and adds production/maturity signals |
| credit | 27.7 | 30.0 | +2.3 | Uses Upstart automation times TransUnion fintech personal-loan denominator |
| medical-dx | 21.7 | 32.3 | +10.6 | Uses FDA live device count and AMA 2026 assistive diagnosis rate |
| legal-ai | 29.0 | 46.3 | +17.3 | Uses Thomson Reuters org GenAI adoption and Clio solo/small legal-work adoption |
| hire | 38.1 | 49.0 | +10.9 | Fixes indicator contract and uses ICIMS/Aptitude TA adoption and screening use-case values |
| education | 18.0 | 35.8 | +17.8 | Renames student indicator to schoolwork and uses Pew/OECD/Turnitin 2026 signals |

## Methodology Notes

- Historical published runs are preserved as archived snapshots: `legacy-2024` and `legacy-2025`.
- The public curve uses the current-method series: `current-method-2024`, `current-method-2025`, and `2026-q2`.
- v2 current views display only current scored indicators; retired v1 indicators remain in historical `indicator_observations`.
- `algo-trade` remains the largest construct correction. The new score is more conservative and more semantically honest.
- Some 2026 values remain survey-based or vendor-reported. The source-lock ledger keeps evidence grade and confidence notes for follow-up.
- The composite should be interpreted as one point in the public current-method curve, with archived published scores available for audit.

## Regeneration Commands

```sh
PATH=/home/bran/code/curve/.tools/go/bin:$PATH make collect
./curve-snapshot-run \
  -seed seed/seed.json \
  -run-id 2026-q2 \
  -label "2026 Q2 source refresh" \
  -measurement-period "2026 Q2" \
  -measurement-year 2026 \
  -published-at 2026-05-16 \
  -methodology-version delegation-curve-v2 \
  -data-freshness "2026 Q2" \
  -replace
PATH=/home/bran/code/curve/.tools/go/bin:$PATH make generate
```
