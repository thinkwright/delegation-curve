# macOS UI Handoff

Date: 2026-05-16
Branch: `research/2026-source-lock`
Remote: `origin/research/2026-source-lock`
Latest implementation commit before this handoff: `4d96a18 Optimize DuckDB data loading and static delivery`

## Goal For Next Session

Pick up this repo on a local macOS machine where Codex can receive screenshots. The next focus is UI review and polish. The research/data/methodology work is in a good working state; use screenshots to inspect the actual rendered experience and resolve visual issues carefully.

## Setup On macOS

```sh
git clone git@github.com:thinkwright/delegation-curve.git
cd delegation-curve
git checkout research/2026-source-lock
cd frontend
npm ci
npm run dev -- --host 127.0.0.1
```

Then open:

```text
http://127.0.0.1:5173/
```

For production-parity server testing on macOS, install/use Go and run:

```sh
cd ..
make server
./curve-server -port 8080
```

The repo already contains the generated static data under `frontend/static/data`, so frontend UI review does not require running the collection pipeline.

## What Changed Today

The 2026 Q2 update was implemented and pushed on `research/2026-source-lock`.

Key research/data state:

- Current public composite: `45.8`.
- Prior public point: `37.7` for 2025 on the current measurement series.
- Public comparison: `+8.1 pts since 2025`.
- Archived legacy points remain in the dataset for audit, but the public curve uses the current-method series.
- The visual curve now keeps prior analysis run markers visible, inspired by the Keeling Curve storytelling model.

Code-gen work:

- The code generation domain was re-evaluated after a source sweep for 2026 state-of-market evidence.
- Current code-gen score is `48.1`.
- The score now uses five inputs: Sonar/GitLab output share, METR technical-work value share, DORA/JetBrains workflow reliance, Anthropic/arXiv agentic task delegation, and low-weight VS Code Marketplace ecosystem reach.
- Stack Overflow was retired from scoring because it is useful context but too biased for the core metric.
- High-profile anecdotal signals such as Claude Code, Amazon, and Meta remain useful narrative context, but were not used as hard scoring inputs without rigorous public measurement.

Method/copy work:

- The public site no longer uses `v2`, `recalculated`, or `baseline` language on first-read surfaces.
- The public comparison language is now plain time-series language such as `since 2025`, `2025 Score`, and `Prior point on the curve`.
- Methodology copy was softened from literal percentage-of-decisions claims to a composite estimate of AI influence and delegated workflow share.
- Domain descriptions were updated to match the actual source mix and formulas.

Performance work:

- DuckDB-WASM no longer starts globally from the layout.
- DuckDB initialization is promise-coalesced, preventing duplicate worker/database startup.
- Parquet table registration is lazy and coalesced.
- `getMeta`, `getDomains`, and `getDomainDetail` are memoized.
- The server now serves precompressed `.gz` files for compressible static assets.
- Hashed app chunks are served with long immutable caching.

## Files Worth Reading First

- `STATUS.md`
- `research/2026-q2-run-summary.md`
- `research/2026-current-method-series.md`
- `research/2026-source-lock-plan.md`
- `research/2026-evidence/code-gen.md`
- `frontend/src/routes/+page.svelte`
- `frontend/src/routes/delegation/+page.svelte`
- `frontend/src/routes/delegation/[domain]/+page.svelte`
- `frontend/src/lib/components/CurveChart.svelte`
- `frontend/src/lib/components/DomainRow.svelte`
- `frontend/src/lib/db/boot.ts`
- `cmd/server/main.go`

## UI Review Checklist

Capture desktop and mobile screenshots for:

- `/`
- `/delegation`
- `/delegation/code-gen`
- `/delegation/content-mod`
- `/about`
- `/data`

Use at least:

- Mobile: 390px or 430px wide.
- Tablet-ish: 768px wide.
- Desktop: 1280px or 1440px wide.

Things to inspect:

- Hero score layout on the home page, especially `45.8` and `+8.1 pts since 2025`.
- Whether `since 2025` is understandable and visually attached to the score.
- Whether prior datapoint markers are visible but not over-emphasized.
- Curve chart label density and marker collision on mobile.
- Domain row density, delta bars, and score alignment on `/delegation`.
- Domain detail pages with long descriptions and source names.
- Bottom navigation spacing on mobile.
- Any text clipping or horizontal overflow.
- Whether first-read pages avoid confusing methodology terms such as `v2`, `recalculated`, and `baseline`.

## Validation From Linux Session

Completed before handoff:

```sh
npm run check
PATH=/home/bran/code/curve/.tools/go/bin:$PATH go test ./...
git diff --check
PATH=/home/bran/code/curve/.tools/go/bin:$PATH make server
```

The Linux environment did not have Chromium/Playwright available for screenshot validation, which is the main reason the next pass should happen on macOS with screenshots attached.

## Known Caveats

- `npm ci` / build output still reports 7 npm audit vulnerabilities. This was pre-existing and not addressed today.
- Some raw JSON/history fields still contain technical labels such as `delegation-curve-v2` and `2025 recalculated baseline`. That is intentional for audit/history data; avoid surfacing those terms on first-read UI.
- The Linux server URL was `http://10.0.0.79:8080/`, but the macOS workflow should prefer a local dev server unless you are explicitly testing the Linux host.
