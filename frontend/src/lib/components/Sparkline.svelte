<script lang="ts">
	import { generatePath } from '$lib/utils/sparkline';

	let {
		data,
		color = 'var(--color-primary)',
		height = 24
	}: {
		data: number[];
		color?: string;
		height?: number;
	} = $props();

	const path = $derived(generatePath(data, 100, height));
</script>

<div class="h-6 w-full relative overflow-hidden" aria-hidden="true" style="height: {height}px;">
	<svg class="absolute bottom-0 left-0 w-full h-full" viewBox="0 0 100 {height}" preserveAspectRatio="none">
		<path d={path} fill="none" stroke={color} stroke-width="1.5" vector-effect="non-scaling-stroke" />
	</svg>
</div>
