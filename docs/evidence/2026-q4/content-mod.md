# Content Moderation: 2026 Q4 Evidence Notes

Run `2026-q4`, method `delegation-curve-v2`. Prior notes: [2026 Q2](../2026-q2/content-mod.md).

Domain score 96.1 → **95.5**. On consistent methods the three platforms are flat to rising; the decline is a correction of two 2026 Q2 inputs (the corrected Q2 equivalent is 95.3).

## Lock Decisions

| Indicator | Action | Value | Source | Window | Grade / confidence |
| --- | --- | ---: | --- | --- | --- |
| Meta Automated Detection | refresh, corrects Q2 | 94.2% | [Meta Community Standards Enforcement Report](https://transparency.meta.com/reports/community-standards-enforcement/), H1 2026, official CSV `CSER-2026_Q2.csv` | Apr-Jun 2026 | A / medium |
| YouTube Automated Flagging | refresh, corrects Q2 | 98.6% | [YouTube Community Guidelines report, removals](https://transparencyreport.google.com/youtube-policy/removals?hl=en) | Apr-Jun 2026 | A / high |
| TikTok Automated Enforcement | refresh | 94.1% | [TikTok seventh DSA transparency report](https://newsroom.tiktok.com/digital-services-act-our-seventh-transparency-report-on-content-moderation-in-europe?lang=en-150), published 2026-08-31 | Jan-Jun 2026 | A / high |

## Method And Corrections

- **Meta.** Meta defines the proactive rate as content actioned that Meta found before users reported it, divided by all content actioned. The scored value weights each policy area's proactive rate by content actioned, across Facebook and Instagram (13 policy areas each), excluding Fake Accounts (no proactive rate is published from 2026) and Spam (whose volume would dominate). On this method the series is 94.5 (Q3 2025), 94.9, 95.0, 94.2 (Q2 2026). The 2026 Q2 input (95.2) could not be reproduced by any natural aggregation of the official CSV.
- **YouTube.** The scored value is the share of removed videos whose first detection was automated flagging: 11,005,420 of 11,159,676 in Apr-Jun 2026. The 2026 Q2 input (99.5%) came from a change-log entry giving the automated share of comment removals for Jul-Sep 2023; the video share for Jul-Sep 2025 was 97.9%.
- **TikTok.** The published headline ("automated systems actioned 94.1% of violating content without human review"); the report's own data template gives 94.0% as a cross-check.

## Comparability

- Meta's January 2025 enforcement shift toward illegal and high-severity violations continues: proactive rates in hateful conduct, bullying, and violence and incitement are 20-30 points below late 2024, while high-severity areas remain at 97-99%. Meta now reports semiannually; the next report (H2 2026) is expected around March 2027. A new policy area (criminal organizations) was added in 2026 and moves the aggregate by under 0.1 point.
- TikTok's methodology break came in H2 2025 (EU reporting template, comments added, de-duplicated counting); H1 2026 is comparable with H2 2025.
- YouTube redesigned its report in September 2026; the change affects channel categories, not video detection sources.

## Candidates For Methodology v3

- **EU DSA Transparency Database.** Public, checksummed monthly aggregates record whether each moderation decision was fully, partially, or not automated across all platforms. In H1 2026, 43.6% of 3.35 billion statements of reasons were fully automated (63.4% fully or partially). It is a decision-level measure, but volume is dominated by marketplaces (Google Shopping alone files about half), and platforms label automation differently (Meta marks almost all decisions "partially" automated), so it is not usable as an all-platform share without a restricted definition.
- **Meta enforcement precision** (about 91-92% at the end of H1 2026) as a guardrail on automated detection.
