export const SITE_URL = 'https://curve.thinkwright.ai';
export const SITE_NAME = 'AI Delegation Curve';
export const PUBLISHER = { '@type': 'Organization', name: 'ThinkWright', url: 'https://thinkwright.ai' };

/** Absolute URL for a site path ('/' or '/about'). */
export function siteUrl(path: string): string {
	return path === '/' ? SITE_URL : `${SITE_URL}${path}`;
}

/** Share image URL, cache-busted per analysis run so previews refresh with new data. */
export function ogImageUrl(runId?: string): string {
	return `${SITE_URL}/og-image.png${runId ? `?v=${encodeURIComponent(runId)}` : ''}`;
}

/** schema.org BreadcrumbList for a page trail, starting at the home page. */
export function breadcrumbs(trail: { name: string; path: string }[]) {
	const items = [{ name: SITE_NAME, path: '/' }, ...trail];
	return {
		'@type': 'BreadcrumbList',
		itemListElement: items.map((item, i) => ({
			'@type': 'ListItem',
			position: i + 1,
			name: item.name,
			item: siteUrl(item.path)
		}))
	};
}
