<script lang="ts">
	let {
		label,
		value,
		subtitle = '',
		icon = '',
		accent = '',
		href = ''
	}: {
		label: string;
		value: string;
		subtitle?: string;
		icon?: string;
		accent?: string;
		href?: string;
	} = $props();
</script>

{#if href}
<a {href} class="flex flex-col justify-between p-4 border border-grid min-h-[100px] hover:bg-surface transition-colors">
	<div class="flex justify-between items-start">
		<span class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">{label}</span>
		{#if icon}
			<span class="material-symbols-outlined text-base text-neutral-300">{icon}</span>
		{/if}
	</div>
	<div class="mt-auto">
		<span class="text-3xl font-bold font-mono tabular-nums tracking-tighter {accent}">{value}</span>
		{#if subtitle}
			<p class="text-[10px] font-mono text-neutral-400 mt-1 uppercase">{subtitle}</p>
		{/if}
	</div>
</a>
{:else}
<div class="flex flex-col justify-between p-4 border border-grid min-h-[100px]">
	<div class="flex justify-between items-start">
		<span class="text-[10px] font-bold uppercase tracking-widest text-neutral-500">{label}</span>
		{#if icon}
			<span class="material-symbols-outlined text-base text-neutral-300">{icon}</span>
		{/if}
	</div>
	<div class="mt-auto">
		<span class="text-3xl font-bold font-mono tabular-nums tracking-tighter {accent}">{value}</span>
		{#if subtitle}
			<p class="text-[10px] font-mono text-neutral-400 mt-1 uppercase">{subtitle}</p>
		{/if}
	</div>
</div>
{/if}
