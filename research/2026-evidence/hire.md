# Hire 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `hire` score: 36.8.

Configured scoring formula:

- `Orgs Using AI Screening`: 40% weight.
- `AI-Screened Applications`: 35% weight.
- `AI Assessment Platform Reach`: 25% weight, normalized as `value / 300M`.

Current observations:

- `Orgs Using AI in Recruiting`: 51%, freshness Feb 2025, display-only because the configured name is `Orgs Using AI Screening`.
- `AI-Screened Applications`: 44%, freshness Q1 2025.
- `AI Assessment Platform Reach`: 80M/yr, freshness 2025, normalized to 26.7%.
- `HR Pros Using AI for Recruiting`: 69%, freshness Feb 2025, display-only.
- `AI-Using Recruiters Screening Resumes`: 44%, freshness Feb 2025, display-only.
- `Deep TA Stack AI Integration`: 14%, freshness Feb 2024, display-only.
- `AI Assessment Volume`: 80M/yr, freshness Q1 2024, display-only.

The active score currently reweights around the two configured indicators that match observations:

```text
(0.35 * 44 + 0.25 * 26.7) / (0.35 + 0.25) = 36.8
```

If the existing 51% `Orgs Using AI in Recruiting` observation were wired into the current score retroactively, the score would become:

```text
0.40 * 51 + 0.35 * 44 + 0.25 * 26.7 = 42.5
```

That should not be applied as a silent historical correction. The name/contract fix should happen in the next methodology version or be explicitly re-snapshotted as a corrected baseline.

## Extracted Candidate Sources

### SHRM State of AI in HR 2026

Source: https://www.shrm.org/topics-tools/research/state-of-ai-hr-2026/full-report
Published: 2026.
Measurement period: December 2025 survey.
Evidence grade: B for broad HR and recruiting adoption.
Confidence: high for broad HR adoption; medium for hiring delegation.

Relevant values:

- SHRM surveyed 1,722 employed HR professionals between 2025-12-05 and 2025-12-23.
- 39% of organizations currently had AI adopted in their HR functions.
- 7% intended to launch AI in HR functions during 2026.
- 62% were currently using AI somewhere in their organizations.
- Recruiting was the most common HR practice area using AI at 27%.
- HR technology followed at 21%, learning and development at 17%, and employee experience at 14%.
- SHRM assessed 138 use cases across 16 HR practice areas and found real-world AI applications concentrated in process-driven tasks, especially recruiting.
- Common examples included resume parsing, interview scheduling, job ad programming, candidate-job matching, and personalized recommendations.
- 49% of organizations currently using or about to pilot AI had policies to regulate AI use.
- In states with workplace AI regulations, 57% of HR professionals were not aware of the relevant rules.
- 72% of HR professionals said nontechnical barriers would still prevent complete HR automation if technical barriers were removed.

Recommendation: use SHRM as the broad organization-level guardrail. It supports a conservative `Orgs Using AI in Recruiting` or `Orgs Using AI in HR Recruiting` indicator, but the 27% value is not directly comparable to the older 51% SHRM Talent Trends recruiting value without checking denominator definitions.

### ICIMS / Aptitude Definitive Guide to AI Adoption in Talent Acquisition

Source: https://www.icims.com/company/newsroom/aiadoptionreport2026/
Published: 2026-04-30.
Measurement period: 2026 survey release.
Evidence grade: B for talent-acquisition workflow adoption.
Confidence: medium-high.

Relevant values:

- Survey covered more than 400 U.S. talent acquisition leaders and practitioners.
- 69% of companies reported using AI in some capacity in talent acquisition.
- 18% reported using AI broadly across hiring processes.
- 58% of talent acquisition leaders were unclear about the difference between AI and automation.
- Screening was the most widely adopted AI use case at 58%.
- Candidate communication followed at 54%, assessments at 50%, and sourcing at 46%.
- Recruiters were the most frequent users of AI tools at 46%, followed by hiring managers at 43%.
- When conflicts arise, recruiter judgment overrides AI recommendations in 58% of organizations.
- 82% of companies said transparency and explainability were important.
- 45% did not yet have a formal AI governance framework.
- 46% of companies were using or planning to use agentic AI for talent acquisition.

Recommendation: use ICIMS/Aptitude as the strongest 2026 talent-acquisition workflow source. Prefer `AI Screening Use Case Adoption` over `AI-Screened Applications` unless a direct application-level denominator is found.

### HireVue 2026 Global AI in Hiring Report

Source: https://www.hirevue.com/resources/report/2026-global-ai-in-hiring-report
Published: 2026.
Measurement period: 2026 survey release.
Evidence grade: C for scoring; B for context.
Confidence: medium.

Relevant values:

- HireVue surveyed more than 3,100 global hiring managers.
- 77% of HR teams use AI regularly.
- 71% of candidates use AI for resumes.
- Only 41% of hiring teams fully trust AI.

Recommendation: keep as context because it is a vendor survey and does not directly refresh the platform-reach denominator. It is useful for the candidate-side AI pressure story.

### HireVue Q1 2024 Platform Volume Update

Source: https://www.hirevue.com/blog/hiring/quarterly-updates-from-hirevue-ceo
Published: 2024-04-22.
Measurement period: Q1 2024.
Evidence grade: B for platform volume.
Confidence: low-medium for market-level scoring.

Relevant values:

- Candidates completed nearly 20 million assessments on HireVue in Q1 2024.
- The assessment count included Virtual Job Tryouts, game-based assessments, coding challenges, and video interviews.
- More than 1,100 customers conducted millions of interviews and sent tens of millions of chat and text messages in Q1 2024.
- HireVue reported volume increases in retail, hospitality, recreation and leisure, government, and communications.

Recommendation: keep the current 80M/yr platform-reach approximation only as a low-confidence reach proxy until a refreshed multi-vendor denominator is found. Do not treat vendor assessment volume as application share.

## Proposed Hire Source Lock

Proposed v2 scoring candidates:

| Indicator | Suggested role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| Orgs Using AI in Talent Acquisition | score input | B | medium | ICIMS 69% in TA sample or SHRM 27% recruiting practice area; denominator choice materially changes score |
| AI Screening Use Case Adoption | score input | B | medium-high | ICIMS 58% screening use case is closer to delegated screening than broad org adoption |
| AI Assessment Platform Reach | low-weight score input or context | B/C | low-medium | Current 80M/yr proxy still needs a refreshed multi-vendor denominator |
| Broad AI Across Hiring Processes | display/context | B | medium | ICIMS 18% broad use across hiring processes captures depth better than headline adoption |

Near-term decision:

- Fix the contract by choosing a canonical first indicator name before the next run. Prefer `Orgs Using AI in Talent Acquisition` for methodology v2.
- Do not silently backfill the current historical score with the display-only 51% value.
- Rename or replace `AI-Screened Applications`; the best locked 2026 source currently measures screening use-case adoption, not application share.
- Keep platform reach low-confidence until a refreshed multi-vendor annual assessment/application denominator is found.
- Add governance and human-override context, because hiring is a high-stakes domain where adoption does not imply autonomous decisioning.
