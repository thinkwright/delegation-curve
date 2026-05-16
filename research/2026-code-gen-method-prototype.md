# Code Generation vNext Method Update

Status: implemented for the 2026 Q2 code-gen score.
Prepared: 2026-05-16.
Current public `code-gen` score: 48.1.

## Purpose

The 2026 source sweep found credible evidence that AI-assisted engineering reliance is higher than the current scoring contract captures. The evidence also shows that generated-code volume, daily AI usage, workflow reliance, and actual task delegation are different constructs.

This note records the replacement scoring frame implemented after source review. The update changes the code-gen methodology and current domain score from 48.2 to 48.1; the composite remains 45.8 after rounding.

## Prior Contract

Prior calculation:

```text
0.45 * 42.0 AI-Generated or Assisted Committed Code
+ 0.35 * 50.0 Technical Work AI Value Share
+ 0.15 * 50.6 Professional Developer Daily AI Use
+ 0.05 * 85.0 IDE AI Extension Installs normalized
= 48.2
```

The prior contract was directionally defensible, but it had one weak seam: Stack Overflow daily AI use was an adoption-frequency proxy, not a workflow-reliance or delegation measure.

## Measurement Constructs

| Construct | Question answered | Candidate sources | Prototype role |
| --- | --- | --- | --- |
| Output share | What share of code or committed code is AI-generated or AI-assisted? | Sonar 2026, GitLab 2026, Augment 2026 | Core score input |
| Technical value share | What share of technical work value is attributed to AI assistance? | METR 2026 | Core score input |
| Workflow reliance | How much ordinary engineering work now depends on AI tools? | DORA 2025, JetBrains AI Pulse 2026 | Replace Stack Overflow daily use |
| Agentic delegation | How much work is actually delegated to coding agents rather than merely assisted by AI? | Anthropic 2026 Agentic Coding Trends, arXiv GitHub coding-agent trace study | New explicit delegation input |
| Tool ecosystem reach | Is the tooling ecosystem broadly available and continuously observable? | VS Code Marketplace | Low-weight continuity input |
| Validation tax | How much review, debugging, CI, and maintenance overhead offsets generated-code volume? | Harness 2026, Lightrun 2026, CircleCI 2026, Faros 2026, AI-code technical-debt research | Guardrail and uncertainty note, not a positive score input |

## Implemented Formula

```text
0.30 * AI-Generated Code Output Share
+ 0.25 * Technical Work AI Value Share
+ 0.25 * AI Workflow Reliance
+ 0.15 * Agentic Task Delegation
+ 0.05 * Tool Ecosystem Reach
```

Rationale:

- Output share remains important, but it should not dominate the domain because generated-code volume can still require substantial human review.
- METR stays in the formula because value share is closer to the Delegation Curve concept than adoption.
- Workflow reliance deserves a larger role than Stack Overflow daily usage because DORA and JetBrains measure professional work dependence more directly.
- Agentic delegation is added explicitly so the score does not confuse AI assistance with true delegated engineering.
- Tool reach remains low weight because installs are continuously refreshable but denominator-poor.

## Scenario Inputs

| Input | Conservative | Base | High | Notes |
| --- | ---: | ---: | ---: | --- |
| AI-generated code output share | 34.0 | 39.2 | 48.0 | Conservative uses GitLab 34%; base blends Sonar 42% at 65% and GitLab 34% at 35%; high uses Augment's AI-forward 48% as an upper context value. |
| Technical work AI value share | 50.0 | 50.0 | 50.0 | METR March 2026 2x value multiplier converted as `(2 - 1) / 2 = 50%`. |
| AI workflow reliance | 65.0 | 67.7 | 90.0 | Conservative uses DORA moderate-to-heavy reliance; base blends DORA 65% at 70% and JetBrains specialized developer-tool adoption 74% at 30%; high uses DORA/JetBrains broad professional AI use around 90%. |
| Agentic task delegation | 10.0 | 17.7 | 28.7 | Conservative uses midpoint of Anthropic's 0-20% fully delegable task range; base blends that midpoint with the arXiv GitHub coding-agent adoption midpoint of 25.4%; high uses the arXiv upper estimate of 28.66%. |
| Tool ecosystem reach | 85.0 | 85.0 | 85.0 | Existing normalized VS Code Marketplace continuity value. |

Base output-share calculation:

```text
0.65 * 42.0 Sonar committed-code estimate
+ 0.35 * 34.0 GitLab code-source-share estimate
= 39.2
```

Base workflow-reliance calculation:

```text
0.70 * 65.0 DORA moderate-to-heavy reliance
+ 0.30 * 74.0 JetBrains specialized AI developer-tool adoption
= 67.7
```

Base agentic-delegation calculation:

```text
0.50 * 10.0 Anthropic fully delegable task midpoint
+ 0.50 * 25.4 arXiv GitHub coding-agent adoption midpoint
= 17.7
```

## Scenario Scores

| Scenario | Calculation | Score |
| --- | --- | ---: |
| Conservative | `0.30*34.0 + 0.25*50.0 + 0.25*65.0 + 0.15*10.0 + 0.05*85.0` | 44.7 |
| Base | `0.30*39.2 + 0.25*50.0 + 0.25*67.7 + 0.15*17.7 + 0.05*85.0` | 48.1 |
| High | `0.30*48.0 + 0.25*50.0 + 0.25*90.0 + 0.15*28.66 + 0.05*85.0` | 58.0 |

## Interpretation

The prior 48.2 score did not miss enough representative high-value evidence to require an upward correction. The stronger 2026 evidence mainly changes the semantics of the method:

- AI-assisted engineering is now widespread enough that Stack Overflow daily use should be retired from the scored contract.
- DORA and JetBrains justify treating workflow reliance as a first-class code-gen construct.
- Anthropic and the arXiv GitHub trace study justify a separate agentic-delegation construct.
- Frontier examples such as Claude Code, Anthropic internal use, Alphabet's nearly-half generated-code benchmark, and Augment's AI-forward survey should shape narrative and upper bands, not the central population estimate.
- Validation-tax sources should narrow confidence and interpretation, not mechanically lower the score before a cross-domain uncertainty method exists.

The implemented score is almost unchanged at 48.1 because broad workflow reliance is counterbalanced by the still-lower evidence for fully delegated agentic tasks. That is a useful result: it says the prior score was not obviously too low, but the new method is more semantically honest.

## Implementation Decision

The scoring contract is updated to the five-input formula above. Stack Overflow is retired as a scored input and retained only as historical bridge context. Validation-tax evidence remains confidence context rather than a direct score penalty.

Implemented changes:

- Added or renamed indicators:
  - `AI-Generated Code Output Share`
  - `Technical Work AI Value Share`
  - `AI Workflow Reliance`
  - `Agentic Task Delegation`
  - `Tool Ecosystem Reach`
- Updated `internal/collect/score.go` code-gen weights.
- Updated `seed/seed.json` current indicators and `analysis_runs`.
- Regenerated Parquet/static data and rebuilt the frontend/server.
