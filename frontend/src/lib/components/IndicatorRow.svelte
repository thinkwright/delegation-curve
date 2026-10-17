<script lang="ts">
	import type { SubIndicator } from '$lib/data/types';
	import { formatWindow } from '$lib/utils/format';

	let { indicator }: { indicator: SubIndicator } = $props();

	const period = $derived(formatWindow(indicator.measurementStart, indicator.measurementEnd));
	const value = $derived(
		typeof indicator.value === 'number' ? indicator.value.toLocaleString() : indicator.value
	);
</script>

<div class="py-2.5 border-b border-grid last:border-0">
	<div class="flex items-baseline">
		<span class="text-sm font-medium text-neutral-600 shrink-0 max-w-[70%]">{indicator.name}</span>
		<div class="leader-dots"></div>
		<span class="font-mono text-sm font-bold tabular-nums shrink-0">{value}{#if indicator.unit}<span class="text-neutral-400 font-normal ml-0.5">{indicator.unit}</span>{/if}</span>
	</div>
	<p class="text-[10px] font-mono text-neutral-400 uppercase mt-1 leading-relaxed">
		{#if indicator.sourceUrl}
			<a href={indicator.sourceUrl} target="_blank" rel="noopener" class="underline decoration-neutral-300 hover:text-neutral-600 transition-colors">{indicator.source}</a>
		{:else}
			<span>{indicator.source}</span>
		{/if}
		{#if period}<span> · {period}</span>{/if}
		{#if indicator.evidenceGrade}<span> · Grade {indicator.evidenceGrade}</span>{/if}
		{#if indicator.confidence}<span> · {indicator.confidence} confidence</span>{/if}
	</p>
	{#if indicator.transform}
		<p class="text-[11px] text-neutral-500 leading-relaxed mt-1">{indicator.transform}</p>
	{/if}
</div>
