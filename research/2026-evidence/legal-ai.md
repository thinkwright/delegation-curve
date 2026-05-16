# Legal AI 2026 Evidence Extraction

Status: implemented for the 2026 Q2 source refresh.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `legal-ai` score: 46.3.

Configured scoring formula:

- `Legal Organization GenAI Adoption`: 40% weight.
- `Solo and Small Firms Using AI for Legal Work`: 30% weight.
- `AI-Assisted Document Review`: 30% weight.

Current observations:

- `Legal Organization GenAI Adoption`: 40.0%, freshness 2026.
- `Solo and Small Firms Using AI for Legal Work`: 73.0%, freshness 2026.
- `AI-Assisted Document Review`: 28.0%, freshness 2024.

The current score calculates as:

```text
0.40 * 40.0 + 0.30 * 73.0 + 0.30 * 28.0 = 46.3
```

The implemented 2026 Q2 contract uses broad legal-organization GenAI adoption, solo/small legal-work adoption, and a direct document-review workflow anchor. Adoption still does not equal autonomous legal decisioning, so legal benchmark and agentic claims remain context.

## Extracted Candidate Sources

### Thomson Reuters 2026 AI in Professional Services Report

Source: https://www.thomsonreuters.com/content/dam/ewp-m/documents/thomsonreuters/en/pdf/reports/2026-ai-in-professional-services-report.pdf
Published: 2026.
Measurement period: 2026 survey release.
Evidence grade: B for broad professional/legal adoption and workflow context.
Confidence: medium.

Relevant values:

- Survey drew on more than 1,500 professionals across more than two dozen countries.
- 40% of respondents said their organizations use GenAI, up from 22% in the prior year.
- Only 19% said their organizations had no plans to adopt GenAI.
- More than 50% of professionals use publicly available GenAI tools for work.
- Professional-grade and industry-specific GenAI tools are moving toward majority use.
- Among current GenAI users, more than 80% use the tools at least weekly.
- 87% of professionals believe GenAI will be central to their workflow within five years.
- 15% of organizations use agentic AI and another 53% are planning or considering it.
- Top legal GenAI use cases among users include legal research at 80%, document review at 74%, document summarization at 73%, brief or memo drafting at 59%, correspondence drafting at 55%, and contract drafting at 49%.
- Only 18% of respondents said their organizations collect ROI metrics for AI.
- Roughly two-thirds of corporate respondents believe outside firms should use AI, but fewer than 20% mandate it.

Recommendation: use as a strong 2026 legal/professional adoption and workflow context source. Do not score the 74% legal document-review use case raw because it is a use-case share among GenAI users, not a whole-market document-review delegation rate.

### Clio Legal Trends for Solo and Small Law Firms 2026

Source: https://www.clio.com/about/press/2026-solo-small-firm-report/
Published: 2026-05-04.
Measurement period: 2026 report release.
Evidence grade: B for solo/small-firm adoption.
Confidence: medium-high.

Relevant values:

- 71% of solo practitioners use AI to complete legal work.
- 75% of small firms use AI to complete legal work.
- Fewer than 33% of solo and small firms have increased revenues with AI.
- Nearly 60% of enterprise firms have increased revenues with AI.
- 86% of solo firms and 78% of small firms have made no pricing changes in response to AI.
- Larger firms are significantly more likely to use specialized legal AI tools for document drafting, e-discovery, analytics, and contract review.
- The full report is based on more than 1,700 respondents and aggregated anonymized data from tens of thousands of legal professionals.

Recommendation: this is the strongest 2026 replacement for the stale `AI Tool Adoption (Solo/Small)` value. Treat it as legal-work AI adoption, not as output share or autonomous legal decisioning.

### Clio Legal Trends for Mid-Sized Law Firms 2026

Source: https://www.clio.com/about/press/ai-is-reshaping-how-mid-sized-law-firms-scale-clio-reports/
Published: 2026-03-09.
Measurement period: 2026 report release.
Evidence grade: B for mid-sized firm adoption and operational impact.
Confidence: medium.

