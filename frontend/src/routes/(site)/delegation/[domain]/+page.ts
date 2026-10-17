import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, parent }) => {
	const { meta } = await parent();
	const summary = meta.domains?.find((d) => d.id === params.domain);
	if (!summary) error(404, `Unknown domain: ${params.domain}`);
	return { domainId: params.domain, summary };
};
