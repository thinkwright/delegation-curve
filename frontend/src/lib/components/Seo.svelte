<script lang="ts">
	import { SITE_NAME, siteUrl } from '$lib/site';

	let {
		title,
		description,
		path,
		image,
		type = 'website',
		publishedAt,
		modifiedAt,
		jsonLd,
		noindex = false
	}: {
		title: string;
		description: string;
		/** Site path of the page, e.g. '/' or '/delegation/code-gen'. */
		path: string;
		image?: string;
		type?: string;
		publishedAt?: string;
		modifiedAt?: string;
		/** One schema.org node or a list of nodes; wrapped in an @graph. */
		jsonLd?: Record<string, unknown> | Record<string, unknown>[];
		noindex?: boolean;
	} = $props();

	const url = $derived(siteUrl(path));
	// Escape "<" so data can never close the script element early.
	const ld = $derived(
		jsonLd
			? JSON.stringify({
					'@context': 'https://schema.org',
					'@graph': Array.isArray(jsonLd) ? jsonLd : [jsonLd]
				}).replace(/</g, '\\u003c')
			: ''
	);
</script>

<svelte:head>
	<title>{title}</title>
	<meta name="description" content={description} />
	{#if noindex}
		<meta name="robots" content="noindex" />
	{:else}
		<link rel="canonical" href={url} />
	{/if}
	<meta property="og:site_name" content={SITE_NAME} />
	<meta property="og:type" content={type} />
	<meta property="og:title" content={title} />
	<meta property="og:description" content={description} />
	<meta property="og:url" content={url} />
	{#if image}
		<meta property="og:image" content={image} />
		<meta property="og:image:width" content="1200" />
		<meta property="og:image:height" content="630" />
		<meta name="twitter:image" content={image} />
	{/if}
	{#if publishedAt}<meta property="article:published_time" content={publishedAt} />{/if}
	{#if modifiedAt}
		<meta property="article:modified_time" content={modifiedAt} />
		<meta name="date" content={modifiedAt} />
	{/if}
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={title} />
	<meta name="twitter:description" content={description} />
	{#if ld}
		{@html `<script type="application/ld+json">${ld}</script>`}
	{/if}
</svelte:head>
