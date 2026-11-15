# Customer Support: 2026 Q4 Evidence Notes

Run `2026-q4`, method `delegation-curve-v2`. Prior notes: [2026 Q2](../2026-q2/support.md).

Domain score unchanged at **40.5**; all four inputs are held.

## Lock Decisions

| Indicator | Action | Value | Source | Window | Grade / confidence |
| --- | --- | ---: | --- | --- | --- |
| Cases Handled by AI | hold | 30% | Salesforce State of Service, 7th edition | Apr 25-Jun 6, 2025 | B / medium-high |
| Bot Deflection Rate | hold | 52.3% | Prior locked Intercom value | not recorded | A / low |
| Production AI Customer Communications Agents | hold | 62% | Sinch AI Production Paradox | Jan-Feb 2026 | B / medium |
| Mature AI Support Deployment | hold | 10% | Intercom 2026 Customer Service Transformation Report | Q4 2025 | B / medium |

## Notes

- **Salesforce.** No 8th edition of State of Service has been published. The May 2026 AI Agents edition (fieldwork March-April 2026) reports adoption of AI agents (66%), not a share of cases.
- **Bot Deflection Rate.** The 52.3% could not be traced to a published source or definition. It is held at low confidence and flagged for replacement or retirement in v3.
- **Sinch.** The July 2026 release reuses the same survey; a September release surveys consumers, a different construct.
- **Intercom** now operates as Fin, which Salesforce agreed to acquire in June 2026. The next transformation report is expected around January 2027.

## Candidates For Methodology v3

- **Fin automation rate** (share of all conversations resolved by AI), published for leading customers by industry. Fin's headline resolution rate (76%) is single-vendor and [changed definition in July 2026](https://www.intercom.com/help/en/articles/15599377-update-to-fin-performance-metrics) to exclude conversations Fin never answered, so it is not a stable series: <https://fin.ai/benchmarks>.
- **Gorgias:** median 45% of AI-handled tickets closed with no human involvement (Oct 2025-Apr 2026): <https://www.gorgias.com/research/the-cheapest-ticket-is-the-one-a-human-never-touches>.
- **Ada:** 72% of conversations contained without a human, 52% graded as resolved: <https://www.ada.cx/resources/guide/what-is-automated-resolution-in-ai-customer-service/>.
- **Gartner (customer survey, Feb-Mar 2026):** 7% of customers used a chatbot in their most recent service interaction, a demand-side guardrail.

## Before Publication

Zendesk CX Trends 2027 and any new Salesforce State of Service edition usually appear in November.
