# 2026 Delegation Curve Source Lock Plan

Status: preparation artifact, not a score update.
Date prepared: 2026-05-16.

## Objective

Prepare the next Delegation Curve research run so it can incorporate 2026 evidence without silently changing what the composite means. The current public curve is best understood as a 2024-2025 measurement snapshot, with a few early 2026 references. The next run should explicitly distinguish:

- observed delegated decisions,
- workflow penetration,
- AI-attributable output or value share,
- ordinary adoption,
- capability benchmarks or contextual evidence.

The source set should be locked before score edits begin.

## Current Dataset Audit

Local state:

- Original score snapshot generated at `2026-03-02T08:36:54Z`; the seed file has since been rewritten only to materialize run history.
- composite last updated `2026-03-02`, `data_year: 2025`.
- The current 2024 and 2025 curve points are now materialized as explicit `analysis_runs` so the next refresh can append a new run instead of overwriting the historical baseline.
- 9 domains, 46 displayed sub-indicators, 42 displayed data-source rows.
- Active scoring config uses 32 configured indicators, but only 31 are present by exact name in the current seed.
- 15 displayed indicators are currently not included in scoring.
- Only `VS Code Marketplace` is a live automated collector. Almost all other collectors are manual stubs resolved by `seed/overrides.yaml` or cached seed values.

Freshness distribution across displayed indicators:

| Freshness label | Count |
| --- | ---: |
| 2021 | 1 |
| 2023 / H2 2023 | 2 |
| 2024 / 2024 quarter labels | 10 |
| 2025 / 2025 quarter labels | 31 |
| 2026 / exact 2026 date | 2 |

Important local drift to fix before scoring:

- `hire` config expects `Orgs Using AI Screening`, but the seed displays `Orgs Using AI in Recruiting`. This causes the configured 40% hiring indicator to be missing and the score to reweight around the remaining two configured indicators.
- `content-mod`, `algo-trade`, `credit`, `medical-dx`, `legal-ai`, `hire`, and `education` display additional indicators that are not scored.
- The methodology page says status tiers are nominal `<30`, elevated `30-60`, autonomous `>60`, but the code classifies nominal `<40`, elevated `40-74.9`, autonomous `>=75`.
- `DataFreshness` is now seed-driven via `delegation.composite.data_freshness`.

## Safe Refresh Workflow

Do not update scores directly in `seed/seed.json` without creating or updating an explicit run snapshot.

1. Lock the existing baseline if it is not already materialized:

   ```sh
   make snapshot-baseline
   ```

2. Refresh source values in `seed/overrides.yaml` or collectors, keeping citations and evidence grades in `research/2026-source-ledger.csv`.
3. Run the collector and generator.
4. Append the new current run after the refreshed scores are present:

   ```sh
   ./curve-snapshot-run \
     -seed seed/seed.json \
     -run-id 2026-q2 \
     -label "2026 Q2 source refresh" \
     -measurement-period "2026 Q2" \
     -measurement-year 2026 \
     -published-at YYYY-MM-DD \
     -methodology-version delegation-curve-v2 \
     -data-freshness "2026 Q2"
   ```

5. Regenerate Parquet and rebuild the frontend/server. The visualizations should then show both the current March 2026 publication baseline and the new 2026 refresh point.

## Evidence Grades

Use these grades when deciding whether a source can enter the composite.

| Grade | Meaning | Use in scoring |
| --- | --- | --- |
| A | Direct observed delegated decisions or task resolutions | Preferred score input |
| B | Workflow penetration or operational deployment | Score input if no direct outcome exists |
| C | AI-attributable output/value/productivity share | Useful for code/support; needs caveats |
| D | Adoption, sentiment, capability benchmark, or context | Context only unless explicitly converted |

## Cross-Cutting 2026 Sources

### METR AI Usage Survey

URL: https://metr.org/blog/2026-05-11-ai-usage-survey/