Relevant values:

- 86% of mid-sized firms report using AI.
- 60% report having formal AI-use policies.
- 65% say AI enables them to take on higher volumes of work.
- 58% of legal professionals in mid-sized firms say AI has enabled them to take on more complex work.
- 57% report improved work-life balance and 50% report lower stress.
- 30% report difficulty incorporating new technology into existing workflows.
- Methodology is a survey of more than 1,000 U.S. legal professionals from both an independent market panel and current customers.

Recommendation: useful for replacing the BigLaw/small split with a firm-size adoption curve if methodology v2 broadens beyond the current BigLaw and solo/small pair. It is not exactly BigLaw.

### ABA 2024 Litigation and TAR TechReport

Source: https://www.americanbar.org/groups/law_practice/resources/tech-report/2024/2024-litigation-and-tar-techreport/
Published: 2025-04-30.
Measurement period: 2024 Legal Technology Survey.
Evidence grade: A for litigation/TAR workflow adoption.
Confidence: medium.

Relevant values:

- 50% of respondents reported having litigation support software available at their firms.
- 31% personally used litigation support software.
- 40% of litigation support software users named Relativity among top products.
- Useful litigation support features included document review at 36%, full-text search at 38%, bates stamping at 30%, redaction at 25%, and OCR at 24%.
- For ESI review and processing, keyword search remained dominant at 85%.
- Natural language search was 65%, concept searching 37%, AI-assisted search 28%, and predictive coding 22%.
- 76% cited unfamiliarity as the primary reason for not using predictive coding.

Recommendation: keep as the direct workflow anchor for `AI-Assisted Document Review`. It is older than the 2026 adoption sources, but it has a cleaner denominator for a delegated legal workflow.

### ABA 2024 Artificial Intelligence TechReport

Source: https://www.americanbar.org/groups/law_practice/resources/tech-report/2024/2024-artificial-intelligence-techreport/
Published: 2025-04-25.
Measurement period: October-December 2024 survey.
Evidence grade: A for broad lawyer AI adoption.
Confidence: medium.

Relevant values:

- For online legal research respondents, 30.2% said their offices were using AI-based technology tools.
- Firms with 500 or more lawyers reported 47.8% usage.
- Firms with 10-49 lawyers reported 29.5% usage.
- Firms with 2-9 lawyers reported 24.1% usage.
- Solo practitioners reported 17.7% usage.
- Leading AI-based research tools already adopted or seriously considered were ChatGPT at 52.1%, Thomson Reuters CoCounsel at 26.0%, and Lexis+ AI at 24.3%.
- The leading perceived benefit was saving time or increasing efficiency at 54.4%.
- Concerns included accuracy at 74.7%, reliability at 56.3%, and data privacy/security at 47.2%.

Recommendation: keep as a 2024/2025 baseline and governance/caution source. It is no longer fresh enough to anchor 2026 scoring by itself.

## Implemented Legal AI Source Lock

Current v2 scoring inputs and retained candidates:

| Indicator | Role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| Legal Organization GenAI Adoption | score input | B | medium | Thomson 40% organization GenAI adoption gives broader professional/legal context |
| Solo and Small Firms Using AI for Legal Work | score input | B | medium-high | Clio 71% solo and 75% small is a major refresh from stale 18.5% |
| AI-Assisted Document Review | score input | A/B | medium | ABA 28% AI-assisted search and 22% predictive coding remain cleaner than raw Thomson use-case shares |
| Legal GenAI Workflow Use | context | B | medium | Thomson top use cases show research 80% and document review 74% among legal GenAI users |

Implementation decision:

- Replace BigLaw-specific adoption with `Legal Organization GenAI Adoption` from Thomson Reuters.
- Refresh solo/small adoption from Clio, but label it adoption rather than delegation.
- Do not replace `AI-Assisted Document Review` with Thomson's 74% raw use-case share; use it as context unless an explicit blended proxy is approved.
- Keep legal benchmark performance and agentic expectations out of the score unless observed deployed workflow data is available.
