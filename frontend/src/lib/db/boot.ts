// API module loaded from CDN (Vite treats URL imports as external, avoiding
// pre-bundle issues with duckdb-wasm). Worker JS is self-hosted — see below.
import * as duckdb from 'https://cdn.jsdelivr.net/npm/@duckdb/duckdb-wasm@1.29.0/+esm';
import { base } from '$app/paths';

let _db: any = null;
let _conn: any = null;
let _initError: Error | null = null;

const TABLES = [
	'delegation',
	'sub_indicators',
	'data_sources',
];

export function getInitError(): Error | null { return _initError; }

export async function initDB() {
	if (_conn) return _conn;
	if (_initError) throw _initError;

	try {

	const basePath = (base || '');

	// Use the package's built-in CDN bundle config, then override worker URLs
	// to self-hosted copies (workers are executable JS — keep them local).
	// Full absolute URLs required because importScripts inside a blob: worker
	// can't resolve root-relative paths (blob origins are opaque).
	const bundles = duckdb.getJsDelivrBundles();
	const origin = globalThis.location.origin + basePath;
	bundles.mvp.mainWorker = `${origin}/duckdb/duckdb-browser-mvp.worker.js`;
	bundles.eh.mainWorker = `${origin}/duckdb/duckdb-browser-eh.worker.js`;
	const bundle = await duckdb.selectBundle(bundles);

	const workerUrl = URL.createObjectURL(
		new Blob([`importScripts("${bundle.mainWorker}");`], { type: 'text/javascript' })
	);
	const worker = new Worker(workerUrl);
	const logger = new duckdb.ConsoleLogger();
	const db = new duckdb.AsyncDuckDB(logger, worker);
	await db.instantiate(bundle.mainModule, bundle.pthreadWorker);
	URL.revokeObjectURL(workerUrl);

	const conn = await db.connect();

	// Fetch all parquet files in parallel as buffers — faster than HTTP range
	// requests for small files (~3-6 KB each).
	const dataPath = basePath + '/data/';
	const buffers = await Promise.all(
		TABLES.map(async (t) => {
			const resp = await fetch(dataPath + t + '.parquet');
			return { name: t, buf: new Uint8Array(await resp.arrayBuffer()) };
		})
	);
	for (const { name, buf } of buffers) {
		await db.registerFileBuffer(name + '.parquet', buf);
		await conn.query(`CREATE VIEW "${name}" AS SELECT * FROM '${name}.parquet'`);
	}

	_db = db;
	_conn = conn;
	return conn;

	} catch (e) {
		_initError = e instanceof Error ? e : new Error(String(e));
		throw _initError;
	}
}

export async function query<T = Record<string, unknown>>(sql: string): Promise<T[]> {
	const conn = await initDB();
	const result = await conn.query(sql);
	return arrowToObjects<T>(result);
}

function arrowToObjects<T>(table: any): T[] {
	const rows: T[] = [];
	const fields = table.schema.fields.map((f: any) => f.name);
	for (let i = 0; i < table.numRows; i++) {
		const row: Record<string, unknown> = {};
		for (const col of fields) {
			let val = table.getChild(col).get(i);
			if (typeof val === 'bigint') val = Number(val);
			row[col] = val;
		}
		rows.push(row as T);
	}
	return rows;
}
