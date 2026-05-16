# Code Generation 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `code-gen` score: 46.6.

Configured indicators:

- `Copilot Code Acceptance`: 50% weight.
- `Developer AI Tool Usage`: 15% weight.
- `AI-Assisted Commits (OSS)`: 25% weight.
- `IDE AI Extension Installs`: 10% weight.

The current contract mixes accepted-code telemetry, broad adoption, OSS proxy data, and install inventory. The 2026 refresh should move toward output/value share while retaining adoption and quality sources as context.

## Extracted Candidate Sources

### METR AI Usage Survey

Source: https://metr.org/blog/2026-05-11-ai-usage-survey/
Published: 2026-05-11.
Measurement period: February-April 2026 survey, with retrospective March 2025 and March 2026 estimates.
Evidence grade: C.
Confidence: medium.

Relevant values:

- Survey sample: 349 technical workers.
- Median self-reported value uplift: 1.4x to 2x.
- Retrospective median value estimate: 1.3x for March 2025.
- Current retrospective median value estimate: 2x for March 2026.
- Forecast median value estimate: 2.5x for March 2027.

Candidate conversion:

```text
AI-attributable technical work value share = (multiplier - 1) / multiplier
```

This gives:

- March 2025: 1.3x -> 23.1%.
- March 2026: 2.0x -> 50.0%.
- 2026 median range: 1.4x to 2.0x -> 28.6% to 50.0%.

Recommendation: add as `Technical Work AI Value Share`, not as a replacement for accepted-code telemetry. It is closer to delegation/value share than ordinary adoption, but it remains self-reported and selected toward technical workers.

### Sonar 2026 State of Code Developer Survey

Source: https://www.sonarsource.com/blog/state-of-code-developer-survey-report-the-current-reality-of-ai-coding
Published: 2026-01-08.
Measurement period: 2026 survey.
Evidence grade: C.
Confidence: medium.

Relevant values:

- Survey sample: more than 1,100 professional developers.
- AI accounts for 42% of committed code today.
- Expected AI share reaches 65% by 2027.
- 96% of developers do not fully trust AI-generated code.
- 48% always verify AI-generated code before committing.
- 38% report that reviewing AI code takes more effort than reviewing human-written code.

Recommendation: add as `AI-Generated or Assisted Committed Code`, with a confidence caveat. This maps more directly to production output than broad AI-tool usage, but it is still a survey report from a developer-tool vendor rather than observed repository telemetry.

### Stack Overflow 2025 Developer Survey

Source: https://survey.stackoverflow.co/2025/ai
Published: 2025.
Measurement period: 2025.
Evidence grade: D.
Confidence: medium.

Relevant values:

- 84% of respondents use or plan to use AI tools in development.
- 51% of professional developers use AI tools daily.
- Professional developers: 50.6% daily, 17.4% weekly, 12.8% monthly or infrequently.

Recommendation: keep as adoption context or redefine the current indicator as `Professional Developer Daily AI Use`. Do not use Stack Overflow as an output-share measure.

### GitHub Octoverse 2025

Source: https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/
Published: 2025-11.
Measurement period: September 2024-August 2025.
Evidence grade: D for adoption context; B if used only as platform workflow penetration.
Confidence: medium.

Relevant values:

- More than 1.1 million public repositories use an LLM SDK.
- 4.3 million AI-related projects exist on GitHub.
- 80% of new developers on GitHub use Copilot within their first week.
- 518.7 million pull requests were merged in 2025.

Recommendation: this is strong context that AI is now default in GitHub workflows, but it does not refresh `Copilot Code Acceptance`. Keep looking for a first-party acceptance or accepted-lines metric before preserving that indicator in the scoring contract.

### GitClear AI Copilot Code Quality Research 2025

Source: https://gitclear-public.s3.us-west-2.amazonaws.com/GitClear-AI-Copilot-Code-Quality-2025.pdf
Published: 2025-02-04.
Measurement period: code changes through 2024.
Evidence grade: D for context.
Confidence: medium.

Relevant values:

- Analyzes 211 million changed lines.
- Newly added code rose to 46.2% of changed lines in 2024.
- Churn rose to 5.7% in 2024.
- Copy/pasted line operations rose to 12.3% in 2024.

Recommendation: use as a quality and maintenance guardrail. It is not a direct measure of AI-assisted commit share and should not remain a primary scoring input unless a clearer AI-attribution method is available.

## Proposed Code-Gen Source Lock

Proposed v2 scoring candidates:

| Indicator | Suggested role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| AI-Generated or Assisted Committed Code | score input | C | medium | Sonar 42% committed-code survey estimate |
| Technical Work AI Value Share | score input | C | medium | METR 50% converted March 2026 value share |
| Copilot accepted-code telemetry | score input if refreshed | A/B | medium | keep only if first-party accepted-code metric is found |
| Professional Developer Daily AI Use | context or low-weight score input | D | medium | Stack Overflow 50.6% daily professional developer use |
| IDE AI Extension Installs | context or low-weight score input | D | medium | continuously refreshable but weak denominator |
| GitClear code quality/churn | context only | D | medium | quality guardrail rather than delegation share |

Near-term decision:

- Do not update the `code-gen` score yet.
- Treat METR and Sonar as strong candidates for a methodology-v2 run.
- Replace or demote `Developer AI Tool Usage` because the current value is being used as a delegation/output proxy even though the source is adoption-oriented.
- Replace or demote `AI-Assisted Commits (OSS)` unless a clear AI-attributed commit-share source is locked.
