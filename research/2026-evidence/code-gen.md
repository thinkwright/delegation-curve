# Code Generation 2026 Evidence Extraction

Status: implemented for the 2026 Q2 source refresh.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `code-gen` score: 48.1.

Configured indicators:

- `AI-Generated Code Output Share`: 30% weight.
- `Technical Work AI Value Share`: 25% weight.
- `AI Workflow Reliance`: 25% weight.
- `Agentic Task Delegation`: 15% weight.
- `Tool Ecosystem Reach`: 5% weight.

The current contract moved away from stale accepted-code telemetry, OSS proxy data, and broad daily-use adoption. It now separates output share, value share, workflow reliance, and agentic task delegation.

Current calculation:

```text
0.30 * 39.2 AI-Generated Code Output Share
+ 0.25 * 50.0 Technical Work AI Value Share
+ 0.25 * 67.7 AI Workflow Reliance
+ 0.15 * 17.7 Agentic Task Delegation
+ 0.05 * 85.0 Tool Ecosystem Reach normalized
= 48.1
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

Recommendation: use in the `AI-Generated Code Output Share` blend with GitLab, with a confidence caveat. This maps more directly to production output than broad AI-tool usage, but it is still a survey report from a developer-tool vendor rather than observed repository telemetry.

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

Recommendation: keep as historical bridge adoption context only. Do not use Stack Overflow as an output-share or workflow-reliance measure.

Follow-up recommendation: continue tracking Stack Overflow as context, but keep it out of scoring unless the survey adds a more direct reliance, output, or delegation metric.

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

### Anthropic 2026 Agentic Coding Trends Report

Source: https://resources.anthropic.com/hubfs/2026%20Agentic%20Coding%20Trends%20Report.pdf?hsLang=en
Published: 2026.
Measurement period: 2025/2026 customer and internal research signals.
Evidence grade: C for quantitative scoring; B for construct definition.
Confidence: medium.

Relevant values:

- Anthropic's Societal Impacts research says developers use AI in roughly 60% of their work.
- The same report says developers report being able to "fully delegate" only 0-20% of tasks.
- The report frames the role shift as engineers orchestrating implementation agents while retaining architecture, validation, and strategic judgment.
- Anthropic reports customer cases including Rakuten using Claude Code for a seven-hour autonomous vLLM task and CRED doubling execution speed.

Recommendation: use the assisted-vs-fully-delegated split as a methodology calibration point. It is directly aligned with the Delegation Curve concept because it distinguishes AI involvement from actual delegation. Do not use the customer examples as population score inputs.

### Alphabet Internal AI Code Generation Benchmark

Source: https://abc.xyz/investor/events/event-details/2025/2025-Q3-Earnings-Call-2025-4OI4Bac_Q9/default.aspx
Published: 2025-10-29.
Measurement period: Q3 2025.
Evidence grade: B for company benchmark; D for population measurement.
Confidence: medium.

Relevant values:

- Alphabet's Q3 2025 transcript says nearly half of all code was generated by AI.

Recommendation: use as a high-confidence top-tech benchmark that the population score should be directionally compared against. It should not replace broad developer measures because Alphabet is an unusually AI-native engineering environment.

### GitLab 2026 Global DevSecOps Report

Source: https://about.gitlab.com/resources/developer-survey/
Published/accessed: 2026.
Measurement period: 2026 survey.
Evidence grade: B.
Confidence: medium.

Relevant values:

- Survey sample: 3,266 DevSecOps professionals across countries and industries.
- GitLab's public report page shows code source share as 34% AI-generated, 37% written from scratch, and 29% copied from other sources.
- The report explicitly asks about the current split between human and AI contributions in software development.

Recommendation: add as `AI-Generated Code Share Triangulation`, not an immediate replacement for Sonar. GitLab's 34% is a useful lower comparator to Sonar's 42% and Augment's AI-forward 48% estimate.

### Agentic Much? Adoption of Coding Agents on GitHub

Source: https://arxiv.org/abs/2601.18341
Published: 2026-01-26; revised 2026-04-08.
Measurement period: first half of 2025 GitHub traces.
Evidence grade: B.
Confidence: medium.

Relevant values:

- Study scope: 128,018 GitHub projects.
- Estimated coding-agent adoption rate: 22.20% to 28.66%.
- Coding agents are identified through explicit software-engineering artifact traces such as co-authored commits or pull requests.
- Agent-assisted commits are larger than human-only commits and include a high proportion of features and bug fixes.

Recommendation: use as a prototype `OSS Coding-Agent Adoption` measure or as context for agentic adoption. It is more behaviorally grounded than surveys, but it is open-source/GitHub-specific and likely undercounts private enterprise adoption.

### Augment Code State of AI-Native Engineering 2026

Source: https://www.augmentcode.com/blog/ai-native-survey-2026
Published: 2026-05-12.
Measurement period: 2026 survey.
Evidence grade: C.
Confidence: low-medium.

Relevant values:

- Survey sample: 219 engineering leaders.
- Sample caveat: Augment states the sample skews AI-forward.
- Respondents estimate 48% of their code is AI-generated.
- 55% are concerned or very concerned about losing shared understanding of the codebase.
- Only 19 of 219 organizations had updated role definitions to match the changed engineering job.

Recommendation: use as a frontier/adopter context source, not a central score input. The 48% estimate is useful because it triangulates Sonar and GitLab, but the sample is intentionally AI-forward.

### Harness State of Engineering Excellence 2026

Source: https://www.harness.io/state-of-engineering-excellence
Published: 2026.
Measurement period: 2026 survey.
Evidence grade: B for guardrail; D for delegation score.
Confidence: medium.

Relevant values:

- Survey sample: 700 developers and engineering leaders from large enterprises across the US, UK, France, Germany, and India.
- Harness states developers are becoming validators of machine-generated output.
- 81% of engineering leaders report code-review time increased after deploying AI.
- 31% of a developer's day is now consumed by AI-related invisible work.
- Top AI friction points: reviewing AI code for accuracy at 53%, fixing subtle bugs from AI code at 52%, explaining AI code to teammates at 48%, and context switching at 45%.

Recommendation: add as `AI Validation Tax` context. This should not raise the code-gen score, but it should protect the methodology from treating generated-code volume as net delegated work.

### Lightrun State of AI-Powered Engineering 2026

Source: https://lightrun.com/ebooks/state-of-ai-powered-engineering-2026/
Published: 2026.
Measurement period: 2026 survey.
Evidence grade: C for guardrail; D for delegation score.
Confidence: medium-low.

Relevant values:

- Survey sample: 200 SRE and DevOps leaders at enterprises in the US, UK, and EU.
- Lightrun states AI adoption in software engineering has reached 90%.
- 43% of AI-generated code still requires manual debugging in production after passing QA and staging.
- 88% of organizations need two to three redeploy cycles to publish a single AI-generated change.

Recommendation: use as runtime-reliability guardrail only. The source is enterprise/SRE-skewed and vendor-reported, but it captures an important downstream cost of agentic coding.

### CircleCI 2026 State of Software Delivery

Source: https://circleci.com/blog/five-takeaways-2026-software-delivery-report/
Published: 2026-02.
Measurement period: September 2025 workflow data.
Evidence grade: B for delivery guardrail.
Confidence: medium.

Relevant values:

- Analysis scope: 28,738,317 CircleCI workflows from projects with at least two contributors and at least five workflow runs.
- Median feature-branch throughput increased 15%, while median main-branch throughput fell 7%.
- Main-branch success rates dropped to 70.8%, the lowest in over five years.
- The top 5% of teams nearly doubled throughput and increased main-branch throughput 26%.

Recommendation: use as a delivery-system guardrail. It is not a code-generation adoption source, but it directly supports the interpretation that AI code volume can outpace validation and production delivery.

### Faros AI Engineering Report 2026

Source: https://www.faros.ai/blog/ai-acceleration-whiplash-takeaways
Published: 2026-04-12.
Measurement period: two years of telemetry.
Evidence grade: B for telemetry guardrail; C because source is vendor platform data.
Confidence: medium.

Relevant values:

- Analysis scope: telemetry from 22,000 developers and more than 4,000 teams.
- 80% of studied teams exceed a 50% weekly-active-user threshold for AI tools.
- AI-generated-code acceptance rate rose from 20% to 60%.
- Epics completed per developer rose 66%, task throughput rose 33.7%, and PR merge rate rose 16.2%.
- Code churn rose 861%, incidents-to-PR ratio rose 242.7%, and bugs per developer rose 54%.
- PRs merged without review rose 31.3%.

Recommendation: use as a high-value guardrail and possible confidence-band source. It should not directly raise the score because the telemetry comes from Faros customers, but it strengthens the case for reporting net delegation with validation overhead.

### Debt Behind the AI Boom

Source: https://arxiv.org/abs/2603.28592
Published: 2026-03-30; revised 2026-04-26.
Measurement period: GitHub repositories containing verified AI-authored commits.
Evidence grade: B for quality guardrail.
Confidence: medium.

Relevant values:

- Dataset: 302.6k verified AI-authored commits from 6,299 GitHub repositories across five AI coding assistants.
- Static analysis identified 484,366 AI-introduced issues.
- Code smells accounted for 89.3% of identified issues.
- More than 15% of commits from every studied AI coding assistant introduced at least one issue.
- 22.7% of tracked AI-introduced issues still survived at the latest repository version.

Recommendation: use as a quality and maintenance guardrail. It is not a delegation-share measure but is highly relevant to confidence intervals and methodology notes.

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

## Implemented Code-Gen Source Lock

Current v2 scoring inputs and retained candidates:

| Indicator | Role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| AI-Generated Code Output Share | score input | B/C | medium | Sonar 42% committed-code estimate blended with GitLab 34% code-source-share estimate; Augment 48% retained as AI-forward upper context |
| Technical Work AI Value Share | score input | C | medium | METR 50% converted March 2026 value share |
| AI Workflow Reliance | score input | B | medium | DORA 65% moderate-to-heavy reliance blended with JetBrains 74% specialized AI developer-tool adoption |
| Agentic Task Delegation | score input | B/C | medium | Anthropic fully delegable task midpoint blended with arXiv GitHub-trace coding-agent adoption midpoint |
| Tool Ecosystem Reach | score input | D | medium | continuously refreshable VS Code Marketplace stock metric, kept at low weight |
| Assisted vs Fully Delegated Work | methodology calibration | C/B | medium | Anthropic reports AI used in roughly 60% of developer work but only 0-20% fully delegable |
| Frontier lab/company benchmark | context or annotation | B/D | medium | Anthropic capability claims and Alphabet/Augment nearly-half AI-generated code |
| AI Validation Tax | guardrail or confidence modifier | B/C | medium | Harness, Lightrun, CircleCI, Faros, and arXiv quality studies show review, debugging, CI, and maintenance costs |
| Copilot accepted-code telemetry | score input if refreshed | A/B | medium | keep only if first-party accepted-code metric is found |
| Professional Developer Daily AI Use | retired from scoring | D | medium | Stack Overflow 50.6% daily professional developer use; self-selected adoption proxy |
| GitClear code quality/churn | context only | D | medium | quality guardrail rather than delegation share |

Implementation decision:

- Update the score contract and current `code-gen` score to 48.1.
- Treat Sonar and GitLab as the current output-share blend; keep Augment as AI-forward context.
- Keep METR as the technical-work value-share input.
- Replace Stack Overflow with DORA and JetBrains workflow-reliance evidence.
- Add a separate agentic-delegation input using Anthropic's assisted-vs-fully-delegated split and the arXiv GitHub-trace study.
- Use Harness, Lightrun, CircleCI, Faros, and AI-code-quality studies as guardrails or uncertainty modifiers, not positive score inputs.
- Use Anthropic, Alphabet, Augment, and similar high-intensity examples as frontier benchmarks or narrative callouts unless a representative denominator is available.
- Continue looking for first-party accepted-code, accepted-lines, or agent-authored-PR telemetry before restoring a direct GitHub/Copilot score input.
