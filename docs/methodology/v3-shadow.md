# Methodology v3: Shadow Method

Status: computed in shadow from the 2026 Q4 run; not the headline method. Under the [restatement policy](restatement-policy.md), a decision on adopting it falls in the first run of 2027.

## Why

The v2 composite mixes two kinds of evidence. Some indicators measure **how much work AI does**: the share of content actioned automatically, of support cases handled by AI, of code written by AI, of loans underwritten without a human. Others measure **how widely AI is used**: the share of organizations that have adopted it, or of people who use it. Both matter, but only the first is delegation.

In the 2026 Q2 run, 46% of the composite weight came from share-of-work indicators, 50% from usage and adoption, and 4% from capacity counts. Three domains (medical diagnosis, legal, hiring) had no share-of-work indicator at all. `make diagnostics` reports this split for every run.

## What The Shadow Computes

Each indicator is tagged with the construct it measures (`Construct` in `internal/collect/score.go`):

| Construct | Meaning | Role in the shadow |
| --- | --- | --- |
| `work` | Share of work, decisions, cases, or content handled by AI | Delegation depth |
| `usage` | Share of people using AI | Reach |
| `adoption` | Share of organizations using AI | Reach |
| `capacity` | Counts or reach metrics (devices, installs, assessments) | Excluded |

For each domain the shadow reports a **depth** score (share-of-work indicators only) and a **reach** score (usage and adoption), each a weighted average of the existing v2 indicator weights. The shadow composite scores a domain as 70% depth plus 30% reach, using only domains that have both, and reports how much of the domain weight it covers.

The shadow also reports:

- **Evidence age:** the weighted median age of the data at publication and the share of weight on data more than 12 or 24 months old, from each value's recorded measurement window.
- **Uncertainty:** a reproducible Monte Carlo interval. Each indicator varies within a 90% range set by its confidence label (high ±5, medium-high ±7.5, medium ±10, low-medium ±12.5, low ±15 normalized points), and domain weights vary by lognormal noise (σ = 0.25). The ranges are explicit assumptions, not sampling error.
- **Sensitivity:** equal domain weights, grade A/B-only evidence, and excluding grade D or capacity metrics.

## Open Work Before Adoption

1. **Depth measures for medical, legal, and hiring.** Candidates include claims-based utilization of AI-specific procedure codes as a share of eligible studies, the share of e-discovery caseloads run through technology-assisted review, and independently audited applicant volumes screened by automated tools.
2. **A representative reach anchor.** The U.S. Census Bureau Business Trends and Outlook Survey reports AI use by sector, but its question changed in November 2025, so a series using it can only start there.
3. **Capacity metrics.** Cumulative counts with fixed caps (FDA devices over 0-2000, installs, assessment volume) rise mechanically and will saturate; replace them with shares or retire them.
4. **Proxy semantics.** FX electronic execution is not AI-driven execution; the algorithmic trading depth score needs an algorithmic or AI share of executed volume.
5. **Weights.** A written rationale for domain weights and for the depth/reach split.
