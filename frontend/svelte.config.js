import adapter from '@sveltejs/adapter-static';
import { readFileSync } from 'node:fs';

// Prerender every domain page listed in the generated data, in addition to
// the pages the crawler reaches from '*'.
const meta = JSON.parse(readFileSync(new URL('./static/data/meta.json', import.meta.url), 'utf8'));
const domainEntries = (meta.domains ?? []).map((d) => `/delegation/${d.id}`);

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			// Served by cmd/server with a 404 status for unknown paths.
			fallback: '404.html'
		}),
		prerender: {
			entries: ['*', ...domainEntries, '/og']
		}
	}
};

export default config;
