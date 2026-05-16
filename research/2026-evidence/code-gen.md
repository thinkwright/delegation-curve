# Code Generation 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `code-gen` score: 48.2.

Configured indicators:

- `AI-Generated or Assisted Committed Code`: 45% weight.
- `Technical Work AI Value Share`: 35% weight.
- `Professional Developer Daily AI Use`: 15% weight.
- `IDE AI Extension Installs`: 5% weight.

The current contract moved away from stale accepted-code telemetry and OSS proxy data toward output and value-share measures. It still uses Stack Overflow as a low-weight adoption proxy, which is directionally useful but biased toward self-selected developer-survey respondents and not ideal for measuring engineering-work reliance.

Current calculation:

```text
0.45 * 42.0 AI-Generated or Assisted Committed Code
+ 0.35 * 50.0 Technical Work AI Value Share
+ 0.15 * 50.6 Professional Developer Daily AI Use
+ 0.05 * 85.0 IDE AI Extension Installs normalized
= 48.2
```

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

Follow-up recommendation: demote Stack Overflow further once a stronger workflow-reliance source is locked. Stack Overflow is useful as a large developer-community survey, but the population is self-selected and the metric is adoption/frequency rather than delegation share.

### Google DORA 2025 State of AI-Assisted Software Development

Source: https://blog.google/innovation-and-ai/technology/developers-tools/dora-report-2025/
Published: 2025-09-23.
Measurement period: 2025 survey.
Evidence grade: B.
Confidence: medium.

Relevant values:

- Survey sample: nearly 5,000 technology professionals globally.
- AI adoption among software development professionals: 90%.
- Median AI use: two hours per day.
- Moderate-to-heavy reliance: 65% total, composed of 37% moderate, 20% a lot, and 8% a great deal.
- More than 80% report productivity gains.
- 59% report positive code-quality influence.

Recommendation: add as a candidate `AI Workflow Reliance` or `Agentic Engineering Reliance` indicator. This is a better replacement candidate for Stack Overflow than another broad adoption survey because it directly measures reliance and work time, not only whether respondents use or plan to use AI tools.

### JetBrains AI Pulse 2026

Source: https://blog.jetbrains.com/research/2026/04/which-ai-coding-tools-do-developers-actually-use-at-work/
Published: 2026-04.
Measurement period: January 2026 survey, with earlier September 2025 and Developer Ecosystem comparison waves.
Evidence grade: B.
Confidence: medium.

Relevant values:

- Survey sample: more than 10,000 professional developers worldwide.
- 90% of developers regularly used at least one AI tool at work for coding and development tasks in January 2026.
- 74% had adopted specialized AI tools for developers, excluding general chatbots.
- Claude Code work adoption reached 18% worldwide and 24% in the US and Canada.
- ChatGPT chatbot use for coding and development tasks at work was 28%.

Recommendation: add as a candidate triangulation source for `AI Workflow Reliance` and `Agentic Coding Tool Adoption`. This is stronger than Stack Overflow for current work usage because it separates developer-specific agents/editors from general chatbots and includes methodology notes about debranded survey promotion and weighting.

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
- GitHub notes that record activity coincided with the Copilot coding agent preview and Copilot code review rollout, while explicitly treating these as observational rather than causal signals.

Recommendation: this is strong context that AI is now default in GitHub workflows, but it does not refresh `Copilot Code Acceptance`. Keep looking for a first-party acceptance or accepted-lines metric before preserving that indicator in the scoring contract.

Follow-up recommendation: keep GitHub as a workflow-platform context source rather than a direct population score unless the denominator for agent-authored pull requests is made explicit enough for scoring.

### Anthropic Claude Code Product and Documentation

Sources:

- https://claude.com/product/claude-code?hsLang=en
- https://code.claude.com/docs/en/how-claude-code-works

Published/accessed: 2026.
Measurement period: product/documentation claims, not population survey.
Evidence grade: B for capability shape; D for population measurement.
Confidence: medium.

Relevant values and signals:

- Claude Code is positioned as an agentic terminal and IDE system that reads issues, writes code, runs tests, and submits PRs.
- Anthropic documentation describes an agentic loop of gathering context, taking action, and verifying results.
- Claude Code product claims include 7.6x more frequent deployments by Claude Code teams.
- The product and documentation directly support measuring broader engineering delegation beyond committed-code volume.

Recommendation: use as a frontier/capability benchmark, not as a scored population input. High-profile Anthropic usage claims are important for the story, but without a representative denominator they should inform a frontier-band annotation or narrative callout rather than the central score.

### Alphabet Internal AI Code Generation Benchmark

Source: https://abc.xyz/investor/events/event-details/2025/2025-Q3-Earnings-Call-2025-4OI4Bac_Q9/default.aspx
Published: 2025-10-29.
Measurement period: Q3 2025.
Evidence grade: B for company benchmark; D for population measurement.
Confidence: medium.

Relevant values:

- Alphabet's Q3 2025 transcript says nearly half of all code was generated by AI.

Recommendation: use as a high-confidence top-tech benchmark that the population score should be directionally compared against. It should not replace broad developer measures because Alphabet is an unusually AI-native engineering environment.

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
| AI Workflow Reliance | prototype score input | B | medium | DORA 65% moderate-to-heavy reliance plus median two hours per day |
| Agentic Coding Tool Adoption | prototype score input or context | B | medium | JetBrains 74% specialized AI dev-tool adoption and 18% Claude Code work adoption |
| Frontier lab/company benchmark | context or annotation | B/D | medium | Anthropic capability claims and Alphabet nearly-half AI-generated code |
| Copilot accepted-code telemetry | score input if refreshed | A/B | medium | keep only if first-party accepted-code metric is found |
| Professional Developer Daily AI Use | retire or keep as low-weight bridge | D | medium | Stack Overflow 50.6% daily professional developer use; self-selected adoption proxy |
| IDE AI Extension Installs | context or low-weight score input | D | medium | continuously refreshable but weak denominator |
| GitClear code quality/churn | context only | D | medium | quality guardrail rather than delegation share |

Near-term decision:

- Keep the 48.2 `code-gen` score unchanged until a score method change is explicitly approved.
- Treat METR and Sonar as the current output/value-share core.
- Prototype replacing Stack Overflow with DORA `AI Workflow Reliance`, optionally triangulated with JetBrains AI Pulse.
- Use Anthropic, Alphabet, and similar top-tech examples as frontier benchmarks or narrative callouts unless a representative denominator is available.
- Continue looking for first-party accepted-code, accepted-lines, or agent-authored-PR telemetry before restoring a direct GitHub/Copilot score input.
