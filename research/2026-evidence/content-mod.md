# Content Moderation 2026 Evidence Extraction

Status: implemented for the 2026 Q2 source refresh.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `content-mod` score: 96.1.

Configured scoring formula:

- `Meta Automated Detection`: 37.5% weight.
- `YouTube Automated Flagging`: 31.25% weight.
- `TikTok Automated Enforcement`: 31.25% weight.

Current observations:

- `Meta Automated Detection`: 95.2%, freshness Q3 2025.
- `YouTube Automated Flagging`: 99.5%, freshness Jul-Sep 2025.
- `TikTok Automated Enforcement`: 93.8%, freshness H2 2025.

The current score calculates as:

```text
0.375 * 95.2 + 0.3125 * 99.5 + 0.3125 * 93.8 = 96.1
```

The implemented 2026 Q2 contract keeps each platform indicator aligned to its source semantics. X and cross-platform DSA data are retained as context until a stable comparable 2025/2026 extraction is locked.

## Extracted Candidate Sources

### Meta Community Standards Enforcement and Enforcement Policy Change

Sources:

- https://transparency.meta.com/reports/community-standards-enforcement/
- https://about.fb.com/news/2025/01/meta-more-speech-fewer-mistakes/amp/

Published or updated: 2025-2026.
Measurement period: current report UI plus Q1 2025 policy-change context.
Evidence grade: A for Meta CSER values; B for methodology-change context.
Confidence: medium until the latest CSER table is extracted reproducibly.

Relevant values and constraints:

- Meta's Community Standards Enforcement Report remains the direct source for proactive-rate values.
- The public report page is dynamic and should be extracted with a reproducible browser or data-source workflow, not copied manually from the UI.
- Meta changed enforcement policy in 2025 by focusing proactive enforcement on illegal and high-severity violations.
- Meta reported roughly a 50% reduction in U.S. enforcement mistakes from Q4 2024 to Q1 2025.
- Meta said it had previously used automated systems to scan for all policy violations and would tune systems to require higher confidence before taking down content.
- Meta said future reporting would include enforcement-mistake metrics.

Recommendation: keep Meta as a scored source, but treat 2025 as a methodology break. Refreshing the value is less important than recording whether the proactive-rate definition is still comparable after the policy change.

### YouTube Community Guidelines Enforcement

Sources:

- https://transparencyreport.google.com/youtube-policy/removals
- https://support.google.com/transparencyreport/answer/9198203?hl=en-GB
- https://support.google.com/transparencyreport/answer/9209072?hl=en

Published or updated: December 2025.
Measurement period: latest visible data includes July-September 2025.
Evidence grade: A.
Confidence: high for automated flagging; medium for current score compatibility.

Relevant values:

- YouTube's visible-change log reports automated flagging at 99.5% in the recent data table.
- Beginning with the July-September 2025 reporting period, YouTube began providing flagging data for suicide, self-harm, and eating-disorder content.
- Beginning with the April-June 2025 period, YouTube changed parts of channel-termination classification and counting.
- YouTube says automated systems scan content when users attempt to publish video or post comments and can prevent re-uploads of known violative content.
- YouTube says some high-confidence cases can be automatically detected and removed while other automated flags go to trained reviewers.

Recommendation: use automated flagging as a direct detection source, but rename `Google Automated Removal` if the scored value is really detection/flagging rather than final removal.

### TikTok DSA Transparency Report H2 2025 and Q4 2025 CGER

Sources:

- https://newsroom.tiktok.com/digital-services-act-our-sixth-transparency-report-on-content-moderation-in-europe?lang=en-150
- https://www.tiktok.com/transparency/en-us/reports/

Published: 2026-04-29 for DSA H2 2025.
Measurement period: July-December 2025 for DSA; Q4 2025 for global CGER.
Evidence grade: A.
Confidence: high for DSA value; medium for global score compatibility.

Relevant values:

