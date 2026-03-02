# AI Delegation Curve

[![License: MIT](https://img.shields.io/badge/License-MIT-black.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8.svg)](https://go.dev)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-5-FF3E00.svg)](https://svelte.dev)
[![Fly.io](https://img.shields.io/badge/Deployed-Fly.io-8B5CF6.svg)](https://curve.thinkwright.ai)

A single number tracking what percentage of consequential decisions are made by AI — measured consistently across 9 decision domains.

**Live at [curve.thinkwright.ai](https://curve.thinkwright.ai)**

## What This Is

The Keeling Curve measured atmospheric CO2 before anyone cared. This does the same for AI decision-making influence. One composite score (0-100), updated regularly, so the shape of the curve is the argument.

9 domains are tracked: content moderation, algorithmic trading, code generation, customer support, credit decisioning, medical diagnostics, legal AI, hiring, and education.

Each domain score is built from 3-4 normalized indicators sourced from transparency reports, regulatory filings, surveys, and public benchmarks.

## Architecture

```
seed/seed.json          Raw data (manual overrides + collected values)
        |
  cmd/collect           Automated data collection pipeline
  cmd/generate          Transform → Parquet export
        |
  frontend/static/data  8 Parquet files (~27 KB total)
        |
  frontend/             SvelteKit 5 SPA with DuckDB-WASM for in-browser queries
        |
  cmd/server            Go server with embed.FS (single binary, no runtime deps)
```

## Quick Start

```sh
# Install dependencies
cd frontend && npm ci && cd ..

# Run the full pipeline: collect → generate Parquet → build frontend
make pipeline

# Or just build and run the server locally
make server
./curve-server -port 8080
```

## Development

```sh
# Start frontend dev server (hot reload)
make dev

# Run tests
make test

# Collect fresh data for a single domain
make collect-contentmod

# Full collection + rebuild
make pipeline
```

## Deployment

Deployed to [Fly.io](https://fly.io) as a single distroless container.

```sh
make deploy
```

The Dockerfile runs a 3-stage build: Node (frontend) → Go (embed static assets) → distroless runtime.

## Data Pipeline

**Collection** (`cmd/collect`): Each domain has a collector that fetches from public sources. Manual overrides in `seed/overrides.yaml` fill gaps where APIs aren't available, with citations for every value.

**Scoring** (`internal/collect/score.go`): Raw values are min-max normalized against historical baselines and theoretical ceilings. Domain scores are weighted averages of their indicators. The composite is a weighted average across domains.

**Export** (`cmd/generate`): Transforms seed data into 8 Parquet tables consumed by the frontend via DuckDB-WASM.

## License

MIT