Use for `code-gen` and possibly a broader technical-work lens. METR surveyed technical workers in February-April 2026 and reports self-assessed value uplift, including retrospective March 2025 and March 2026 comparisons. This is not direct committed-code measurement, but it is closer to value delegation than ordinary tool adoption.

Proposed conversion for exploration:

```text
AI-attributable value share = (multiplier - 1) / multiplier
```

Example: a `2x` value multiplier implies `50%` AI-attributable value share. Keep this as a separate indicator rather than replacing all code-gen measures.

### Stanford HAI 2026 AI Index

URL: https://hai.stanford.edu/ai-index/2026-ai-index-report

Use as a high-quality secondary source and source-discovery map. It has strong 2026 relevance for economy, education, medicine, technical capability, and productivity. Do not treat capability benchmarks as delegation scores. Treat them as context unless they describe real deployment or workflow outcomes.

Relevant extracts for source selection:

- Organizational AI adoption reached 88%; generative AI is used in at least one business function by 70% of organizations; agent deployment remains single-digit across most functions.
- Productivity studies cited by AI Index report gains of 14-15% in customer support, 26% in software development, and 50% in marketing output.
- Four out of five U.S. high school and college students use AI for schoolwork.
- FDA authorized 258 AI medical devices in 2025.
- AI agents rose to 66.3% on OSWorld, but this is capability context, not delegated social decision-making.

## Domain Recommendations

### content-mod

Current state: mostly fresh through Q3/H1 2025, but `DSA Cross-Platform Automated Rate` is displayed and ignored by scoring.

Recommended lock:

- Keep Meta, YouTube/Google, TikTok, and X as direct platform transparency sources.
- Promote DSA only if we can reproduce a stable, cross-platform automated-decision extraction. Otherwise keep DSA as context/display only.
- Refresh all platform values to the latest available 2025 or 2026 reporting period.

Candidate sources:

- Meta Transparency Center Community Standards Enforcement Reports.
- YouTube Community Guidelines Enforcement report and data download.
- TikTok Community Guidelines Enforcement and DSA reports. TikTok has reported over 85% of removed content identified and removed by automation and Q4 2025 automated detection counts.
- X Global Transparency Reports. Current public result located for H2 2024; needs check for H1/H2 2025 availability.
- EU DSA Transparency Database or platform DSA transparency reports.

Decision: keep current source family; investigate DSA reproducibility before scoring.

### algo-trade

Current state: one displayed signal is from 2021. Direct AI-vs-algorithmic trading measurement remains weak.

Recommended lock:

- Keep SEC/FINRA market structure and BIS FX as anchors.
- Keep Cboe for options market structure/volume, but avoid converting ordinary options volume into AI delegation.
- Keep a survey source such as Greenwich/The TRADE for institutional algo adoption if the methodology is available.
- Do not import generic "AI trading market size" vendor estimates into the composite.

Candidate sources:

- BIS 2025 Triennial Survey and FX execution analysis.
- SEC market structure and algorithmic trading materials.
- Cboe 2025 and Q1 2026 options industry reports.
- The TRADE 2025 Algorithmic Trading Survey as a possible survey/proxy source.

Decision: targeted source refresh only; no obvious high-fidelity 2026 replacement found yet.

### code-gen

Current state: high-value domain for a 2026 refresh. Existing sources mix acceptance, adoption, OSS proxy, and install inventory.

Extraction note: see `research/2026-evidence/code-gen.md`.

Recommended lock:

- Add METR as `Technical Work AI Value Share` or `AI-Attributable Technical Output Value`.
- Add Sonar 2026 as `AI-Generated or Assisted Committed Code` if we accept survey-based committed-code estimates.
- Keep Stack Overflow 2025 as adoption/daily-use context; do not use it as output share unless the question maps directly to output.
- Keep VS Code Marketplace as a weak but continuously refreshable stock metric.
- Keep GitClear only if the report can be refreshed and the definition remains stable.

Candidate sources:

