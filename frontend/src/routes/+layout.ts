import { base } from '$app/paths';
import type { CompositeData } from '$lib/data/types';
import type { LayoutLoad } from './$types';

// Every page is prerendered at build time from the generated data, so the
// headline numbers, titles, and descriptions are in the HTML that crawlers
// and link previews read. Data-heavy detail still loads client-side.
export const prerender = true;
export const trailingSlash = 'never';

export const load: LayoutLoad = async ({ fetch }) => {
	const res = await fetch(`${base}/data/meta.json`);
	if (!res.ok) throw new Error(`Failed to load meta.json: ${res.status}`);
	const meta: CompositeData = await res.json();
	return { meta };
};
