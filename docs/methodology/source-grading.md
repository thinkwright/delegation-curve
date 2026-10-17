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

Every score update keeps source review, scoring, run snapshotting, and generated artifacts separate. Run history (previous scores, deltas, trends) is derived from explicit analysis runs and is never edited by hand.

1. **Lock sources.** Update each value in `seed/overrides.yaml` together with its audit trail (see below) and record the decision in the run's source ledger (`docs/runs/<run>-source-ledger.csv`).
2. **Review the change set.** `go run ./cmd/collect -dry-run` prints every indicator, domain, and composite change without writing anything. Explain any domain move over 5 points in the run record.
3. **Collect.** `make collect` writes the new current values and scores into `seed/seed.json`. It does not touch history.
4. **Snapshot the run.** `curve-snapshot-run -run-id <run> -label ... -measurement-period ... -measurement-year ... -published-at YYYY-MM-DD -methodology-version ...` records the run and rewrites previous scores, deltas, trends, and `last_updated` from the public run series.
5. **Generate.** `make generate` writes Parquet, `meta.json`, `sitemap.xml`, and the static seed copy. It refuses to publish scores that are not snapshotted in the current run, or history fields that disagree with the run series.
6. **Review diagnostics.** `make diagnostics` reports composition, evidence age, sensitivity, the uncertainty interval, and the v3 shadow. Include the output in the run record.
7. **Refresh the share image.** `make og-image` renders `frontend/static/og-image.png` from the new data (requires Chrome).
8. **Build and check.** `make server`, then `make run-server` and `make smoke` in another shell. CI repeats the reproducibility, generated-data, and smoke checks on every pull request.
9. **Record the run** under `docs/runs/` and `docs/evidence/<run>/`.

## Audit Trail For Every Value

Each override in `seed/overrides.yaml` records:

| Field | Meaning |
| --- | --- |
| `source_url` | Primary source supporting the value. May be empty only for a low-confidence held value, which must explain itself in `transform`. |
| `published_at` | Publication date of the source. |
| `measurement_start`, `measurement_end` | The period the data describes (fieldwork, reporting quarter, filing year). |
| `evidence_grade` | A-D, per the table above. |
| `confidence` | high, medium-high, medium, low-medium, or low. |
| `transform` | Any conversion, blend, or denominator adjustment used to derive the value. |

Dates are ISO 8601 at the precision the source discloses: `2025`, `2025-09`, or `2025-09-30`. When fieldwork dates are not disclosed, `measurement_end` is the publication date (an upper bound) and `measurement_start` is left empty. Ages and staleness are computed from the start of a reduced-precision period, so they are never understated.

A data-lint test enforces these fields, and the metadata is exported with every run in `indicator_observations` and `sub_indicators`.

## Reproducibility Notes

- `go run ./cmd/collect -dry-run -offline -check` must reproduce the published scores from `seed/overrides.yaml` and cached values alone; CI runs it on every change.
- Dynamic reports and APIs should be extracted reproducibly where possible. If an API requires access tokens or rolling retention windows, the run should record enough detail (query, window, checksum) to reproduce the published value later.
- Automated collectors must fail rather than return a partial value: an indicator defined over a fixed set of sources falls back to its last recorded value if any source is unavailable.