- TikTok's sixth DSA report covers H2 2025.
- TikTok reported 178 million monthly active EU recipients.
- TikTok removed around 112 million pieces of content under DSA-scoped terms and policies during H2 2025.
- Automated systems actioned 93.8% of all violating content without human review.
- 97.6% of automated enforcement decisions were confirmed as correct.
- Public reporting around the Q4 2025 global Community Guidelines Enforcement Report indicates 175.3 million videos removed, 152.6 million detected and taken down using automated detection technologies, and a 99.1% proactive removal rate. This should be verified from the official CGER data before scoring.

Recommendation: TikTok is refreshable, but choose between EU DSA automated enforcement and global CGER proactive removal. The two are not the same denominator.

### X Global Transparency and DSA Reports

Sources:

- https://transparency.x.com/content/dam/transparency-twitter/2025/x-global-transparency-report_h2_2024.pdf
- https://transparency.x.com/dsa-transparency-report-2025-april.html
- https://transparency.x.com/dsa-transparency-report-2025-october.html

Published: 2025 and 2026 pages located.
Measurement period: H2 2024 global report plus 2025 DSA reports.
Evidence grade: A for platform reports.
Confidence: low-medium for current score compatibility.

Relevant values:

- The latest official global report located in search is H2 2024, not a global H1/H2 2025 report.
- The H2 2024 global PDF exposes automated versus human counts by enforcement category.
- Automation differs sharply by category: platform manipulation and spam actions appear heavily automated, while some content-removal categories report no automated removals.
- X also has DSA transparency pages for 2025, but EU DSA reports are not a direct substitute for the global report.
- The current seed value of 88.6% with freshness H1 2025 needs source provenance before it is refreshed or reused.

Recommendation: keep X in the family only if a stable comparable global report can be found. Otherwise demote X to context or replace it with DSA-derived cross-platform data in methodology v2.

### EU DSA Transparency Database Research API

Source: https://transparency.dsa.ec.europa.eu/page/research-api
Published or accessed: 2026-05-16.
Measurement period: rolling last six months.
Evidence grade: A if snapshotted.
Confidence: medium until access and extraction are implemented.

Relevant values and constraints:

- The Research API supports search, SQL-like queries, counts, query, aggregates, labels, and platforms endpoints.
- All endpoints require Bearer-token authentication.
- The `statement_index` contains only statements of reasons from the last six months; older statements are not available through the Research API.
- Relevant fields include `automated_decision` and `automated_detection`.
- The documentation includes an example for analysis of automated means in content moderation using `automated_detection` and grouping by platform/content type.

Recommendation: DSA is valuable for methodology v2 only if we request access, implement a reproducible extraction, and snapshot the dataset for each analysis run. Otherwise it should remain display/context.

## Implemented Content Moderation Source Lock

Current v2 scoring inputs and retained candidates:

| Indicator | Role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| Meta Automated Detection | score input | A | medium | Keep but mark 2025 methodology break after policy shift |
| YouTube Automated Flagging | score input | A | high | Rename from Google automated removal if using flagging |
| TikTok Automated Enforcement or Proactive Removal | score input | A | medium-high | Choose EU DSA 93.8% or global CGER proactive rate but not both without reconciliation |
| X Automated Enforcement | hold or context | A | low-medium | Direct 2025 global report not locked |
| DSA Cross-Platform Automated Rate | future score input only after extractor | A | medium | Requires token and run-level snapshots because retention is rolling |

Implementation decision:

- Keep the Meta, YouTube, and TikTok source family, with indicator names normalized to what each source actually measures.
- Treat Meta 2025 as a methodology break and verify CSER definitions before comparing to prior proactive rates.
- Use YouTube automated flagging directly instead of the old `Google Automated Removal` label.
- Use TikTok DSA H2 2025 automated enforcement as the current score input unless global CGER data can be extracted directly.
- Keep X and cross-platform DSA data as context until a comparable current global report or tokened DSA extractor and snapshot workflow exists.
