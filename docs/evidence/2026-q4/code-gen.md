# Software Development: 2026 Q4 Evidence Notes

Run `2026-q4`, method `delegation-curve-v2`. Prior notes: [2026 Q2](../2026-q2/code-gen.md) and [code-gen method](../../methodology/code-gen-2026.md).

Domain score 48.1 → **48.4**.

## Lock Decisions

| Indicator | Action | Value | Source | Window | Grade / confidence |
| --- | --- | ---: | --- | --- | --- |
| AI-Generated Code Output Share | hold (window fix) | 39.2% | Sonar 2026 (42%) and GitLab 2026 (34%) | Sonar fieldwork October 2025; GitLab published November 2025 | C / medium |
| Technical Work AI Value Share | hold | 50% | [METR AI usage survey](https://metr.org/blog/2026-05-11-ai-usage-survey/) | Feb-Apr 2026 | C / medium |
| AI Workflow Reliance | hold | 67.7% | DORA 2025 (65%) and JetBrains AI Pulse (74%) | 2025 to January 2026 | B / medium |
| Agentic Task Delegation | refresh | 19.8% | Anthropic range (held) and [arXiv:2606.07448](https://arxiv.org/abs/2606.07448) | GitHub traces January 2025 to early April 2026 | C / medium |
| Tool Ecosystem Reach | auto | 93.2M installs | VS Code Marketplace API | 2026-10-17 | D / medium |

## Notes

- **Agentic Task Delegation.** "Agentic Very Much!" (June 2026) is a rerun by the same authors of the GitHub-trace study used in 2026 Q2, on the same 127,670 older projects: 26.46% (conservative) to 32.89% (high) show coding-agent traces, midpoint 29.675%, up from 25.4%. The blend is 0.5 × 10% (Anthropic's fully delegable task range, held) + 0.5 × 29.675% = 19.8%. The 2026 Q2 ledger gave the earlier study's window as H1 2025; the paper covers January 2025 to February 2026.
- **AI Workflow Reliance.** No DORA 2026 report has been published. The JetBrains Developer Ecosystem 2026 survey (May-July 2026, more than 15,000 developers) finds [90% of professional developers using AI coding agents at work at least weekly](https://blog.jetbrains.com/research/2026/08/ai-coding-agent-adoption-2026/). That is a different construct from the AI Pulse adoption figure behind the 2026 Q2 input, so it is not substituted under the frozen method.
- **Output share is probably understated.** Newer surveys put AI-written code higher: JetBrains finds [about 47% of code written fully by agents](https://blog.jetbrains.com/research/2026/08/how-much-code-do-developers-really-let-agents-write/) (May-July 2026) and DX reports [51.9% AI-authored](https://newsletter.getdx.com/p/ai-authored-code-has-nearly-doubled) (Q2 2026, self-reported). Neither is a new edition of the Sonar or GitLab sources, and both measure narrower constructs than Sonar's "generated or assisted" code.
- **Tool Ecosystem Reach is decaying as a measure.** Of the five tracked extensions, Codeium is unpublished, Supermaven has shut down, and Tabnine is labelled legacy; Claude Code and Codex, now among the largest, are not tracked; GitHub Copilot is about 80% of the total and VS Code has announced it will leave the Marketplace. The collector now falls back to the last value if any tracked extension fails. Changing the list is a v3 method change.

## Context

- [Stack Overflow 2026 Developer Survey](https://survey.stackoverflow.co/2026/ai) (published 2026-10-06): 73% of respondents using AI coding assistants or agents use them daily; coding agents (66%) and chatbots (63%) are the main modes. Context only since 2026 Q2.

## Candidates For Methodology v3

- JetBrains' split of code into agent-written, AI-assisted, and manual shares.
- Anthropic's analysis of Claude Code sessions (Claude makes about 80% of execution decisions across about 400,000 sessions): <https://www.anthropic.com/research/claude-code-expertise>.
- Telemetry on pull requests authored by coding agents on GitHub (not yet checked).
