<script lang="ts">
	let {
		filters,
		active,
		onchange
	}: {
		filters: string[];
		active: string;
		onchange: (filter: string) => void;
	} = $props();
</script>

<div class="sticky top-14 z-40 bg-white hairline-b hairline-t flex divide-x divide-black border-black">
	{#each filters as filter}
		<button
			class="flex-1 py-3 px-2 text-xs font-bold tracking-widest uppercase transition-colors
				{active === filter ? 'bg-black text-white' : 'bg-white text-black hover:bg-surface'}"
			onclick={() => onchange(filter)}
		>
			{filter}
		</button>
	{/each}
</div>
