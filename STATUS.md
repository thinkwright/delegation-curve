# Project Status

Updated: 2026-05-16T23:00:07Z

Branch: `research/2026-source-lock`
Remote: `origin/research/2026-source-lock`
Latest implementation commit before this status update: `4d96a18 Optimize DuckDB data loading and static delivery`

## Current State

- The 2026 Q2 Delegation Curve run is implemented and pushed.
- Public composite score: `45.8`.
- Public prior point: `37.7` for 2025 under the current measurement series.
- Public movement: `+8.1 pts since 2025`.
- Public first-read copy intentionally avoids `v2`, `recalculated`, and `baseline` language.
- Historical public curve points are shown with markers; current run remains emphasized.
- Archived legacy runs remain available in the data for audit.
- DuckDB-WASM startup and static delivery have been optimized:
  - no global DuckDB initialization from the layout,
  - coalesced DuckDB initialization,
  - lazy parquet table registration,
  - memoized data accessors,
  - gzip precompression for static assets,
  - immutable cache headers for hashed app chunks.

## Local Server

The Linux workspace server was last running at:

```sh
http://10.0.0.79:8080/
```

Process at handoff:

```sh
/home/bran/code/curve/curve-server -port 8080
```

This binary is Linux-specific. On macOS, use the frontend dev server for UI screenshot work, or rebuild the Go server locally.

## Validation Completed

```sh
npm run check
PATH=/home/bran/code/curve/.tools/go/bin:$PATH go test ./...
git diff --check
PATH=/home/bran/code/curve/.tools/go/bin:$PATH make server
```

Runtime header checks confirmed:

- DuckDB worker JS serves with `Content-Encoding: gzip` when requested.
- Hashed Svelte assets serve with `Cache-Control: public, max-age=31536000, immutable`.
- Data files serve with `Cache-Control: public, max-age=300, stale-while-revalidate=86400`.
- HTML serves with `Cache-Control: no-cache`.

## Known Follow-Ups

- UI polish needs screenshot-driven review on macOS, especially mobile widths.
- Confirm the main curve and domain charts remain legible with prior datapoint markers.
- Check copy and layout on `/`, `/delegation`, `/delegation/code-gen`, `/about`, and `/data`.
- The npm audit warning remains: 7 vulnerabilities reported by `npm ci` / build output. This appears pre-existing and was not addressed in this research/UI pass.

See `research/2026-05-16-macos-ui-handoff.md` for the detailed next-session handoff.