- METR 2026 AI usage survey.
- Sonar 2026 State of Code Developer Survey.
- Stack Overflow 2025 Developer Survey AI section.
- DORA 2025 State of AI-assisted Software Development.
- VS Code Marketplace API.

Decision: add METR and consider Sonar; reduce conceptual weight on install counts.

### support

Current state: all displayed indicators are 2025, survey/proxy heavy.

Recommended lock:

- Keep Zendesk, Intercom, and Salesforce as core source families.
- Prefer resolved-case or deflection-rate measures over broad "using AI" metrics.
- Add Gartner or AI Index productivity figures as context only unless tied to support resolution/deflection.
- Consider Intercom 2026 as a stronger operational maturity source, but validate whether it provides a direct resolution or mature deployment percentage.

Candidate sources:

- Intercom 2026 Customer Service Transformation Report.
- Zendesk 2026 CX Trends report.
- Salesforce customer service trends/state-of-service materials.
- Gartner 2026 support workforce/AI implementation survey.
- Stanford HAI 2026 economy chapter for productivity context.

Extraction note: `research/2026-evidence/support.md`.

Decision: keep support as a high-priority score-refresh domain, but rename `AI Resolution Rate` to `Cases Handled by AI` unless a true autonomous-resolution source is found. Add Sinch as production/reliability context and avoid counting rollback-prone deployment as positive impact without a methodology decision.

### credit

Current state: one bank decisioning signal is H2 2024, one UK source is displayed but unscored.

Recommended lock:

- Keep Upstart/fintech filings as the best direct automation source.
- Keep TransUnion CIIR for personal-loan market denominator, but avoid overextending personal-loan fintech share to all credit without an explicit blending rule.
- Keep McKinsey/OCC/BoE/FCA as bank-adoption context and guardrail; banks appear more cautious than fintechs in high-stakes decisioning.
- Make the blend formula explicit in the ledger.

Candidate sources:

- Upstart 2025 Form 10-K: 91% of loans on platform fully automated in 2025, with definition revised in Q4 2025.
- TransUnion Q1 2026 CIIR and Q4 2025 CIIR.
- McKinsey 2026 risk/banking articles for proof-of-concept and adoption context.
- Bank of England February 2026 AI roundtables and FCA/BoE AI survey materials.

Extraction note: `research/2026-evidence/credit.md`.

Decision: lock Upstart plus TransUnion, but do not use the raw 91% Upstart platform automation rate as a broad market value. The current candidate is `91% * 42% = 38.2%` for a personal-loan fintech automation proxy. If that blended proxy is scored, demote `Fintech Lending Market Share` to denominator/context to avoid double counting. Treat bank AI adoption as cautious context unless a credit-decisioning deployment measure is found.

### medical-dx

Current state: many displayed medical signals are stale or unscored; active scoring uses four indicators.

Recommended lock:

- Keep FDA AI-enabled device list as a primary source, but use it as capacity/deployment proxy, not actual clinical diagnosis share.
- Keep radiology, diagnosis, and pathology adoption as separate workflow penetration indicators.
- Treat clinical-note automation as medical workflow context, not diagnosis delegation.
- Treat benchmark diagnostic systems as context unless deployed clinically.

Candidate sources:

- FDA AI-Enabled Medical Device List.
- Stanford HAI 2026 medicine chapter: 258 FDA AI medical devices authorized in 2025; broad clinical note adoption; diagnostic benchmark/system context.
- AMA physician AI sentiment/adoption surveys.
- RSNA/radiology adoption surveys.
- CAP/pathology digital pathology adoption sources.

Extraction note: `research/2026-evidence/medical-dx.md`.

Decision: refresh FDA from the direct CSV and use AMA 2026 `assistive diagnosis` as the best near-term diagnosis-use candidate. Keep radiology pending until a comparable deployment survey is locked, keep pathology low-confidence, and keep note automation plus diagnostic benchmarks out of the diagnosis score.

### legal-ai

Current state: mostly 2024-2025 adoption measures; direct delegated legal decision-making remains low and hard to observe.

