// Fails when the DuckDB-WASM pieces drift apart: the CDN API module version
// pinned in src/lib/db/boot.ts, the installed package, and the self-hosted
// worker scripts in static/duckdb/ must all be the same release.
import { readFileSync } from 'node:fs';

const read = (p) => readFileSync(new URL(`../${p}`, import.meta.url));
const boot = read('src/lib/db/boot.ts').toString();
const pinned = boot.match(/DUCKDB_VERSION = '([^']+)'/)?.[1];
const typed = boot.match(/duckdb-wasm@([\d.]+)\/\+esm'\)/)?.[1];
const installed = JSON.parse(read('node_modules/@duckdb/duckdb-wasm/package.json')).version;

const problems = [];
if (!pinned) problems.push('DUCKDB_VERSION not found in boot.ts');
if (pinned !== installed) problems.push(`boot.ts pins ${pinned}, package installs ${installed}`);
if (typed !== pinned) problems.push(`boot.ts type import uses ${typed}, DUCKDB_VERSION is ${pinned}`);
for (const worker of ['duckdb-browser-eh.worker.js', 'duckdb-browser-mvp.worker.js']) {
	const hosted = read(`static/duckdb/${worker}`);
	const dist = read(`node_modules/@duckdb/duckdb-wasm/dist/${worker}`);
	if (!hosted.equals(dist)) problems.push(`static/duckdb/${worker} differs from the installed ${installed} build`);
}

if (problems.length) {
	console.error(problems.map((p) => `duckdb: ${p}`).join('\n'));
	process.exit(1);
}
console.log(`duckdb-wasm ${installed}: CDN pin and self-hosted workers match`);
