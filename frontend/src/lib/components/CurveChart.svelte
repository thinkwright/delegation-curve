<script lang="ts">
	let {
		data,
		labels,
		seriesKeys,
		height = 160,
		color = 'var(--color-primary)',
		showGrid = true,
		endYear
	}: {
		data: number[];
		labels?: string[];
		seriesKeys?: string[];
		height?: number;
		color?: string;
		showGrid?: boolean;
		endYear?: number;
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

	const activeSeriesKeys = $derived(
		seriesKeys?.length === data.length ? seriesKeys : data.map(() => 'default')
	);
	const segments = $derived(() => {
		const result: number[][] = [];
		for (let i = 0; i < data.length; i += 1) {
			if (i === 0 || activeSeriesKeys[i] !== activeSeriesKeys[i - 1]) {
				result.push([i]);
			} else {
				result[result.length - 1].push(i);
			}
		}
		return result;
	});
	const breakIndexes = $derived(() => {
		const indexes: number[] = [];
		for (let i = 1; i < data.length; i += 1) {
			if (activeSeriesKeys[i] !== activeSeriesKeys[i - 1]) indexes.push(i);
		}
		return indexes;
	});

	// Y-axis ticks: evenly spaced within the auto-scaled range
	const yTicks = $derived(() => {
		const { min, max } = yRange();
		const step = (max - min) / 4;
		return Array.from({ length: 5 }, (_, i) => Math.round(min + i * step));
	});

	// Prefer explicit run labels, with year labels retained for legacy trend arrays.
	const startYear = $derived((endYear ?? new Date().getUTCFullYear()) - data.length + 1);
	const xLabels = $derived(() => {
		if (labels?.length === data.length) return labels;
		return Array.from({ length: data.length }, (_, i) => String(startYear + i));
	});
	const toX = (index: number, total: number) => total <= 1 ? chartW : (index / (total - 1)) * chartW;
	const linePathFor = (indexes: number[]) => {
		if (!indexes.length) return '';
		if (indexes.length === 1) {
			const x = toX(indexes[0], data.length);
			const y = toY(data[indexes[0]]);
			return `M${(x - 6).toFixed(1)},${y.toFixed(1)} L${(x + 6).toFixed(1)},${y.toFixed(1)}`;
		}
		return indexes
			.map((index, i) => {
				const x = toX(index, data.length);
				const y = toY(data[index]);
				return `${i === 0 ? 'M' : 'L'}${x.toFixed(1)},${y.toFixed(1)}`;
			})
			.join(' ');
	};
	const areaPathFor = (indexes: number[]) => {
		if (indexes.length < 2) return '';
		const line = linePathFor(indexes);
		const firstX = toX(indexes[0], data.length);
		const lastX = toX(indexes[indexes.length - 1], data.length);
		return `${line} L${lastX.toFixed(1)},${chartH} L${firstX.toFixed(1)},${chartH} Z`;
	};

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
				{#each xLabels() as label, i}
					{@const x = toX(i, xLabels().length)}
					<text x={x} y={chartH + 16} text-anchor="middle" class="text-[8px] fill-neutral-400" font-family="'JetBrains Mono', monospace">{label}</text>
				{/each}
			{/if}

			{#each breakIndexes() as index}
				{@const x = (toX(index - 1, data.length) + toX(index, data.length)) / 2}
				<line x1={x} y1="0" x2={x} y2={chartH} stroke="var(--color-grid)" stroke-width="1" stroke-dasharray="3 3" />
				<text x={x + 3} y="10" class="text-[7px] fill-neutral-400 uppercase" font-family="'JetBrains Mono', monospace">new series</text>
			{/each}

			<!-- Area fill -->
			{#each segments() as segment}
				{#if areaPathFor(segment)}
					<path d={areaPathFor(segment)} fill={color} opacity="0.04" />
				{/if}
			{/each}

			<!-- Line -->
			{#each segments() as segment}
				<path d={linePathFor(segment)} fill="none" stroke={color} stroke-width="2" vector-effect="non-scaling-stroke" />
			{/each}

			<!-- Current value dot -->
			<rect x={chartW - 3} y={lastY - 3} width="6" height="6" fill={color} />
		</g>
	</svg>
</div>
