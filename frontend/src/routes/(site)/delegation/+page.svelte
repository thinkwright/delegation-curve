<script lang="ts">
	import DomainRow from '$lib/components/DomainRow.svelte';
	import FilterBar from '$lib/components/FilterBar.svelte';
	import Seo from '$lib/components/Seo.svelte';
	import { breadcrumbs, ogImageUrl } from '$lib/site';

	let { data } = $props();

	const filters = ['All', 'Nominal', 'Elevated', 'Autonomous'];
	let active = $state('All');

	const domains = $derived(data.meta.domains ?? []);
	const maxDelta = $derived(Math.max(...domains.map((d) => d.score - d.previousScore), 1));
	const sorted = $derived([...domains].sort((a, b) => b.score - a.score));
	const filtered = $derived(
		active === 'All' ? sorted : sorted.filter((d) => d.status === active.toLowerCase())
	);
	const priorLabel = $derived(`${data.meta.delegation.previousPeriod ?? 'prior run'} score`);
	const description = $derived(
		`AI influence scores for ${domains.length} consequential domains in the ${data.meta.dataFreshness} run: ${sorted
			.map((d) => `${d.fullName.toLowerCase()} ${d.score}`)
			.join(', ')}.`
	);
</script>

<Seo
	title="AI Delegation Scores by Domain — AI Delegation Curve"
	{description}
	path="/delegation"
	image={ogImageUrl(data.meta.delegation.runId)}
	modifiedAt={data.meta.delegation.publishedAt}
	jsonLd={breadcrumbs([{ name: 'Domains', path: '/delegation' }])}
/>


<div class="px-4 py-8 pb-4">
	<h2 class="text-4xl font-black tracking-tight uppercase leading-none mb-1">Delegation<br/>Curve</h2>
	<p class="font-mono text-xs text-neutral-500 uppercase tracking-widest">AI influence across consequential domains</p>
</div>

<FilterBar {filters} {active} onchange={(f) => active = f} />

<div class="flex flex-col">
	{#each filtered as domain (domain.id)}
		<DomainRow {domain} {maxDelta} {priorLabel} />
	{/each}
	{#if filtered.length === 0}
		<div class="p-8 text-center text-neutral-400 text-sm">No domains match this filter.</div>
	{/if}
</div>
