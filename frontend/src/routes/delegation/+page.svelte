<script lang="ts">
	import { getDomains } from '$lib/db/client';
	import type { DelegationDomain } from '$lib/data/types';
	import DomainRow from '$lib/components/DomainRow.svelte';
	import FilterBar from '$lib/components/FilterBar.svelte';

	const filters = ['All', 'Nominal', 'Elevated', 'Autonomous'];
	let active = $state('All');

	let domains = $state<DelegationDomain[]>([]);
	let error = $state<string | null>(null);
	$effect(() => { getDomains().then(d => { domains = d; }).catch(e => { error = e.message; }); });

	const maxDelta = $derived(
		Math.max(...domains.map(d => d.score - d.previousScore), 1)
	);

	const sorted = $derived(
		[...domains].sort((a, b) => b.score - a.score)
	);

	const filtered = $derived(
		active === 'All'
			? sorted
			: sorted.filter((d) => d.status === active.toLowerCase())
	);
</script>

<svelte:head>
	<title>AI Delegation Scores by Domain — AI Delegation Curve</title>
	<meta name="description" content="AI delegation scores ranked across 9 decision domains: content moderation, algorithmic trading, AI code generation, customer support automation, AI credit scoring, medical AI diagnostics, legal AI tools, AI hiring systems, and AI in education. Each domain scored 0–100 based on normalized indicators from transparency reports, regulatory filings, and public benchmarks." />
	<link rel="canonical" href="https://curve.thinkwright.ai/delegation" />
	<meta property="og:title" content="AI Delegation Scores by Domain — AI Delegation Curve" />
	<meta property="og:description" content="AI delegation scores ranked across 9 decision domains. See which sectors have delegated the most decision-making power to AI systems." />
	<meta property="og:url" content="https://curve.thinkwright.ai/delegation" />
	<meta property="og:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
	<meta property="og:image:width" content="1200" />
	<meta property="og:image:height" content="630" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
	{@html `<script type="application/ld+json">${JSON.stringify({
		"@context": "https://schema.org",
		"@type": "BreadcrumbList",
		"itemListElement": [
			{ "@type": "ListItem", "position": 1, "name": "AI Delegation Curve", "item": "https://curve.thinkwright.ai" },
			{ "@type": "ListItem", "position": 2, "name": "Domains", "item": "https://curve.thinkwright.ai/delegation" }
		]
	})}</script>`}
</svelte:head>

<div class="px-4 py-8 pb-4">
	<h2 class="text-4xl font-black tracking-tight uppercase leading-none mb-1">Delegation<br/>Curve</h2>
	<p class="font-mono text-xs text-neutral-500 uppercase tracking-widest">% of consequential decisions made or influenced by AI</p>
</div>

<FilterBar {filters} {active} onchange={(f) => active = f} />

{#if error}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[40vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300">error_outline</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Failed to load data</p>
	<p class="text-xs text-neutral-400 mt-1 font-mono">{error}</p>
</div>
{:else if domains.length === 0}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[40vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300 animate-pulse">show_chart</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Loading domains</p>
</div>
{:else}
<div class="flex flex-col">
	{#each filtered as domain (domain.id)}
		<DomainRow {domain} {maxDelta} />
	{/each}
	{#if filtered.length === 0}
		<div class="p-8 text-center text-neutral-400 text-sm">No domains match this filter.</div>
	{/if}
</div>
{/if}