Recommended lock:

- Keep BigLaw and solo/small adoption split.
- Add Thomson Reuters 2026 as a stronger cross-professional and legal-sector anchor.
- Keep document review/TAR as the most directly delegated legal workflow.
- Do not score legal reasoning benchmark performance as delegation.

Candidate sources:

- Thomson Reuters 2026 AI in Professional Services Report.
- Clio 2026 Legal Trends for Solo and Small Law Firms.
- ABA Legal Technology Survey / TAR materials.
- Law360 Pulse 2026 AI Survey if accessible and methodology is adequate.

Decision: refresh with Thomson Reuters and Clio/ABA; keep document review as direct workflow anchor.

### hire

Current state: active scoring has a name mismatch and reweights around two configured indicators. Several displayed hiring indicators are not scored.

Recommended lock:

- Fix the indicator contract before collection: decide whether the first scored indicator is `Orgs Using AI in Recruiting` or `Orgs Using AI Screening`.
- Use one org-level adoption metric, one application/screening metric, and one platform-reach metric.
- Add 2026 HR/recruiting reports where they measure screening, scheduling, or candidate communication automation.

Candidate sources:

- SHRM 2026 State of AI in HR.
- ICIMS/Aptitude 2026 AI adoption in talent acquisition report.
- SHRM 2025 Talent Trends for resume-screening specifics if no 2026 equivalent is available.
- Vendor platform disclosures such as HireVue only for reach, with lower confidence.

Decision: high-priority cleanup domain. Do not collect until naming and inclusion choices are fixed.

### education

Current state: one 2026 Turnitin-style signal is displayed, but core scoring is mostly 2025.

Recommended lock:

- Replace or supplement `Students Using AI Tutors` with a broader student schoolwork-use indicator if the scope is "AI influence in education" rather than tutoring specifically.
- Keep AI grading/assessment as a separate direct delegation indicator.
- Keep faculty AI teaching use as workflow penetration.
- Treat AI detection flags as evidence of student output, not delegated assessment, unless tied to institutional enforcement decisions.

Candidate sources:

- Stanford HAI 2026 education chapter: four out of five U.S. high school and college students use AI for schoolwork.
- EDUCAUSE 2026 Students and Technology Report.
- EDUCAUSE 2026 Impact of AI on Work in Higher Education.
- Turnitin 2026 AI writing report/data.
- OECD TALIS 2024 / OECD 2026 Digital Education Outlook for teacher AI use.
- Pew 2026 teen AI survey as K-12 context.

Extraction note: `research/2026-evidence/education.md`.

Decision: rename or split the student indicator before scoring. Use Pew for K-12 schoolwork depth, HAI for broad cross-level context, OECD for teacher workflow, and keep grading plus AI-written-output signals separate until direct assessment-automation usage is found.

## Proposed Source Lock Gates

Before changing `seed/seed.json`, complete these gates:

1. One row per candidate source in `research/2026-source-ledger.csv`.
2. For every scored indicator, record `source_url`, `published_at`, `measurement_period`, `evidence_grade`, `confidence`, and `included_in_score`.
3. For each displayed-but-unscored indicator, explicitly choose `score`, `display_only`, or `drop`.
4. Fix the hiring indicator-name mismatch.
5. Align methodology copy with code thresholds or change the thresholds deliberately.
6. Decide whether `DataFreshness` is derived from source ledger dates rather than hardcoded.
7. Recompute the current 2025 score from current locked sources before adding 2026 values, to separate methodology drift from data drift.

## Near-Term Worklist

1. Build a source-ledger schema into the repo, either as CSV/YAML or by extending seed metadata.
2. Populate exact latest URLs and extracted values for each locked source.
3. Add `included_in_score` and `source_url` metadata to data export so public users can distinguish scored inputs from contextual indicators.
4. Run a dry recomputation with the same current values but cleaned source contracts.
5. Only then apply 2026 refreshed values and compare deltas.
