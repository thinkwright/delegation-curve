// Headless Chrome checks for the built site, driven over the DevTools
// protocol with Node's built-in fetch and WebSocket (no extra dependencies).
//
//   node scripts/browser.mjs smoke <base-url>
//       Loads each page, waits for its client-rendered content, and fails on
//       uncaught exceptions or console errors.
//   node scripts/browser.mjs probe <url> <text> [timeout-ms]
//       Waits for text on one page and prints what rendered (debugging).
//   node scripts/browser.mjs screenshot <url> <out.png> [width] [height] [selector]
//       Captures a viewport screenshot once fonts and the optional selector
//       are ready (used for static/og-image.png).
//
// Set CHROME to the browser binary if it is not in a standard location.

import { spawn } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

const CHROME_CANDIDATES = [
	process.env.CHROME,
	'/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
	'/usr/bin/google-chrome',
	'/usr/bin/google-chrome-stable',
	'/usr/bin/chromium',
	'/usr/bin/chromium-browser'
].filter(Boolean);

// Third-party noise that says nothing about the site itself.
const IGNORED_CONSOLE = [/googletagmanager|google-analytics/i, /fonts\.(googleapis|gstatic)/i];

const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

async function launch() {
	const bin = CHROME_CANDIDATES.find((p) => existsSync(p));
	if (!bin) throw new Error('Chrome not found; set CHROME');
	const profile = mkdtempSync(join(tmpdir(), 'curve-chrome-'));
	const proc = spawn(bin, [
		'--headless=new',
		'--disable-gpu',
		'--no-first-run',
		'--no-default-browser-check',
		'--hide-scrollbars',
		`--user-data-dir=${profile}`,
		'--remote-debugging-port=0',
		'about:blank'
	], { stdio: 'ignore' });

	const portFile = join(profile, 'DevToolsActivePort');
	for (let i = 0; i < 100 && !existsSync(portFile); i++) await sleep(100);
	if (!existsSync(portFile)) throw new Error('Chrome did not start');
	const [port, path] = readFileSync(portFile, 'utf8').trim().split('\n');

	const ws = new WebSocket(`ws://127.0.0.1:${port}${path}`);
	await new Promise((resolve, reject) => {
		ws.addEventListener('open', resolve, { once: true });
		ws.addEventListener('error', reject, { once: true });
	});

	let nextId = 1;
	const pending = new Map();
	const listeners = new Set();
	ws.addEventListener('message', (ev) => {
		const msg = JSON.parse(ev.data);
		if (msg.id && pending.has(msg.id)) {
			const { resolve, reject } = pending.get(msg.id);
			pending.delete(msg.id);
			msg.error ? reject(new Error(msg.error.message)) : resolve(msg.result);
		} else if (msg.method) {
			for (const fn of listeners) fn(msg);
		}
	});
	const send = (method, params = {}, sessionId) =>
		new Promise((resolve, reject) => {
			const id = nextId++;
			pending.set(id, { resolve, reject });
			ws.send(JSON.stringify({ id, method, params, sessionId }));
		});

	return {
		send,
		on: (fn) => listeners.add(fn),
		async close() {
			ws.close();
			proc.kill();
			await sleep(200);
			rmSync(profile, { recursive: true, force: true });
		}
	};
}

async function openPage(browser, { width = 390, height = 844 } = {}) {
	const { targetId } = await browser.send('Target.createTarget', { url: 'about:blank' });
	const { sessionId } = await browser.send('Target.attachToTarget', { targetId, flatten: true });
	const send = (method, params) => browser.send(method, params, sessionId);
	const problems = [];
	browser.on((msg) => {
		if (msg.sessionId !== sessionId) return;
		if (msg.method === 'Runtime.exceptionThrown') {
			const d = msg.params.exceptionDetails;
			problems.push(`exception: ${d.exception?.description ?? d.text}`);
		}
		if (msg.method === 'Runtime.consoleAPICalled' && msg.params.type === 'error') {
			const text = msg.params.args.map((a) => a.value ?? a.description ?? '').join(' ');
			if (!IGNORED_CONSOLE.some((re) => re.test(text))) problems.push(`console.error: ${text}`);
		}
		if (msg.method === 'Log.entryAdded' && msg.params.entry.level === 'error') {
			const { text, url = '' } = msg.params.entry;
			if (!IGNORED_CONSOLE.some((re) => re.test(text) || re.test(url))) problems.push(`log: ${text} ${url}`);
		}
	});
	await send('Runtime.enable');
	await send('Log.enable');
	await send('Page.enable');
	await send('Emulation.setDeviceMetricsOverride', { width, height, deviceScaleFactor: 1, mobile: false });
	return {
		problems,
		async goto(url) {
			await send('Page.navigate', { url });
		},
		async eval(expression) {
			const r = await send('Runtime.evaluate', { expression, returnByValue: true, awaitPromise: true });
			return r.result.value;
		},
		async waitFor(expression, timeoutMs) {
			const deadline = Date.now() + timeoutMs;
			while (Date.now() < deadline) {
				try {
					if (await this.eval(expression)) return true;
				} catch {
					// Navigation in progress.
				}
				await sleep(250);
			}
			return false;
		},
		async screenshot() {
			const { data } = await send('Page.captureScreenshot', { format: 'png' });
			return Buffer.from(data, 'base64');
		}
	};
}

