<script lang="ts">
	import type { DelegationDomain } from '$lib/data/types';
	import { getDomainDetail } from '$lib/db/client';
	import { base } from '$app/paths';
	import CurveChart from '$lib/components/CurveChart.svelte';
	import StatusBadge from '$lib/components/StatusBadge.svelte';
	import IndicatorRow from '$lib/components/IndicatorRow.svelte';
	import Seo from '$lib/components/Seo.svelte';
	import { formatDelta } from '$lib/utils/format';
	import { breadcrumbs, ogImageUrl } from '$lib/site';
	import { sourceDisclosures } from '$lib/data/sourceDisclosures';

	let { data } = $props();

	// Headline fields come from the prerendered summary; run history, indicators,
	// and sources load from Parquet in the browser.
	const summary = $derived(data.summary);
	let domain = $state<DelegationDomain | null>(null);
	let error = $state<string | null>(null);

	$effect(() => {
		const id = data.domainId;
		domain = null;
		error = null;
		getDomainDetail(id)
			.then((d) => {
				if (d) domain = d;
				else error = 'Domain detail not found';
			})
			.catch((e) => { error = e.message; });
	});

	const delta = $derived(summary.score - summary.previousScore);
	const priorPeriodLabel = $derived(data.meta.delegation.previousPeriod ?? 'prior run');
	const history = $derived(domain?.runHistory ?? []);
	const domainScores = $derived(history.length ? history.map((p) => p.score) : domain?.trend ?? []);
	const domainPeriods = $derived(history.length ? history.map((p) => p.measurementPeriod) : undefined);
	const domainDates = $derived(history.map((p) => p.periodMidpoint));
	const supplementalSources = $derived(sourceDisclosures[summary.id] ?? []);
	const path = $derived(`/delegation/${summary.id}`);
</script>

<Seo
	title={`${summary.fullName} — AI Delegation Score — AI Delegation Curve`}
	description={`${summary.fullName} AI influence score: ${summary.score}/100 (${summary.status}) in the ${data.meta.dataFreshness} run. ${summary.description}`}
	{path}
	image={ogImageUrl(data.meta.delegation.runId)}
	modifiedAt={data.meta.delegation.publishedAt}
	jsonLd={breadcrumbs([
		{ name: 'Domains', path: '/delegation' },
		{ name: summary.fullName, path }
	])}
/>

<!-- Domain Identity -->
<div class="px-4 pt-8 pb-4">
	<nav aria-label="Breadcrumb" class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">
		<a href="{base}/delegation" class="hover:text-black transition-colors">Domains</a>
		<span class="text-neutral-300 mx-1" aria-hidden="true">/</span>
		<span aria-current="page">{summary.fullName}</span>
	</nav>
	<div class="flex items-start justify-between mb-4">
		<div>
			<h2 class="text-3xl font-black tracking-tight uppercase leading-none mb-1">{summary.name}</h2>
			<p class="font-mono text-xs text-neutral-500 mt-1">{summary.fullName}</p>
		</div>
		<div class="flex flex-col items-end gap-1.5 mt-1">
			<StatusBadge status={summary.status} />
			<span class="text-[10px] font-mono text-neutral-400 uppercase">Tier {summary.tier}</span>
		</div>
	</div>

	<!-- Hero Score -->
	<div class="flex items-end gap-3 mb-2">
		<span class="text-[72px] font-black font-mono tabular-nums tracking-tighter leading-none
			{summary.status === 'autonomous' ? 'text-rose' : ''}">{summary.score}</span>
		<div class="flex flex-col mb-3">
			<span class="text-base font-mono font-bold tabular-nums {delta >= 0 ? 'text-sage' : 'text-rose'}">{formatDelta(delta)} pts</span>
			<span class="text-[10px] font-mono text-neutral-400 uppercase">since {priorPeriodLabel}</span>
		</div>
	</div>

	<p class="text-xs text-neutral-500 leading-relaxed">{summary.description}</p>
</div>

<!-- Full Curve -->
<div class="px-4 pb-2 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">Historical Trend</p>
	{#if domain}
		<CurveChart
			data={domainScores}
			labels={domainPeriods}
			dates={domainDates}
			height={180}
			color={summary.status === 'autonomous' ? 'var(--color-rose)' : 'var(--color-primary)'}
			endYear={data.meta.delegation.dataYear}
		/>
	{:else}
		<div class="h-[180px] flex items-center justify-center">
			<span class="text-[10px] font-bold uppercase tracking-widest text-neutral-400 {error ? '' : 'animate-pulse'}">{error ? 'History unavailable' : 'Loading history'}</span>
		</div>
	{/if}
</div>

<!-- KPI Row -->
<div class="grid grid-cols-3 hairline-b">
	<div class="p-3 border-r border-grid text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Score</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{summary.score}</p>
	</div>
	<div class="p-3 border-r border-grid text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Weight</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{(summary.weight * 100).toFixed(0)}%</p>
	</div>
	<div class="p-3 text-center">
		<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">Tier</p>
		<p class="text-xl font-bold font-mono tabular-nums mt-1">{summary.tier}</p>
	</div>
</div>

{#if error}
<div class="px-4 py-6 hairline-b">
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400">Failed to load indicator data</p>
	<p class="text-xs text-neutral-400 mt-1 font-mono">{error}</p>
</div>
{:else if !domain}
<div class="px-4 py-10 hairline-b flex flex-col items-center">
	<span class="material-symbols-outlined text-3xl text-neutral-300 animate-pulse">show_chart</span>
	<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Loading indicators</p>
</div>
{:else}
<!-- Sub-Indicators -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">Sub-Indicators</p>
	<p class="text-[11px] text-neutral-500 leading-relaxed mb-3">Each value links to its source, with the period it measures, its <a href="{base}/about#evidence-grades" class="underline hover:text-neutral-600">evidence grade</a>, and confidence.</p>
	{#each domain.subIndicators as indicator}
		<IndicatorRow {indicator} />
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
	{#if supplementalSources.length}
		<div class="mt-6 pt-5 border-t border-grid">
			<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Source Context</p>
			<div class="space-y-4">
				{#each supplementalSources as source}
					<div class="border-b border-grid pb-4 last:border-0 last:pb-0">
						<div class="flex items-start justify-between gap-4">
							<div>
								{#if source.url}
									<a href={source.url} target="_blank" rel="noopener" class="text-sm font-medium hover:text-neutral-500 transition-colors">{source.name}</a>
								{:else}
									<p class="text-sm font-medium">{source.name}</p>
								{/if}
								<p class="text-[10px] font-mono text-neutral-400 uppercase mt-0.5">{source.role}</p>
							</div>
							<div class="text-right shrink-0">
								<p class="text-xs font-mono text-neutral-500">{source.cadence}</p>
								<p class="text-[10px] font-mono text-neutral-400 uppercase mt-0.5">{source.type}</p>
							</div>
						</div>
						<p class="text-xs text-neutral-500 leading-relaxed mt-2">{source.note}</p>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
{/if}

<!-- Methodology Link -->
<a href="{base}/about" class="flex items-center justify-between p-4 hairline-b hover:bg-surface transition-colors group">
	<div>
		<span class="text-sm font-bold uppercase tracking-wide">Methodology</span>
		<p class="text-xs text-neutral-500 mt-0.5">How this domain score is calculated</p>
	</div>
	<span class="material-symbols-outlined text-xl text-neutral-400 group-hover:translate-x-1 transition-transform">arrow_forward</span>
</a>
