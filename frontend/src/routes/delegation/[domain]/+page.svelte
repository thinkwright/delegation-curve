<script lang="ts">
	import type { DelegationDomain } from '$lib/data/types';
	import { getDomainDetail, getMeta } from '$lib/db/client';
	import { base } from '$app/paths';
	import CurveChart from '$lib/components/CurveChart.svelte';
	import StatusBadge from '$lib/components/StatusBadge.svelte';
	import MetricRow from '$lib/components/MetricRow.svelte';
	import { formatDelta } from '$lib/utils/format';

	let { data } = $props();

	let domain = $state<DelegationDomain | null>(null);
	let dataYear = $state<number>(2025);
	let error = $state<string | null>(null);

	$effect(() => {
		domain = null;
		error = null;
		Promise.all([getDomainDetail(data.domainId), getMeta()])
			.then(([d, meta]) => {
				if (!d) { error = 'Domain not found'; return; }
				domain = d;
				dataYear = meta.delegation.dataYear;
			})
			.catch(e => { error = e.message; });
	});

	const delta = $derived(domain ? domain.score - domain.previousScore : 0);
	const tierLabel = $derived(domain ? `Tier ${domain.tier}` : '');
	const weightLabel = $derived(domain ? `${(domain.weight * 100).toFixed(0)}% of composite` : '');
</script>

<svelte:head>
	<title>{domain?.fullName ?? data.domainId} — AI Delegation Score — AI Delegation Curve</title>
	<meta name="description" content="{domain ? `${domain.fullName} AI delegation score: ${domain.score}/100 (${domain.status}). ${domain.description}` : `AI delegation score for ${data.domainId}. Measuring AI decision-making influence with normalized indicators.`}" />
	<link rel="canonical" href="https://curve.thinkwright.ai/delegation/{data.domainId}" />
	<meta property="og:title" content="{domain?.fullName ?? data.domainId} — AI Delegation Score" />
	<meta property="og:description" content="{domain ? `Score: ${domain.score}/100. ${domain.description}` : `AI delegation score for ${data.domainId}.`}" />
	<meta property="og:url" content="https://curve.thinkwright.ai/delegation/{data.domainId}" />
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
			{ "@type": "ListItem", "position": 2, "name": "Domains", "item": "https://curve.thinkwright.ai/delegation" },
			{ "@type": "ListItem", "position": 3, "name": domain?.fullName ?? data.domainId, "item": "https://curve.thinkwright.ai/delegation/" + data.domainId }
		]
	})}</script>`}
</svelte:head>

{#if error}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[60vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300">error_outline</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Failed to load data</p>
	<p class="text-xs text-neutral-400 mt-1 font-mono">{error}</p>
</div>
{:else if !domain}
<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[60vh]">
	<span class="material-symbols-outlined text-3xl text-neutral-300 animate-pulse">show_chart</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Loading domain data</p>
</div>
{:else}

<!-- Domain Identity -->
<div class="px-4 pt-8 pb-4">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Delegation Curve</p>
	<div class="flex items-start justify-between mb-4">
		<div>
			<h2 class="text-3xl font-black tracking-tight uppercase leading-none mb-1">{domain.name}</h2>
			<p class="font-mono text-xs text-neutral-500 mt-1">{domain.fullName}</p>
		</div>
		<div class="flex flex-col items-end gap-1.5 mt-1">
			<StatusBadge status={domain.status} />
			<span class="text-[10px] font-mono text-neutral-400 uppercase">{tierLabel}</span>
		</div>
	</div>

	<!-- Hero Score -->
	<div class="flex items-end gap-3 mb-2">
		<span class="text-[72px] font-black font-mono tabular-nums tracking-tighter leading-none
			{domain.status === 'autonomous' ? 'text-rose' : ''}">{domain.score}</span>
		<div class="flex flex-col mb-3">
			<span class="text-base font-mono font-bold tabular-nums {delta >= 0 ? 'text-sage' : 'text-rose'}">{formatDelta(delta)}</span>
			<span class="text-[10px] font-mono text-neutral-400 uppercase">vs prior period</span>
		</div>
	</div>

	<p class="text-xs text-neutral-500 leading-relaxed">{domain.description}</p>
</div>

<!-- Full Curve -->
<div class="px-4 pb-2 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">Historical Trend</p>
	<CurveChart
		data={domain.trend}
		height={180}
		color={domain.status === 'autonomous' ? 'var(--color-rose)' : 'var(--color-primary)'}
		endYear={dataYear}
	/>
</div>

<!-- KPI Row -->
<div class="grid grid-cols-3 hairline-b">
	<div class="p-3 border-r border-grid text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Score</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{domain.score}</p>
	</div>
	<div class="p-3 border-r border-grid text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Weight</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{weightLabel.split('%')[0]}%</p>
	</div>
	<div class="p-3 text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Tier</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{domain.tier}</p>
	</div>
</div>

<!-- Sub-Indicators -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-4">Sub-Indicators</p>
	{#each domain.subIndicators as indicator}
		<MetricRow
			label={indicator.name}
			value={typeof indicator.value === 'number' ? indicator.value.toLocaleString() : indicator.value}
			unit={indicator.unit}
		/>
	{/each}
</div>

<!-- Data Sources -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-4">Data Sources</p>
	<div class="space-y-3">
		{#each domain.dataSources as source}
			<div class="flex items-start justify-between py-2 border-b border-grid last:border-0">
				<div>
					<p class="text-sm font-medium">{source.name}</p>
					<p class="text-[10px] font-mono text-neutral-400 uppercase mt-0.5">{source.type}</p>
				</div>
				<span class="text-xs font-mono text-neutral-500 shrink-0">{source.cadence}</span>
			</div>
		{/each}
	</div>
</div>

<!-- Methodology Link -->
<a href="{base}/about" class="flex items-center justify-between p-4 hairline-b hover:bg-surface transition-colors group">
	<div>
		<span class="text-sm font-bold uppercase tracking-wide">Methodology</span>
		<p class="text-xs text-neutral-500 mt-0.5">How this domain score is calculated</p>
	</div>
	<span class="material-symbols-outlined text-xl text-neutral-400 group-hover:translate-x-1 transition-transform">arrow_forward</span>
</a>
{/if}
