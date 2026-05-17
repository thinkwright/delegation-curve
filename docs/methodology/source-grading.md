# Source Grading And Refresh Workflow

This document defines the source-quality rubric and the refresh workflow used for the 2026 Q2 Delegation Curve run.

## Evidence Grades

| Grade | Meaning | Use in scoring |
| --- | --- | --- |
| A | Direct observed delegated decisions, task resolutions, operational outcomes, regulatory disclosures, or platform telemetry | Preferred score input |
| B | Workflow penetration, operational deployment, official databases, or strong representative surveys | Score input when direct outcome evidence is unavailable |
| C | AI-attributable output, value, productivity share, or narrower proxy evidence | Useful with caveats and confidence notes |
| D | Adoption, sentiment, capability benchmark, vendor market estimate, or broad context | Context only unless explicitly converted with a documented rule |

## Source Roles

Sources can have different roles in the public record:

- **Scored input:** directly contributes to a domain score.
- **Denominator:** constrains a proxy so a platform-specific metric is not overstated as a whole-market metric.
- **Context:** helps explain adoption pressure, deployment conditions, or market structure.
- **Guardrail:** narrows interpretation or confidence, especially when a positive metric can overstate actual delegation.
- **Held source:** retained from a prior locked source because no stronger current source exists.
- **Candidate source:** promising but not scored until extraction, denominator, or reproducibility requirements are met.

## Refresh Workflow

Future score updates should keep source review, scoring, run snapshotting, and generated artifacts separate.

1. Update source evidence in `seed/overrides.yaml` or collectors.
2. Record source decisions, grades, confidence, measurement periods, and URLs in the source ledger.
3. Run the collector to update `seed/seed.json`.
4. Snapshot the run with `curve-snapshot-run`.
5. Regenerate static data with `make generate`.
6. Rebuild and test the frontend/server.
7. Preserve the run record and evidence notes under `docs/`.

Do not manually edit current scores in `seed/seed.json` without creating or updating an explicit run snapshot.

## Reproducibility Notes

Manual override values must include enough information to audit the source:

- source name,
- URL,
- measurement period,
- freshness label,
- raw value,
- transformation or normalization rule,
- confidence caveat when applicable.

Dynamic reports and APIs should be extracted reproducibly where possible. If an API requires access tokens or rolling retention windows, the run should snapshot enough detail to reproduce the published value later.
