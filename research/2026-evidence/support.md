# Customer Support 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `support` score: 45.9.

Configured indicators:

- `AI Resolution Rate`: 30% weight.
- `Bot Deflection Rate`: 25% weight.
- `Orgs Using AI Support`: 25% weight.
- `AI Copilot Adoption (Agents)`: 20% weight.

Current observations:

- `AI Resolution Rate`: 41%, freshness 2025.
- `Bot Deflection Rate`: 52.3%, freshness 2025.
- `Orgs Using AI Support`: 63%, freshness 2025.
- `AI Copilot Adoption (Agents)`: 24%, freshness 2025.

The current contract is directionally right: customer support is one of the clearest real-world delegation domains. The main risk is metric mismatch. 2026 sources often report adoption, investment, maturity, or expected resolution rather than measured resolved-case or true deflection rates.

## Extracted Candidate Sources

### Intercom 2026 Customer Service Transformation Report

Source: https://www.intercom.com/customer-transformation-report
Published: 2026.
Measurement period: Q4 2025.
Evidence grade: B for support AI adoption and maturity.
Confidence: medium.

Relevant values:

- Sample size: 2,470 support professionals across NAMER, EMEA, LATAM, and APAC.
- 82% of senior leaders say their teams invested in AI for customer service over the last 12 months.
- 87% of senior leaders plan to invest in AI for customer service in 2026.
- 10% of respondents say they have reached mature deployment, defined by Intercom as AI fully integrated into support operations and working at scale.
- 87% of mature-deployment teams report improved metrics since implementing AI, compared with 62% overall.
- 58% of teams cite improving customer experience as the top 2026 priority.
- 40% of teams report agents spending more time training and optimizing AI systems.
- 52% of organizations plan to scale AI beyond support in 2026, and nearly one-third say customer service teams are leading that effort.

Recommendation: use Intercom as an operational maturity source and as context for `Orgs Using AI Support`. Do not use it directly for `AI Resolution Rate` or `Bot Deflection Rate` unless the full report exposes directly comparable resolution/deflection values.

### Salesforce State of Service 7th Edition

Source: https://www.salesforce.com/news/stories/state-of-service-report-announcement-2025/
Published: 2025-11-13.
Measurement period: survey conducted 2025-04-25 to 2025-06-06.
Evidence grade: B for case-handled-by-AI and service workflow adoption.
Confidence: medium-high.

Relevant values:

- Global double-anonymous survey of 6,500 service professionals and decision makers.
- Service teams estimate 30% of cases are currently handled by AI.
- Service teams project AI will handle 50% of cases by 2027.
- 79% of service leaders believe investing in AI agents is essential to meet current business demands.
- Companies expect AI agents to reduce service costs and case resolution times by 20% on average.
- Service representatives using AI report spending 20% less time on routine cases.
- 51% of service leaders say security concerns have delayed or limited AI initiatives.

Recommendation: treat the 30% currently handled by AI value as the strongest candidate for a refreshed `AI Resolution Rate` or a renamed `Cases Handled by AI` indicator. It is closer to delegation than broad org adoption, but still survey-estimated and not necessarily true end-to-end resolution.

### Zendesk 2026 CX Trends

Source: https://cxtrends.zendesk.com/
Published: 2026.
Measurement period: 2026 report.
Evidence grade: B for customer expectations and CX leader sentiment.
Confidence: medium.

Relevant values:

- 83% of CX leaders say memory-rich AI agents are key to personalized customer journeys.
- 74% of consumers say AI has made them expect customer service to be available 24/7.
- 88% of customers expect faster response times than one year earlier.
- 95% of consumers expect explanations for AI-made decisions.
- 37% of CX leaders currently offer reasoning behind AI decisions.

Recommendation: use Zendesk as context for customer-facing AI expectations, transparency, and design pressure. Do not use the public CX Trends page for the current `AI Resolution Rate` until a direct and cited automated-resolution metric is available.

### Gartner 2026 Customer Service and Support Survey

Source: https://www.gartner.com/en/newsroom/press-releases/2026-02-18-gartner-survey-finds-ninety-one-percent-of-customer-service-leaders-under-pressure-to-implement-ai-in-2026
Published: 2026-02-18.
Measurement period: survey conducted 2025-10.
Evidence grade: B for executive pressure and service-model redesign.
Confidence: medium.

Relevant values:

