# 2026 Delegation Curve Source Lock Plan

Status: implemented for the 2026 Q2 run; public-series handling is recorded in `research/2026-current-method-series.md`.
Date prepared: 2026-05-16.

## Objective

Record the source lock and implementation decisions for the 2026 Q2 Delegation Curve run. The March 2026 published curve is best understood as a 2024-2025 measurement snapshot, with a few early 2026 references. The 2026 Q2 run explicitly distinguishes:

- observed delegated decisions,
- workflow penetration,
- AI-attributable output or value share,
- ordinary adoption,
- capability benchmarks or contextual evidence.

This file began as a pre-run source-lock plan and now also records the implemented decisions. The run summary is in `research/2026-q2-run-summary.md`.

## Dataset Audit

Pre-refresh local state:

- Original score snapshot generated at `2026-03-02T08:36:54Z`; the seed file has since been rewritten only to materialize run history.
- composite last updated `2026-03-02`, `data_year: 2025`.
- 9 domains, 46 displayed sub-indicators, 42 displayed data-source rows.
- Active scoring config used 32 configured indicators, but only 31 were present by exact name in the seed.
- 15 displayed indicators were not included in scoring.

Implemented 2026 Q2 state:

- Current run ID: `2026-q2`.
- Composite score: 45.8.
- Prior public comparison: 37.7 from `current-method-2025`.
- Public movement: +8.1 index points.
- Data freshness label: `2026 Q2`.
- Public analysis runs: `current-method-2024`, `current-method-2025`, and `2026-q2`.
- Archived publication baselines: `legacy-2024` and `legacy-2025`.
- 9 domains, 31 displayed current sub-indicators, 29 displayed data-source rows.
- Current displayed sub-indicators are the scored v2 indicators; retired v1 indicators remain in historical `indicator_observations`.
- Only `VS Code Marketplace` is a live automated collector. Almost all other collectors are manual stubs resolved by `seed/overrides.yaml` or cached seed values.

Freshness distribution across current displayed indicators:

| Freshness label | Count |
| --- | ---: |
| 2024 / H2 2024 | 3 |
| 2025 / 2025 quarter labels | 13 |
| 2025/2026 | 2 |
| 2026 / exact 2026 date | 13 |

Resolved local drift in the implemented run:

- The hiring indicator contract now uses `Orgs Using AI in Talent Acquisition`, `AI Screening Use Case Adoption`, `Broad AI Across Hiring Processes`, and `AI Assessment Platform Reach`.
- Current domain pages display only current scored v2 indicators.
- The methodology page status tiers match the code: nominal `<40`, elevated `40-74.9`, autonomous `>=75`.
- `DataFreshness` is seed-driven via `delegation.composite.data_freshness`.
- Code-gen no longer scores Stack Overflow daily use; it uses the five-input output-share, value-share, workflow-reliance, agentic-delegation, and tool-reach method.

## Refresh Workflow

The 2026 Q2 run followed this pattern, and future score updates should keep the same separation between source refresh, scoring, run snapshotting, and generated artifacts. Do not update scores directly in `seed/seed.json` without creating or updating an explicit run snapshot.

1. Lock the existing baseline if it is not already materialized:

   ```sh
   make snapshot-baseline
   ```

2. Refresh source values in `seed/overrides.yaml` or collectors, keeping citations and evidence grades in `research/2026-source-ledger.csv`.
3. Run the collector and generator.
4. Append the new current run after the refreshed scores are present. Example from the 2026 Q2 run:

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

5. Regenerate Parquet and rebuild the frontend/server. The visualizations should show the current-method public curve, while archived publication baselines remain available in the run data for audit.

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

Implemented state: current v2 scoring uses Meta Q3 2025 automated detection, YouTube Jul-Sep 2025 automated flagging, and TikTok H2 2025 automated enforcement. X and cross-platform DSA data are retained as context, not displayed current indicators.

Locked approach:

- Keep Meta, YouTube/Google, and TikTok as direct platform transparency score sources.
- Keep X on hold until a comparable 2025/2026 global report is found.
- Promote DSA only if we can reproduce a stable, cross-platform automated-decision extraction. Otherwise keep DSA as context/display only.

