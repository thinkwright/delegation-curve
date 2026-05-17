# AI Delegation Curve

**How much does AI influence work and society?**

The AI Delegation Curve is a public measurement project tracking AI influence across consequential domains of work and civic life. It turns scattered adoption signals, transparency reports, surveys, public datasets, and research findings into one time-series index.

**Live report:** [curve.thinkwright.ai](https://curve.thinkwright.ai)

## Current Reading

The Q2 2026 update estimates the Delegation Curve at:

```text
45.8
```

That is up **8.1 points since 2025** on the current measurement series.

The curve preserves prior public points, while archived published runs remain available in the dataset for audit. The goal is to maintain a continuous signal over time, not just a one-off score.

## What The Score Means

The score is a 0-100 composite estimate of AI influence and delegated workflow share. It is not a literal claim that a fixed percentage of all human decisions are made by AI.

Each domain combines direct workflow evidence where available with proxy signals where direct measurement is still incomplete. Examples include automated enforcement rates, adoption surveys, regulatory datasets, workflow telemetry, platform disclosures, and research studies.

Status bands:

- **Nominal:** below 40
- **Elevated:** 40 to 74.9
- **Autonomous:** 75 and above

## Domains Tracked

The current index tracks nine domains:

- **Content moderation** — automated detection, flagging, and enforcement actions on major platforms.
- **Algorithmic trading** — market execution automation and AI-assisted trade-execution adoption.
- **Software development** — AI-generated code output, workflow reliance, and agentic coding delegation.
- **Customer support** — AI-handled cases, bot deflection, production agents, and mature support deployments.
- **Credit decisioning** — AI involvement in underwriting and lending decisions.
- **Medical diagnostics** — AI influence in clinical diagnosis, imaging, pathology, and FDA-cleared medical AI capacity.
- **Legal research and review** — AI adoption in legal work, research, and document review.
- **Recruitment and screening** — AI involvement in talent acquisition, screening, assessments, and hiring workflows.
- **Education and assessment** — AI influence in student work, grading, teacher workflow, and AI-written submissions.

## Methodology At A Glance

1. **Collect public evidence.** Each domain starts with sourced indicators from reports, databases, filings, surveys, and research.
2. **Normalize indicators.** Raw values are converted to a 0-100 scale. Direct percentages stay direct; count and reach metrics use fixed caps or scaling.
3. **Weight indicators and domains.** Each domain score is a weighted average of its indicators. The composite score is a weighted average across domains.
4. **Maintain the curve.** When sources or scoring rules improve, public prior points can be restated onto the current measurement series. Older published points remain available for audit.
5. **Separate method from interpretation.** Source choices, formulas, weights, and confidence notes are kept distinct from narrative conclusions.

## Source Confidence

The project treats source quality as part of the measurement, not a footnote.

- **Direct workflow evidence** is preferred: transparency reports, operational datasets, regulatory disclosures, and first-party telemetry.
- **Strong proxies** are used where direct delegated-decision data is unavailable: adoption reports, official databases, product-market denominators, and industry surveys.
- **Context and guardrails** help interpret scope and confidence: benchmark studies, quality reports, governance findings, and reliability research.

Some domains are more mature than others. Content moderation has strong platform-level evidence. Software development now has several high-value 2026 signals. Hiring, medicine, law, and education still rely more heavily on surveys and proxy measures, so confidence should be interpreted accordingly.

## Data Access

The underlying data is published with the site:

- [`/seed.json`](https://curve.thinkwright.ai/seed.json) contains the full JSON seed dataset.
- [`/data`](https://curve.thinkwright.ai/data) provides downloadable data access from the report.

The repository also includes the data-generation pipeline used to materialize the public dataset. Analysts can inspect `seed/seed.json`, `seed/overrides.yaml`, and `internal/collect/score.go` to review values, formulas, and weights.

## Reproducing Locally

```sh
git clone git@github.com:thinkwright/delegation-curve.git
cd delegation-curve
cd frontend
npm ci
npm run dev
```

For a full data rebuild, use the Makefile targets in the repository root:

```sh
cd ..
make generate
make test
```

## License

MIT
