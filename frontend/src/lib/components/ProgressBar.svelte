<script lang="ts">
	let {
		value,
		max = 100,
		color = 'bg-black'
	}: {
		value: number;
		max?: number;
		color?: string;
	} = $props();

	const pct = $derived(Math.min(100, (value / max) * 100));
</script>

<div class="w-full bg-grid h-1.5">
	<div class="{color} h-full transition-all duration-300" style="width: {pct}%"></div>
</div>