Candidate sources:

- Meta Transparency Center Community Standards Enforcement Reports.
- YouTube Community Guidelines Enforcement report and data download.
- TikTok Community Guidelines Enforcement and DSA reports. TikTok has reported over 85% of removed content identified and removed by automation and Q4 2025 automated detection counts.
- X Global Transparency Reports. Current public result located for H2 2024; needs check for H1/H2 2025 availability.
- EU DSA Transparency Database or platform DSA transparency reports.

Extraction note: `research/2026-evidence/content-mod.md`.

Decision: use Meta, YouTube, and TikTok as the current scored source family, with indicator names matched to source semantics. YouTube is scored as automated flagging, TikTok is scored using EU DSA automated enforcement, and Meta carries a 2025 methodology-break note. X stays on hold until a comparable 2025 global report is found. Promote cross-platform DSA only after tokened extraction and run-level snapshots exist.

### algo-trade

Implemented state: current v2 scoring uses BIS 2025 FX electronic trading share and Coalition Greenwich current internal-AI trade-execution adoption. Broad U.S. equities, options, and stale EU equity proxies are demoted from the current score.

Locked approach:

- Keep BIS FX as the scored market-execution automation anchor.
- Keep Coalition Greenwich as the narrow AI-specific trade-execution adoption source.
- Keep SEC/FINRA market structure as context.
- Keep Cboe for options market structure/volume, but avoid converting ordinary options volume into AI delegation.
- Do not import generic "AI trading market size" vendor estimates into the composite.

Candidate sources:

- BIS 2025 Triennial Survey and FX execution analysis.
- SEC Market Structure Analytics and algorithmic trading materials.
- Cboe 2025 and Q1 2026 options industry reports.
- The TRADE 2025 Algorithmic Trading Survey as a possible survey/proxy source.
- Coalition Greenwich 2025 buy-side AI in equity trading as a narrow AI-execution adoption source.

Extraction note: `research/2026-evidence/algo-trade.md`.

Decision: targeted source refresh plus methodology cleanup. BIS is scored as `FX Electronic Trading Share`, not as AI-specific trading. SEC and Cboe remain market-automation context unless a direct algorithmic-execution share is found. The prior `Institutional AI Adoption` value appeared to be expected impact on algorithm optimization rather than adoption; it is replaced with Coalition Greenwich 15% current internal-AI trade-execution adoption. Options and stale EU equity signals are demoted unless better direct sources are found.

### code-gen

Implemented state: current v2 scoring is 48.1 under a five-input method covering output share, technical-work value share, workflow reliance, agentic task delegation, and tool ecosystem reach. Stack Overflow daily use is retired from scoring and retained only as bridge adoption context.

Extraction note: see `research/2026-evidence/code-gen.md`.

Locked approach:

- Keep METR as `Technical Work AI Value Share`.
- Blend Sonar 2026 and GitLab 2026 as `AI-Generated Code Output Share`.
- Retire Stack Overflow 2025 from scoring and keep it as bridge adoption context only.
- Blend DORA 2025 and JetBrains AI Pulse 2026 as `AI Workflow Reliance`.
- Blend the arXiv GitHub coding-agent adoption study with Anthropic's fully delegated task range as `Agentic Task Delegation`.
- Use Anthropic's 2026 Agentic Coding Trends report to keep AI-assisted work separate from fully delegated tasks.
- Use Anthropic Claude Code, Alphabet internal-code claims, and Augment's AI-native survey as frontier benchmarks or narrative callouts, not as population score inputs.
- Use Harness, Lightrun, CircleCI, Faros, and AI-code technical-debt research as validation and quality guardrails rather than positive score inputs.
- Keep VS Code Marketplace as a weak but continuously refreshable stock metric.
- Keep GitClear only if the report can be refreshed and the definition remains stable.

Candidate sources:

