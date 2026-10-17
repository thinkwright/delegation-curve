<script lang="ts">
	import { base } from '$app/paths';
	import { formatDelta, formatIsoDate } from '$lib/utils/format';
	import { PUBLISHER, SITE_NAME, SITE_URL, breadcrumbs, ogImageUrl } from '$lib/site';
	import CurveChart from '$lib/components/CurveChart.svelte';
	import MetricCard from '$lib/components/MetricCard.svelte';
	import Seo from '$lib/components/Seo.svelte';

	let { data } = $props();

	const composite = $derived(data.meta);
	const history = $derived(composite.delegation.runHistory ?? []);
	const delegationScores = $derived(history.length ? history.map((p) => p.score) : composite.delegation.trend);
	const delegationPeriods = $derived(history.length ? history.map((p) => p.measurementPeriod) : undefined);
	const delegationDates = $derived(history.map((p) => p.periodMidpoint));
	const priorPeriodLabel = $derived(composite.delegation.previousPeriod ?? 'prior run');
	const publishedAt = $derived(composite.delegation.publishedAt ?? composite.delegation.lastUpdated);
	const firstPublishedAt = $derived(history.map((p) => p.publishedAt).filter(Boolean).sort()[0] ?? publishedAt);
	const domains = $derived(composite.domains ?? []);
	const tierSummary = $derived.by(() => {
		const tier1 = domains.filter((d) => d.tier === 1).length;
		return `${tier1} Tier 1 · ${domains.length - tier1} Tier 2/3`;
	});
	const highestDomainId = $derived(
		domains.find((d) => d.name === composite.highestDomain.name)?.id ?? composite.highestDomain.name.toLowerCase()
	);

	const title = 'How much does AI influence work and society?';
	const description = $derived(
		`The AI Delegation Curve tracks ${composite.dataFreshness} AI influence across ${composite.domainsTracked} high-impact domains, from code and support to medicine, law, hiring, and education.`
	);
</script>

<Seo
	{title}
	{description}
	path="/"
	image={ogImageUrl(composite.delegation.runId)}
	publishedAt={firstPublishedAt}
	modifiedAt={publishedAt}
	jsonLd={[
		{ ...PUBLISHER, logo: `${SITE_URL}/og-image.png` },
		{
			'@type': 'WebSite',
			name: SITE_NAME,
			url: SITE_URL,
			description,
			datePublished: firstPublishedAt,
			dateModified: publishedAt,
			publisher: PUBLISHER
		},
		breadcrumbs([])
	]}
/>


<!-- Hero Section -->
<div class="px-4 pt-8 pb-6">
	<div class="flex justify-between items-start mb-6">
		<div>
			<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">Measuring AI Influence</p>
			<h2 class="text-4xl font-black tracking-tight uppercase leading-none">Delegation<br/>Curve</h2>
		</div>
		<div class="flex items-center gap-1.5 mt-1">
			<div class="w-2 h-2 bg-sage animate-pulse"></div>
			<span class="text-[10px] font-bold uppercase tracking-wider text-neutral-400">{composite.dataFreshness}</span>
		</div>
	</div>

	<!-- Hero Number -->
	<div class="flex items-end gap-3 mb-1">
		<span class="text-[96px] font-black font-mono tabular-nums tracking-tighter leading-none">{composite.delegation.current}</span>
		<div class="flex flex-col mb-4">
			<span class="text-lg font-mono font-bold text-sage tabular-nums">{formatDelta(composite.delegation.delta)} pts</span>
			<span class="text-[10px] font-mono text-neutral-400 uppercase">since {priorPeriodLabel}</span>
		</div>
	</div>

	<p class="text-xs text-neutral-500 leading-relaxed max-w-xs">
		A weighted index tracking AI influence across {composite.domainsTracked} domains where automated systems increasingly shape work, access, judgment, and public life.
	</p>
	<p class="text-[10px] font-mono text-neutral-400 uppercase mt-2">Current run: {composite.dataFreshness}</p>
</div>

<!-- Curve Chart -->
<div class="px-4 pb-2 hairline-b">
	<CurveChart
		data={delegationScores}
		labels={delegationPeriods}
		dates={delegationDates}
		height={160}
		endYear={composite.delegation.dataYear}
	/>
</div>

<!-- Metric Grid -->
<div class="grid grid-cols-2">
	<MetricCard
		label="Domains Tracked"
		value={composite.domainsTracked}
		subtitle={tierSummary}
		icon="grid_view"
		href={`${base}/delegation`}
	/>
	<MetricCard
		label={`${priorPeriodLabel} Score`}
		value={composite.delegation.previous}
		subtitle="Prior point on the curve"
		icon="history"
		href={`${base}/delegation`}
	/>
	<MetricCard
		label="Highest Domain"
		value={composite.highestDomain.score}
		subtitle={composite.highestDomain.name}
		icon="trending_up"
		href={`${base}/delegation/${highestDomainId}`}
	/>
	<MetricCard
		label="Data Freshness"
		value={composite.dataFreshness}
		subtitle={`Last updated ${formatIsoDate(publishedAt)}`}
		icon="schedule"
		href={`${base}/about`}
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
