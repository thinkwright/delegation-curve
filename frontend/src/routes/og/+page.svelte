<script lang="ts">
	import CurveChart from '$lib/components/CurveChart.svelte';
	import { formatDelta, formatIsoDate } from '$lib/utils/format';

	// Source for static/og-image.png (1200x630), captured by
	// scripts/render-og-image.sh after each run. Not linked or indexed.
	let { data } = $props();

	const m = $derived(data.meta);
	const history = $derived(m.delegation.runHistory ?? []);
	const domains = $derived(m.domains ?? []);
	const top = $derived([...domains].sort((a, b) => b.score - a.score).slice(0, 3));
	const publishedAt = $derived(m.delegation.publishedAt ?? m.delegation.lastUpdated);
</script>

<svelte:head>
	<title>Share image — AI Delegation Curve</title>
	<meta name="robots" content="noindex" />
</svelte:head>

<div class="dark og">
	<div class="og-left">
		<div>
			<p class="text-[14px] font-bold uppercase tracking-widest text-neutral-400">Measuring AI Influence</p>
			<h1 class="text-[56px] font-black tracking-tight uppercase leading-[0.95] mt-3">Delegation<br />Curve</h1>
		</div>
		<div>
			<div class="flex items-end gap-5">
				<span class="text-[168px] font-black font-mono tabular-nums tracking-tighter leading-[0.8]">{m.delegation.current.toFixed(1)}</span>
				<div class="flex flex-col mb-2">
					<span class="text-[30px] font-mono font-bold text-sage tabular-nums whitespace-nowrap">{formatDelta(m.delegation.delta)} pts</span>
					<span class="text-[16px] font-mono text-neutral-400 uppercase mt-1 whitespace-nowrap">since {m.delegation.previousPeriod ?? 'prior run'}</span>
				</div>
			</div>
			<p class="text-[19px] text-neutral-400 leading-snug mt-6 max-w-[520px]">
				A weighted index estimating AI influence across {m.domainsTracked} consequential domains, from platform enforcement to code, support, credit, medicine, law, hiring, and education.
			</p>
		</div>
		<p class="text-[15px] font-mono text-neutral-500 uppercase">curve.thinkwright.ai · {m.dataFreshness} run · {formatIsoDate(publishedAt)}</p>
	</div>

	<div class="og-right">
		<div class="og-chart">
			<CurveChart
				data={history.map((p) => p.score)}
				labels={history.map((p) => p.measurementPeriod)}
				dates={history.map((p) => p.periodMidpoint)}
				height={230}
			/>
		</div>
		<div class="og-tiles">
			{#each top as d}
				<div class="og-tile">
					<p class="text-[13px] font-bold uppercase tracking-widest text-neutral-500">{d.fullName}</p>
					<p class="text-[40px] font-bold font-mono tabular-nums leading-none mt-2 {d.status === 'autonomous' ? 'text-rose' : ''}">{d.score.toFixed(1)}</p>
					<p class="text-[13px] font-mono uppercase text-neutral-500 mt-2">{d.status}</p>
				</div>
			{/each}
		</div>
	</div>
</div>

<style>
	.og {
		width: 1200px;
		height: 630px;
		overflow: hidden;
		display: grid;
		grid-template-columns: 640px 1fr;
		background: var(--color-background);
		color: var(--color-primary);
	}
	.og-left {
		padding: 56px 48px 44px 64px;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}
	.og-right {
		padding: 72px 64px 44px 16px;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		border-left: 1px solid var(--color-grid);
	}
	.og-chart :global(text) {
		font-size: 10px;
	}
	.og-tiles {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		border-top: 1px solid var(--color-grid);
	}
	.og-tile {
		padding: 22px 16px 0 0;
	}
</style>
