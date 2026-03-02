<script lang="ts">
	import { generatePath, generateAreaPath } from '$lib/utils/sparkline';

	let {
		data,
		height = 160,
		color = 'var(--color-primary)',
		showGrid = true,
		endYear
	}: {
		data: number[];
		height?: number;
		color?: string;
		showGrid?: boolean;
		endYear: number;
	} = $props();

	const width = 320;
	const padding = { top: 8, right: 8, bottom: 24, left: 32 };
	const chartW = $derived(width - padding.left - padding.right);
	const chartH = $derived(height - padding.top - padding.bottom);

	// Auto-scale Y axis to data range with padding (min span 15)
	const yRange = $derived(() => {
		const lo = Math.min(...data);
		const hi = Math.max(...data);
		const span = hi - lo;
		const pad = Math.max((span * 0.3), (15 - span) / 2, 3);
		return {
			min: Math.max(0, Math.floor((lo - pad) / 5) * 5),
			max: Math.min(100, Math.ceil((hi + pad) / 5) * 5)
		};
	});

	// Map a value to chart Y coordinate
	const toY = $derived((v: number) => {
		const { min, max } = yRange();
		return chartH - ((v - min) / (max - min)) * chartH;
	});

	// For 2-point data, use a cubic bezier; otherwise use the standard polyline
	const linePath = $derived(() => {
		const { min, max } = yRange();
		if (data.length === 2) {
			const y0 = toY(data[0]);
			const y1 = toY(data[1]);
			return `M0,${y0.toFixed(1)} C${(chartW * 0.4).toFixed(1)},${y0.toFixed(1)} ${(chartW * 0.6).toFixed(1)},${y1.toFixed(1)} ${chartW},${y1.toFixed(1)}`;
		}
		return generatePath(data, chartW, chartH, 0, min, max);
	});
	const areaPath = $derived(() => {
		const line = linePath();
		if (!line) return '';
		return `${line} L${chartW},${chartH} L0,${chartH} Z`;
	});

	// Y-axis ticks: evenly spaced within the auto-scaled range
	const yTicks = $derived(() => {
		const { min, max } = yRange();
		const step = (max - min) / 4;
		return Array.from({ length: 5 }, (_, i) => Math.round(min + i * step));
	});

	// Derive year labels from data length: each point = one year, ending at endYear.
	const startYear = $derived(endYear - data.length + 1);
	const years = $derived(Array.from({ length: data.length }, (_, i) => startYear + i));

	const lastValue = $derived(data[data.length - 1] ?? 0);
	const lastY = $derived(toY(lastValue));
</script>

<div class="w-full" style="height: {height}px;">
	<svg viewBox="0 0 {width} {height}" class="w-full h-full" preserveAspectRatio="xMidYMid meet">
		<g transform="translate({padding.left}, {padding.top})">
			{#if showGrid}
				<!-- Y grid lines and labels -->
				{#each yTicks() as tick}
					{@const y = toY(tick)}
					<line x1="0" y1={y} x2={chartW} y2={y} stroke="var(--color-grid)" stroke-width="0.5" />
					<text x="-4" y={y + 3} text-anchor="end" class="text-[8px] fill-neutral-400" font-family="'JetBrains Mono', monospace">{tick}</text>
				{/each}
				<!-- X labels -->
				{#each years as year, i}
					{@const x = (i / (years.length - 1)) * chartW}
					<text x={x} y={chartH + 16} text-anchor="middle" class="text-[8px] fill-neutral-400" font-family="'JetBrains Mono', monospace">{year}</text>
				{/each}
			{/if}

			<!-- Area fill -->
			<path d={areaPath()} fill={color} opacity="0.04" />

			<!-- Line -->
			<path d={linePath()} fill="none" stroke={color} stroke-width="2" vector-effect="non-scaling-stroke" />

			<!-- Current value dot -->
			<rect x={chartW - 3} y={lastY - 3} width="6" height="6" fill={color} />
		</g>
	</svg>
</div>
