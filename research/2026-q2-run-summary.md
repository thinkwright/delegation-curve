# 2026 Q2 Run Summary

Status: implemented as the current run in `seed/seed.json`.
Run ID: `2026-q2`.
Methodology version: `delegation-curve-v2`.
Published: 2026-05-16.
Data freshness label: `2026 Q2`.

## Result

The 2026 Q2 refresh yields a composite Delegation Curve score of **45.8**, down 0.2 points from the 2025 baseline score of 46.0.

This is not a finding that AI delegation declined broadly. The flat composite is mainly the result of a methodology correction in `algo-trade`: the prior 2025 score mixed market-automation proxies with AI-specific adoption. The v2 run narrows the scored inputs to:

```text
0.55 * 59.0 FX Electronic Trading Share
+ 0.45 * 15.0 Buy-Side AI Trade Execution Adoption
= 39.2
```

The 15.0 value is current internal-AI trade-execution adoption from Coalition Greenwich. The 39.0 current-plus-planned value is documented but not used for current scoring.

## Domain Scores

| Domain | 2025 baseline | 2026 Q2 | Change | Notes |
| --- | ---: | ---: | ---: | --- |
| content-mod | 93.0 | 96.1 | +3.1 | X demoted; YouTube automated flagging and TikTok DSA automated enforcement used |
| algo-trade | 68.4 | 39.2 | -29.2 | Adjusted from broad market automation toward defensible current AI execution adoption |
| code-gen | 46.6 | 48.2 | +1.6 | Adds Sonar committed-code estimate and METR value-share estimate |
| support | 45.9 | 40.5 | -5.4 | Renames resolution to cases handled by AI and adds production/maturity signals |
| credit | 28.9 | 30.0 | +1.1 | Uses Upstart automation times TransUnion fintech personal-loan denominator |
| medical-dx | 25.1 | 32.3 | +7.2 | Uses FDA live device count and AMA 2026 assistive diagnosis rate |
| legal-ai | 29.0 | 46.3 | +17.3 | Uses Thomson Reuters org GenAI adoption and Clio solo/small legal-work adoption |
| hire | 36.8 | 49.0 | +12.2 | Fixes indicator contract and uses ICIMS/Aptitude TA adoption and screening use-case values |
| education | 19.4 | 35.8 | +16.4 | Renames student indicator to schoolwork and uses Pew/OECD/Turnitin 2026 signals |

## Methodology Notes

- Historical runs are preserved as explicit snapshots: `legacy-2024`, `legacy-2025`, and `2026-q2`.
- v2 current views display only current scored indicators; retired v1 indicators remain in historical `indicator_observations`.
- `algo-trade` should be treated as a methodology-break domain. The new score is more conservative and more semantically honest.
- Some 2026 values remain survey-based or vendor-reported. The source-lock ledger keeps evidence grade and confidence notes for follow-up.
- The composite should be interpreted as a new v2 point on the curve, not a pure like-for-like data refresh.

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