// innerText applies CSS text-transform, so compare case-insensitively.
const textIncludes = (...needles) =>
	`(() => { const t = (document.body ? document.body.innerText : '').toLowerCase(); return ${JSON.stringify(needles.map((n) => n.toLowerCase()))}.every((n) => t.includes(n)); })()`;

async function smoke(base) {
	const meta = await (await fetch(`${base}/data/meta.json`)).json();
	const current = String(meta.delegation.current);
	const checks = [
		{ path: '/', expect: [current, 'DELEGATION'] },
		{ path: '/delegation', expect: [meta.domains[0].name] },
		{ path: '/delegation/code-gen', expect: ['Technical Work AI Value Share', 'Grade'], slow: true },
		{ path: '/about', expect: ['Evidence Grades'] },
		{ path: '/data', expect: ['analysis_runs'], slow: true },
		{ path: '/delegation/not-a-domain', expect: ['Page not found'], expectStatus: 404 }
	];

	const browser = await launch();
	let failed = 0;
	try {
		for (const check of checks) {
			const page = await openPage(browser);
			await page.goto(base + check.path);
			const ok = await page.waitFor(textIncludes(...check.expect), check.slow ? 45000 : 15000);
			// Let late errors (hydration, worker startup) surface.
			await sleep(500);
			// The document's own 404 is expected on the not-found check.
			if (check.expectStatus === 404) {
				const self = base + check.path;
				page.problems.splice(0, page.problems.length, ...page.problems.filter((p) => !(p.includes('404') && p.includes(self))));
			}
			const status = ok && page.problems.length === 0 ? 'ok  ' : 'FAIL';
			if (status === 'FAIL') failed++;
			console.log(`${status} ${check.path}${ok ? '' : ` (missing: ${check.expect.join(', ')})`}`);
			for (const p of page.problems) console.log(`       ${p}`);
		}
	} finally {
		await browser.close();
	}
	if (failed) {
		console.error(`${failed} page check(s) failed`);
		process.exit(1);
	}
}

async function screenshot(url, out, width = 1200, height = 630, selector) {
	const browser = await launch();
	try {
		const page = await openPage(browser, { width: Number(width), height: Number(height) });
		await page.goto(url);
		const ready = selector
			? `document.readyState === "complete" && document.fonts.status === "loaded" && !!document.querySelector(${JSON.stringify(selector)})`
			: 'document.readyState === "complete" && document.fonts.status === "loaded"';
		if (!(await page.waitFor(ready, 20000))) {
			throw new Error(`page not ready${selector ? ` (no ${selector})` : ''}; refusing to write ${out}`);
		}
		await sleep(500);
		writeFileSync(out, await page.screenshot());
		for (const p of page.problems) console.log(`warning: ${p}`);
		console.log(`wrote ${out}`);
	} finally {
		await browser.close();
	}
}

async function probe(url, needle, timeoutMs = 45000) {
	const browser = await launch();
	try {
		const page = await openPage(browser);
		const started = Date.now();
		await page.goto(url);
		const ok = await page.waitFor(textIncludes(needle), Number(timeoutMs));
		console.log(`${ok ? 'found' : 'NOT found'} ${JSON.stringify(needle)} after ${Date.now() - started}ms`);
		for (const p of page.problems) console.log(`  ${p}`);
		const text = await page.eval('document.body.innerText');
		console.log(text.slice(0, 1500));
	} finally {
		await browser.close();
	}
}

const [cmd, ...args] = process.argv.slice(2);
if (cmd === 'smoke' && args[0]) await smoke(args[0].replace(/\/$/, ''));
else if (cmd === 'probe' && args[1]) await probe(...args);
else if (cmd === 'screenshot' && args[1]) await screenshot(...args);
else {
	console.error('usage: browser.mjs smoke <base-url> | probe <url> <text> [timeout-ms] | screenshot <url> <out.png> [width] [height] [selector]');
	process.exit(2);
}