- METR 2026 AI usage survey.
- Sonar 2026 State of Code Developer Survey.
- GitLab 2026 Global DevSecOps Report.
- Stack Overflow 2025 Developer Survey AI section.
- DORA 2025 State of AI-assisted Software Development.
- JetBrains AI Pulse 2026.
- Agentic Much? Adoption of Coding Agents on GitHub.
- Anthropic 2026 Agentic Coding Trends Report.
- Anthropic Claude Code product and documentation.
- Alphabet Q3 2025 earnings-call code-generation benchmark.
- Augment Code State of AI-Native Engineering 2026.
- Harness State of Engineering Excellence 2026.
- Lightrun State of AI-Powered Engineering 2026.
- CircleCI 2026 State of Software Delivery.
- Faros AI Engineering Report 2026.
- Debt Behind the AI Boom arXiv study.
- VS Code Marketplace API.

Decision: update the current code-gen contract to a five-input frame: output share from Sonar plus GitLab, technical-work value share from METR, workflow reliance from DORA plus JetBrains, agentic task delegation from Anthropic's assisted-versus-fully-delegated split plus the arXiv GitHub trace study, and low-weight tool ecosystem reach from VS Code Marketplace. The resulting `code-gen` score is 48.1, essentially unchanged from the pre-refresh 48.2, because high workflow reliance is counterbalanced by still-lower evidence for fully delegated agentic tasks. Delivery and reliability sources should be used as guardrails or uncertainty modifiers so the curve does not confuse generated-code volume with net delegated work. Implementation note: see `research/2026-code-gen-method-prototype.md`.

### support

Implemented state: current v2 scoring uses Salesforce current cases handled by AI, an existing bot deflection source, Sinch production AI customer communications agents, and Intercom mature support deployment. Survey/proxy limitations remain explicit.

Locked approach:

- Keep Salesforce and Intercom as core source families.
- Keep Zendesk or comparable operational telemetry as a future deflection-refresh target.
- Prefer resolved-case or deflection-rate measures over broad "using AI" metrics.
- Add Gartner or AI Index productivity figures as context only unless tied to support resolution/deflection.
- Keep Sinch production and rollback evidence as a reliability caveat, not an uncaveated positive adoption measure.

Candidate sources:

- Intercom 2026 Customer Service Transformation Report.
- Zendesk 2026 CX Trends report.
- Salesforce customer service trends/state-of-service materials.
- Sinch AI Production Paradox.
- Gartner 2026 support workforce/AI implementation survey.
- Stanford HAI 2026 economy chapter for productivity context.

Extraction note: `research/2026-evidence/support.md`.

Decision: rename `AI Resolution Rate` to `Cases Handled by AI`, add Sinch as production/reliability context, and keep rollback-prone deployment caveated. Avoid projected 2027 values in the current score.

### credit

Implemented state: current v2 scoring uses a blended Upstart-plus-TransUnion personal-loan fintech automation proxy and a conservative bank AI credit-decisioning signal. UK financial-services AI credit-risk data remains context.

Locked approach:

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

Decision: lock Upstart plus TransUnion, but do not use the raw 91% Upstart platform automation rate as a broad market value. Score `91% * 42% = 38.2%` as a personal-loan fintech automation proxy. Treat `FinTech Personal Loan Origination Share` as the denominator inside that proxy, not as an independent scored indicator. Treat bank AI adoption as cautious context unless a credit-decisioning deployment measure is found.

### medical-dx

Implemented state: current v2 scoring uses the FDA AI-enabled device count, radiology/imaging AI adoption, AMA assistive diagnosis use, and pathology AI adoption. Clinical-note automation and diagnostic benchmarks remain context.

Locked approach:

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

Decision: refresh FDA from the direct CSV, use AMA 2026 `assistive diagnosis` as the diagnosis-use score input, use KLAS as the current radiology/imaging adoption score input, keep pathology low-confidence, and keep note automation plus diagnostic benchmarks out of the diagnosis score.

### legal-ai

Implemented state: current v2 scoring uses Thomson Reuters legal-organization GenAI adoption, Clio solo/small-firm AI legal-work adoption, and ABA AI-assisted document review. Direct delegated legal decision-making remains hard to observe.

Locked approach:

- Replace the old BigLaw-specific adoption label with a broader legal-organization GenAI adoption indicator.
- Keep solo/small adoption split from broad legal-organization adoption.
- Add Thomson Reuters 2026 as a stronger cross-professional and legal-sector anchor.
- Keep document review/TAR as the most directly delegated legal workflow.
- Do not score legal reasoning benchmark performance as delegation.

