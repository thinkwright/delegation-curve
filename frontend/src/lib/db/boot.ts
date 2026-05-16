// API module loaded from CDN (Vite treats URL imports as external, avoiding
// pre-bundle issues with duckdb-wasm). Worker JS is self-hosted — see below.
import * as duckdb from 'https://cdn.jsdelivr.net/npm/@duckdb/duckdb-wasm@1.29.0/+esm';
import { base } from '$app/paths';

let _db: any = null;
let _conn: any = null;
let _initPromise: Promise<any> | null = null;
let _initError: Error | null = null;
const _tablePromises = new Map<string, Promise<void>>();

const TABLES = new Set([
	'delegation',
	'sub_indicators',
	'data_sources',
	'analysis_runs',
	'domain_scores',
	'indicator_observations',
]);

export async function initDB() {
	if (_conn) return _conn;
	if (_initError) throw _initError;
	if (_initPromise) return _initPromise;

	_initPromise = createDB().catch((e) => {
		_initError = e instanceof Error ? e : new Error(String(e));
		_initPromise = null;
		throw _initError;
	});

	return _initPromise;
}

async function createDB() {
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
	try {
		await db.instantiate(bundle.mainModule, bundle.pthreadWorker);
	} finally {
		URL.revokeObjectURL(workerUrl);
	}

	const conn = await db.connect();
	_db = db;
	_conn = conn;
	return conn;
}

export async function query<T = Record<string, unknown>>(
	sql: string,
	tables: string[]
): Promise<T[]> {
	const conn = await initDB();
	await ensureTables(conn, tables);
	const result = await conn.query(sql);
	return arrowToObjects<T>(result);
}

async function ensureTables(conn: any, tables: string[]): Promise<void> {
	await Promise.all(tables.map((name) => ensureTable(conn, name)));
}

async function ensureTable(conn: any, name: string): Promise<void> {
	if (!TABLES.has(name)) throw new Error(`Unknown table: ${name}`);

	let promise = _tablePromises.get(name);
	if (!promise) {
		promise = registerTable(conn, name).catch((e) => {
			_tablePromises.delete(name);
			throw e;
		});
		_tablePromises.set(name, promise);
	}
	return promise;
}

async function registerTable(conn: any, name: string): Promise<void> {
	const basePath = (base || '');
	const resp = await fetch(`${basePath}/data/${name}.parquet`);
	if (!resp.ok) throw new Error(`Failed to load ${name}.parquet: ${resp.status}`);
	await _db.registerFileBuffer(`${name}.parquet`, new Uint8Array(await resp.arrayBuffer()));
	await conn.query(`CREATE VIEW "${name}" AS SELECT * FROM '${name}.parquet'`);
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
