<script lang="ts">
	let {
		label,
		value,
		unit = ''
	}: {
		label: string;
		value: string | number;
		unit?: string;
	} = $props();
</script>

<div class="flex items-baseline py-2">
	<span class="text-sm font-medium text-neutral-600 shrink-0">{label}</span>
	<div class="leader-dots"></div>
	<span class="font-mono text-sm font-bold tabular-nums shrink-0">{value}{#if unit}<span class="text-neutral-400 font-normal ml-0.5">{unit}</span>{/if}</span>
</div>
