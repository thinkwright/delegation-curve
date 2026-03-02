<script lang="ts">
	import { base } from '$app/paths';
	import JsonTree from '$lib/components/JsonTree.svelte';

	let seedData = $state<Record<string, unknown> | null>(null);
	let copied = $state(false);
	let rawJson = $state('');

	$effect(() => {
		fetch(`${base}/seed.json`)
			.then(r => r.json())
			.then(d => {
				seedData = d;
				rawJson = JSON.stringify(d, null, 2);
			})
			.catch(() => {});
	});

	function copyToClipboard() {
		navigator.clipboard.writeText(rawJson).then(() => {
			copied = true;
			setTimeout(() => { copied = false; }, 2000);
		});
	}
</script>

<svelte:head>
	<title>Data — Raw Seed Data — AI Delegation Curve</title>
	<meta name="description" content="Raw seed data for the AI Delegation Curve: composite scores, domain scores, sub-indicators, data sources, and trend history. JSON format, freely downloadable." />
	<link rel="canonical" href="https://curve.thinkwright.ai/data" />
	<meta property="og:title" content="Data — Raw Seed Data" />
	<meta property="og:description" content="Raw seed data underlying the AI Delegation Curve. Composite scores, domain scores, sub-indicators, and trend history." />
	<meta property="og:url" content="https://curve.thinkwright.ai/data" />
	<meta property="og:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
	<meta property="og:image:width" content="1200" />
	<meta property="og:image:height" content="630" />
	<meta property="og:type" content="website" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
	{@html `<script type="application/ld+json">${JSON.stringify({
		"@context": "https://schema.org",
		"@graph": [
			{
				"@type": "Dataset",
				"name": "AI Delegation Curve Seed Data",
				"description": "Composite scores, domain scores, sub-indicators, data sources, and trend history for the AI Delegation Curve — measuring AI decision-making influence across 9 domains.",
				"url": "https://curve.thinkwright.ai/data",
				"license": "https://creativecommons.org/licenses/by/4.0/",
				"creator": { "@type": "Organization", "name": "ThinkWright", "url": "https://thinkwright.ai" },
				"distribution": {
					"@type": "DataDownload",
					"encodingFormat": "application/json",
					"contentUrl": "https://curve.thinkwright.ai/seed.json"
				},
				"temporalCoverage": "2024/2025",
				"keywords": ["AI autonomy", "AI delegation", "AI decision-making", "AI governance", "composite index"]
			},
			{
				"@type": "BreadcrumbList",
				"itemListElement": [
					{ "@type": "ListItem", "position": 1, "name": "AI Delegation Curve", "item": "https://curve.thinkwright.ai" },
					{ "@type": "ListItem", "position": 2, "name": "Data", "item": "https://curve.thinkwright.ai/data" }
				]
			}
		]
	})}</script>`}
</svelte:head>

<div class="px-4 py-8 pb-4">
	<h2 class="text-4xl font-black tracking-tight uppercase leading-none mb-1">Data</h2>
	<p class="font-mono text-xs text-neutral-500 uppercase tracking-widest">Raw seed file</p>
</div>

<div class="px-4 py-6 hairline-b hairline-t">
	<p class="text-sm leading-relaxed mb-4">
		The complete dataset underlying the AI Delegation Curve. Composite scores, domain scores, sub-indicators, data sources, and trend history — all in a single JSON file.
	</p>
	<div class="flex gap-3 items-center">
		<a
			href="{base}/seed.json"
			download="seed.json"
			class="inline-flex items-center gap-2 px-3 py-1.5 text-xs font-bold uppercase tracking-wider bg-black text-white hover:bg-neutral-800 transition-colors"
		>
			<span class="material-symbols-outlined text-sm">download</span>
			Download JSON
		</a>
		<button
			onclick={copyToClipboard}
			class="inline-flex items-center gap-2 px-3 py-1.5 text-xs font-bold uppercase tracking-wider border border-black hover:bg-surface transition-colors"
		>
			<span class="material-symbols-outlined text-sm">{copied ? 'check' : 'content_copy'}</span>
			{copied ? 'Copied' : 'Copy'}
		</button>
	</div>
</div>

<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Schema</p>
	<div class="space-y-2 text-sm">
		<div class="flex gap-3">
			<span class="font-mono text-xs text-rose shrink-0 w-20">composite</span>
			<span class="text-neutral-500">Weighted average delegation score, delta, and trend array</span>
		</div>
		<div class="flex gap-3">
			<span class="font-mono text-xs text-rose shrink-0 w-20">domains[]</span>
			<span class="text-neutral-500">9 domain objects with score, status, weight, sub-indicators, and data sources</span>
		</div>
		<div class="flex gap-3">
			<span class="font-mono text-xs text-sage shrink-0 w-20">score</span>
			<span class="text-neutral-500">0-100 scale representing AI decision-making influence</span>
		</div>
		<div class="flex gap-3">
			<span class="font-mono text-xs text-sage shrink-0 w-20">trend</span>
			<span class="text-neutral-500">Array of historical scores (one per data year)</span>
		</div>
	</div>
</div>

<div class="py-4 hairline-b">
	<div class="px-4 flex items-center justify-between mb-3">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">seed.json</p>
		<span class="font-mono text-[10px] text-neutral-400">{rawJson ? `${(rawJson.length / 1024).toFixed(1)} KB` : ''}</span>
	</div>
	<div class="px-4 overflow-x-auto">
		{#if seedData}
			<JsonTree data={seedData} />
		{:else}
			<div class="py-8 text-center">
				<span class="text-xs font-mono text-neutral-400">Loading...</span>
			</div>
		{/if}
	</div>
</div>

<div class="px-4 py-6">
	<p class="text-sm leading-relaxed text-neutral-500">
		This data is freely available for research and analysis. See <a href="{base}/about" class="underline hover:text-black transition-colors">Methodology</a> for how scores are computed.
	</p>
</div>
