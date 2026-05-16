# 2026 Public Curve Series

Status: implemented for the public curve.
Date: 2026-05-16.

## Decision

The public Delegation Curve should remain one continuous curve. The March 2026 publication scores remain archived, but the public prior points are now recalculated under the 2026 scoring frame:

| Run ID | Period | Score | Public curve | Notes |
| --- | --- | ---: | --- | --- |
| `legacy-2024` | 2024 | 39.7 | no | Archived original scoring series |
| `legacy-2025` | 2025 | 46.0 | no | Archived March 2026 publication baseline |
| `current-method-2024` | 2024 | 31.5 | yes | Recalculated history |
| `current-method-2025` | 2025 | 37.7 | yes | Recalculated baseline |
| `2026-q2` | 2026 Q2 | 45.8 | yes | Current source refresh |

The headline comparison is therefore:

```text
45.8 - 37.7 = +8.1 index points
```

## Rationale

The prior public score of 46.0 was not a comparable denominator for the 2026 Q2 score because the 2026 update changed source contracts, especially in algorithmic trading. Showing `45.8` as `-0.2` against the March score implied a decline in AI delegation that the evidence did not support.

The correction is to preserve the old publication as an archived run and compare the 2026 score against a recalculated 2025 baseline under the current scoring frame.

## 2025 Recalculated Baseline

The 2025 recalculated baseline uses the 2026 indicator contracts where a prior-period value or close predecessor exists. Domain scores:

| Domain | 2025 comparable | Basis |
| --- | ---: | --- |
| content-mod | 94.1 | 2025 Meta, Google/YouTube, and TikTok automation values mapped to renamed v2 indicators |
| algo-trade | 39.2 | v2 construct correction: BIS FX electronic trading plus Coalition Greenwich internal-AI trade-execution adoption |
| code-gen | 32.6 | prior OSS/code/adoption values plus METR March 2025 value-share conversion |
| support | 37.0 | prior support automation signals mapped to the v2 cases/deflection/deployment frame |
| credit | 27.7 | prior loan automation plus bank decisioning values under the v2 two-input frame |
| medical-dx | 21.7 | prior FDA, radiology, diagnosis, and pathology values under the v2 frame |
| legal-ai | 29.0 | prior legal adoption and document-review values under the v2 frame |
| hire | 38.1 | prior recruiting/screening/platform values mapped to the v2 hiring frame |
| education | 18.0 | prior student, grading, and teacher-use values mapped to the v2 education frame |

Weighted composite:

```text
0.10*94.1 + 0.15*39.2 + 0.15*32.6 + 0.15*37.0 + 0.10*27.7
+ 0.12*21.7 + 0.08*29.0 + 0.08*38.1 + 0.07*18.0 = 37.7
```

## 2024 Comparable Estimate

Full source-level 2024 backcasting is not available for every v2 indicator. The 2024 public point is therefore a continuity backcast: it starts from the recalculated 2025 domain score and subtracts the original 2024-to-2025 domain movement.

Example:

```text
code-gen 2024 comparable = 32.6 - (46.6 - 41.0) = 27.0
```

This keeps the public curve continuous while making clear that the most defensible recalculation anchor is 2025, not 2024.

## Product Treatment

The homepage and domain detail pages should not explain methodology versioning in the hero. They should use a plain visitor-facing baseline label:

```text
45.8
+8.1 pts
vs 2025 baseline
```

The `2025 baseline` label refers to `current-method-2025`, the recalculated comparison point under the current scoring frame. The archive/methodology pages explain that earlier published scores are retained for audit but are not the current public comparison series.

Curve charts should mark each public analysis run. Prior public points use small outlined markers; the current run remains the stronger filled marker.
