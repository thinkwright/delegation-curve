// DuckDB-WASM loaded via CDN at runtime.
declare module 'https://cdn.jsdelivr.net/npm/@duckdb/duckdb-wasm@1.29.0/+esm' {
	export function getJsDelivrBundles(): any;
	export function selectBundle(bundles: any): Promise<any>;

	export class ConsoleLogger {}

	export class AsyncDuckDB {
		constructor(logger: any, worker: Worker);
		instantiate(mainModule: string, pthreadWorker?: string | null): Promise<void>;
		registerFileBuffer(name: string, buffer: Uint8Array): Promise<void>;
		registerFileURL(name: string, url: string, protocol: number, directIO: boolean): Promise<void>;
		connect(): Promise<any>;
	}

	export const DuckDBDataProtocol: { HTTP: number };
}
