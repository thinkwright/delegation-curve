<script lang="ts">
	import type { DelegationDomain } from '$lib/data/types';

	let { domains }: { domains: DelegationDomain[] } = $props();

	const profiles = [
		{ id: 'current', name: 'Current', desc: 'Published Thinkwright methodology' },
		{ id: 'equal', name: 'Equal', desc: 'All domains weighted equally' },
		{ id: 'consumer', name: 'Consumer', desc: 'Emphasizes consumer-facing AI systems' },
		{ id: 'enterprise', name: 'Enterprise', desc: 'Emphasizes financial and regulatory AI' }
	];

	let activeId = $state('current');

	const consumerW: Record<string, number> = {
		'content-mod': 0.20, 'support': 0.20, 'education': 0.15,
		'code-gen': 0.10, 'medical-dx': 0.10, 'hire': 0.10,
		'algo-trade': 0.05, 'credit': 0.05, 'legal-ai': 0.05
	};

	const enterpriseW: Record<string, number> = {
		'algo-trade': 0.25, 'credit': 0.20, 'legal-ai': 0.15,
		'hire': 0.10, 'code-gen': 0.10, 'support': 0.05,
		'content-mod': 0.05, 'medical-dx': 0.05, 'education': 0.05
	};

	function weightFor(domainId: string, defaultW: number): number {
		if (activeId === 'current') return defaultW;
		if (activeId === 'equal') return 1 / domains.length;
		if (activeId === 'consumer') return consumerW[domainId] ?? defaultW;
		if (activeId === 'enterprise') return enterpriseW[domainId] ?? defaultW;
		return defaultW;
	}

	const breakdown = $derived(
		domains
			.map(d => {
				const w = weightFor(d.id, d.weight);
				return {
					id: d.id,
					name: d.name,
					score: d.score,
					weight: w,
					contribution: d.score * w,
					status: d.status
				};
			})
			.sort((a, b) => b.contribution - a.contribution)
	);

	const composite = $derived(breakdown.reduce((s, b) => s + b.contribution, 0));
	const currentComposite = $derived(domains.reduce((s, d) => s + d.score * d.weight, 0));
	const delta = $derived(composite - currentComposite);
	const maxContrib = $derived(Math.max(...breakdown.map(b => b.contribution)));

	const barColor = (status: string) =>
		status === 'autonomous' ? 'bg-rose' :
		status === 'elevated' ? 'bg-sage' :
		'bg-neutral-300';
</script>

<!-- Profile tabs -->
<div class="flex gap-1 mb-4 flex-wrap">
	{#each profiles as p}
		<button
			class="px-2.5 py-1.5 text-[10px] font-bold uppercase tracking-wider transition-colors
				{activeId === p.id ? 'bg-black text-white' : 'bg-neutral-100 text-neutral-500 hover:bg-neutral-200'}"
			onclick={() => activeId = p.id}
		>{p.name}</button>
	{/each}
</div>

<!-- Composite result -->
<div class="flex items-end gap-3 mb-1">
	<span class="text-4xl font-black font-mono tabular-nums tracking-tighter leading-none">{composite.toFixed(1)}</span>
	{#if activeId !== 'current'}
		<span class="text-sm font-mono font-bold tabular-nums mb-0.5
			{delta >= 0 ? 'text-sage' : 'text-rose'}">
			{delta >= 0 ? '+' : ''}{delta.toFixed(1)} vs published
		</span>
	{/if}
</div>
<p class="text-xs text-neutral-500 mb-4">{profiles.find(p => p.id === activeId)?.desc}</p>

<!-- Contribution breakdown -->
<div class="border border-grid">
	{#each breakdown as b}
		<div class="flex items-center gap-2 px-3 py-2 border-b border-grid last:border-0">
			<span class="text-[10px] font-bold uppercase w-20 shrink-0 truncate">{b.name}</span>
			<span class="font-mono text-[10px] tabular-nums text-neutral-400 w-8 shrink-0 text-right">{(b.weight * 100).toFixed(0)}%</span>
			<div class="flex-1 h-1.5 bg-neutral-100 rounded-sm overflow-hidden">
				<div
					class="h-full transition-all duration-300 {barColor(b.status)}"
					style="width: {(b.contribution / maxContrib * 100).toFixed(1)}%"
				></div>
			</div>
			<span class="font-mono text-[10px] tabular-nums w-8 text-right shrink-0">{b.contribution.toFixed(1)}</span>
		</div>
	{/each}
</div>