- Survey of 321 customer service and support leaders.
- 91% report pressure from executive leadership to implement AI.
- Leaders identify customer satisfaction, operational efficiency, and self-service success as top 2026 priorities.
- Nearly 80% of organizations plan to transition at least some agents into new roles.
- 84% plan to add new skills to the agent role and adjust hiring profiles.
- 58% aim to upskill agents into knowledge-management specialist roles.

Recommendation: use Gartner for governance and labor-redesign context. It should not directly score support delegation because it measures pressure and planned operating-model change rather than resolved cases.

### Hiver State of AI Customer Support in 2026

Source: https://hiverhq.com/reports/state-of-ai-customer-support-2026
Published: 2026.
Measurement period: 2026 report.
Evidence grade: C for outcome sentiment.
Confidence: medium-low.

Relevant values:

- Survey of 700+ support leaders across the United States.
- 14% of support leaders say AI has significantly improved resolution times.
- 48% say their teams are moderately confident using AI.
- 50% say AI has not significantly lowered cost per ticket.
- 61% remain cautious about AI representing their brand.
- 25% say AI has clearly reshaped their team structure.

Recommendation: use Hiver as a cautionary counterweight to vendor-success narratives. It is useful for measuring realized outcome skepticism, but it is not a direct resolution-rate or deflection-rate source.

### Sinch AI Production Paradox

Source: https://sinch.com/news/sinch-releases-ai-production-paradox/
Published: 2026-05-13.
Measurement period: survey conducted 2026-01 to 2026-02.
Evidence grade: B for production deployment and rollback risk.
Confidence: medium.

Relevant values:

- Independent survey of 2,527 senior decision makers across 10 countries and six industries.
- 62% of enterprises already have AI agents live in production across customer communications.
- 88% expect to have AI agents in production by the end of 2026.
- 74% have rolled back or shut down a deployed AI agent after a governance failure.
- 81% rollback rate among organizations with fully mature guardrails.
- 98% are increasing AI communications investment in 2026.
- 84% of AI communications engineering teams spend at least half their time on safety infrastructure.
- 55% have to build custom infrastructure for cross-channel context.

Recommendation: add Sinch as a new production-readiness and reliability source. It is broader than customer support, but customer communications is adjacent and highly relevant. Keep it separate from success scoring unless the next methodology adds a reliability or rollback adjustment.

### Stanford HAI 2026 AI Index Economy Chapter

Source: https://hai.stanford.edu/assets/files/ai_index_report_2026_chapter_4_economy.pdf
Published: 2026.
Measurement period: synthesis of productivity studies through 2025.
Evidence grade: B for productivity context.
Confidence: medium-high.

Relevant values:

- Customer support agents using a conversational AI assistant resolved 14%-15% more issues per hour.
- The AI Index frames support work as one of the clearest task-level productivity-gain areas, alongside software development and marketing.
- The support productivity study is not a 2026 deployment prevalence measure; it is evidence that agent assist can improve issue throughput in a real support setting.

Recommendation: use AI Index only as productivity context. It supports the importance of the domain, but it should not replace adoption or resolution share metrics.

## Proposed Support Source Lock

Proposed v2 scoring candidates:

| Indicator | Suggested role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| Cases Handled by AI | score input or rename of `AI Resolution Rate` | B | medium-high | Salesforce 30% current and 50% projected by 2027; closer to delegation than org adoption |
| Bot Deflection Rate | score input if direct source found | A/B | low | Still needs a refreshed operational deflection source with clear definition |
| Orgs Using AI Support | score input or context | B | medium | Intercom 82% invested and 10% mature deployment; Salesforce 79% investment-essential |
| AI Copilot Adoption (Agents) | score input if agent-level source found | B | low | Current sources speak to reps using AI or agents spending less time on routine work, not clean agent-level adoption |
| Production AI Customer Communications Agents | display or reliability modifier | B | medium | Sinch 62% live in production and 74% rollback after governance failure |
| AI Support Productivity Lift | context only | B | medium-high | AI Index 14%-15% more issues resolved per hour |

Near-term decision:

- Rename `AI Resolution Rate` to `Cases Handled by AI` unless a true autonomous-resolution source is found.
- Keep `Bot Deflection Rate` locked only after the source defines whether deflection means no escalation, customer-confirmed resolution, or completed backend workflow.
- Add Sinch as a new production-reliability source, but keep it out of positive scoring unless a reliability penalty or maturity modifier is added.
- Use Hiver and Gartner as caution/context, not score inputs.
- Avoid mixing projected 2027 values into a 2026 score; use Salesforce 30% current as the candidate score value if we proceed before stronger telemetry is found.
