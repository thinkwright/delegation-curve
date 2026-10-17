# Domain Weights

The composite is a weighted average of nine domain scores. Within each domain, indicator weights sum to 1; across domains, weights sum to 1. Both are enforced by tests in `internal/collect/score_test.go`, and the configuration lives in `internal/collect/score.go`.

| Domain | Weight | Tier |
| --- | ---: | ---: |
| Algorithmic trading | 15% | 1 |
| Software development | 15% | 1 |
| Customer support | 15% | 1 |
| Content moderation | 10% | 1 |
| Medical diagnosis | 12% | 2 |
| Credit and lending | 10% | 2 |
| Legal research and review | 8% | 2 |
| Recruitment and screening | 8% | 2 |
| Education and assessment | 7% | 3 |

## How To Read Them

The weights are an editorial choice, not an estimate from data. Tiers group domains broadly by the maturity of their evidence: tier 1 domains draw on platform transparency reports, market-structure data, and operational measures; tier 2 and 3 domains rely more on surveys and proxies.

## Robustness

The headline does not depend much on the weights. For the 2026 Q2 run, equal weights give 46.4 against the published 45.8, and the uncertainty interval, which also perturbs the weights, is 42.8 to 48.9. `make diagnostics` reports these figures for every run, and the methodology page lets readers re-weight the domains themselves.

A written rationale for each weight is part of the 2027 methodology review (see [v3-shadow.md](v3-shadow.md)).
