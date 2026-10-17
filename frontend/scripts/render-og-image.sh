#!/bin/sh
# Renders static/og-image.png (1200x630) from the prerendered /og route, so
# the share image always matches the current run. Requires Chrome (see
# scripts/browser.mjs). Run after `make generate`, then commit the PNG.
set -eu
cd "$(dirname "$0")/.."

PORT=${PORT:-4179}
npm run build >/dev/null
npx vite preview --host 127.0.0.1 --port "$PORT" --strictPort >/dev/null 2>&1 &
PREVIEW=$!
trap 'kill $PREVIEW 2>/dev/null || true' EXIT

ready=
for _ in $(seq 1 100); do
	if curl -sf "http://127.0.0.1:$PORT/og" >/dev/null; then ready=1; break; fi
	sleep 0.2
done
[ -n "$ready" ] || { echo "preview server did not start on port $PORT" >&2; exit 1; }
node scripts/browser.mjs screenshot "http://127.0.0.1:$PORT/og" static/og-image.png 1200 630 .og
