<script lang="ts">
	import { getMeta } from '$lib/db/client';
	import type { CompositeData } from '$lib/data/types';
	import { base } from '$app/paths';
	import { formatDelta } from '$lib/utils/format';
	import CurveChart from '$lib/components/CurveChart.svelte';
	import MetricCard from '$lib/components/MetricCard.svelte';

	let composite = $state<CompositeData | null>(null);
	let error = $state<string | null>(null);
	$effect(() => { getMeta().then(d => { composite = d; }).catch(e => { error = e.message; }); });
</script>

<svelte:head>
	<title>AI Delegation Curve — A measure of consequential decision-making by AI</title>
	<meta name="description" content="The AI Delegation Curve: a single composite score (0–100) measuring what percentage of consequential decisions are made or influenced by AI. Tracks AI autonomy and AI adoption across content moderation, algorithmic trading, code generation, customer support, credit decisioning, medical diagnostics, legal AI, hiring automation, and education." />
	<link rel="canonical" href="https://curve.thinkwright.ai" />
	<meta property="og:title" content="AI Delegation Curve — A measure of consequential decision-making by AI" />
	<meta property="og:description" content="A composite index tracking AI decision-making influence across 9 domains. How much autonomy have we delegated to AI systems?" />
	<meta property="og:url" content="https://curve.thinkwright.ai" />
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
				"@type": "Organization",
				"name": "ThinkWright",
				"url": "https://thinkwright.ai",
				"logo": "https://curve.thinkwright.ai/og-image.png"
			},
			{
				"@type": "WebSite",
				"name": "AI Delegation Curve",
				"url": "https://curve.thinkwright.ai",
				"description": "A composite index tracking what percentage of consequential decisions are delegated to AI systems across 9 domains.",
				"publisher": { "@type": "Organization", "name": "ThinkWright" }
			},
			{
				"@type": "BreadcrumbList",
				"itemListElement": [
					{ "@type": "ListItem", "position": 1, "name": "AI Delegation Curve", "item": "https://curve.thinkwright.ai" }
				]
			}
		]
	})}</script>`}
</svelte:head>

{#if error}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[60vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300">error_outline</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Failed to load data</p>
	<p class="text-xs text-neutral-400 mt-1 font-mono">{error}</p>
</div>
{:else if !composite}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[60vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300 animate-pulse">show_chart</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Loading index data</p>
</div>
{:else}
<!-- Hero Section -->
<div class="px-4 pt-8 pb-6">
	<div class="flex justify-between items-start mb-6">
		<div>
			<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">Measuring AI Decision Influence</p>
			<h2 class="text-4xl font-black tracking-tight uppercase leading-none">Delegation<br/>Curve</h2>
		</div>
		<div class="flex items-center gap-1.5 mt-1">
			<div class="w-2 h-2 bg-sage animate-pulse"></div>
			<span class="text-[10px] font-bold uppercase tracking-wider text-neutral-400">Measuring</span>
		</div>
	</div>

	<!-- Hero Number -->
	<div class="flex items-end gap-3 mb-1">
		<span class="text-[96px] font-black font-mono tabular-nums tracking-tighter leading-none">{composite.delegation.current}</span>
		<div class="flex flex-col mb-4">
			<span class="text-lg font-mono font-bold text-sage tabular-nums">{formatDelta(composite.delegation.delta)}%</span>
			<span class="text-[10px] font-mono text-neutral-400 uppercase">vs prior</span>
		</div>
	</div>

	<p class="text-xs text-neutral-500 leading-relaxed max-w-xs">
		What percentage of consequential decisions are made or influenced by AI? A weighted composite of {composite.domainsTracked} decision domains.
	</p>
	<p class="text-[10px] font-mono text-neutral-400 uppercase mt-2">Data through {composite.delegation.dataYear}</p>
</div>

<!-- Curve Chart -->
<div class="px-4 pb-2 hairline-b">
	<CurveChart data={composite.delegation.trend} height={160} endYear={composite.delegation.dataYear} />
</div>

<!-- Metric Grid -->
<div class="grid grid-cols-2">
	<MetricCard
		label="Domains Tracked"
		value="{composite.domainsTracked}"
		subtitle="{4} Tier 1 · {5} Tier 2/3"
		icon="grid_view"
		href="{base}/delegation"
	/>
	<MetricCard
		label="Prior Year"
		value="{composite.delegation.previous}"
		subtitle="{composite.delegation.dataYear - 1} composite"
		icon="history"
		href="{base}/delegation"
	/>
	<MetricCard
		label="Highest Domain"
		value="{composite.highestDomain.score}"
		subtitle="{composite.highestDomain.name}"
		icon="trending_up"
		href="{base}/delegation/{composite.highestDomain.name.toLowerCase()}"
	/>
	<MetricCard
		label="Data Freshness"
		value="{composite.dataFreshness}"
		subtitle="Last updated {composite.delegation.lastUpdated}"
		icon="schedule"
		href="{base}/about"
	/>
</div>

<!-- Quick Link -->
<a href="{base}/delegation" class="flex items-center justify-between p-4 hover:bg-surface transition-colors group">
	<div>
		<span class="text-sm font-bold uppercase tracking-wide">View Delegation Curve Index</span>
		<p class="text-xs text-neutral-500 mt-0.5">{composite.domainsTracked} domains · Scores, trends, and data sources</p>
	</div>
	<span class="material-symbols-outlined text-xl text-neutral-400 group-hover:translate-x-1 transition-transform">arrow_forward</span>
</a>
{/if}