Candidate sources:

- Thomson Reuters 2026 AI in Professional Services Report.
- Clio 2026 Legal Trends for Solo and Small Law Firms.
- ABA Legal Technology Survey / TAR materials.
- Law360 Pulse 2026 AI Survey if accessible and methodology is adequate.

Extraction note: `research/2026-evidence/legal-ai.md`.

Decision: replace `AI Tool Adoption (BigLaw)` with `Legal Organization GenAI Adoption`, refresh solo/small adoption with Clio 2026, use Thomson Reuters 2026 for broad legal workflow context, and keep ABA/TAR as the direct document-review workflow anchor. Do not score Thomson's raw 74% document-review use-case share without a denominator adjustment.

### hire

Implemented state: current v2 scoring fixes the old indicator-name mismatch and uses talent-acquisition adoption, screening use-case adoption, broad AI across hiring processes, and AI assessment platform reach.

Locked approach:

- Use `Orgs Using AI in Talent Acquisition` as the canonical org-level adoption indicator.
- Use one org-level adoption metric, one application/screening metric, and one platform-reach metric.
- Add 2026 HR/recruiting reports where they measure screening, scheduling, or candidate communication automation.

Candidate sources:

- SHRM 2026 State of AI in HR.
- ICIMS/Aptitude 2026 AI adoption in talent acquisition report.
- SHRM 2025 Talent Trends for resume-screening specifics if no 2026 equivalent is available.
- Vendor platform disclosures such as HireVue only for reach, with lower confidence.

Extraction note: `research/2026-evidence/hire.md`.

Decision: high-priority cleanup domain. The implemented v2 contract uses `Orgs Using AI in Talent Acquisition`, replaces `AI-Screened Applications` with `AI Screening Use Case Adoption`, adds `Broad AI Across Hiring Processes`, and keeps `AI Assessment Platform Reach` low-confidence until refreshed multi-vendor volume is available. The strongest 2026 workflow signal is ICIMS/Aptitude screening use-case adoption at 58%, while SHRM provides a conservative broad-market guardrail at 27% recruiting-practice AI use.

### education

Implemented state: current v2 scoring uses student AI schoolwork use, AI-graded assessments, teacher AI work use, and a low-weight student AI-written-output signal.

Locked approach:

- Replace `Students Using AI Tutors` with the broader `Students Using AI for Schoolwork` indicator.
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

Decision: rename the student indicator to `Students Using AI for Schoolwork`. Use Pew for K-12 schoolwork depth, HAI for broad cross-level context, OECD for teacher workflow, and keep grading plus AI-written-output signals separate. `Student Papers 80%+ AI-Written` is scored as a low-weight output-share signal, not as assessment delegation.

## Future Source Lock Gates

For future score changes, complete these gates before updating `seed/seed.json`:

1. One row per candidate source in `research/2026-source-ledger.csv`.
2. For every scored indicator, record `source_url`, `published_at`, `measurement_period`, `evidence_grade`, `confidence`, and `included_in_score`.
3. For each displayed-but-unscored indicator, explicitly choose `score`, `display_only`, or `drop`.
4. Recheck indicator-name contracts against `internal/collect/score.go`.
5. Align methodology copy with code thresholds or change the thresholds deliberately.
6. Decide whether `DataFreshness` should be derived from source ledger dates rather than set in the run snapshot.
7. Recompute the prior public baseline from current locked sources before adding new values, to separate methodology drift from data drift.

## Maintenance Worklist

1. Promote the source ledger from research CSV into structured seed or export metadata.
2. Add `included_in_score`, `source_url`, `measurement_period`, evidence grade, and confidence metadata to the public data export.
3. Add run-level source snapshots so future updates can show which sources changed between public analysis runs.
4. Recheck low-confidence proxy indicators before the next score update, especially AI assessment platform reach, production support-agent deployment, and AI-assisted diagnosis rate.
5. Decide whether `DataFreshness` should be computed from source-ledger dates once source metadata is first-class.
