<script lang="ts">
	import type { DelegationDomain } from '$lib/data/types';
	import { base } from '$app/paths';
	import { formatDelta } from '$lib/utils/format';

	let { domain, maxDelta = 1 }: { domain: DelegationDomain; maxDelta?: number } = $props();

	const statusColor = $derived(
		domain.status === 'autonomous' ? 'text-rose' :
		domain.status === 'elevated' ? 'text-sage' :
		'text-neutral-400'
	);

	const dotColor = $derived(
		domain.status === 'autonomous' ? 'bg-rose' :
		domain.status === 'elevated' ? 'bg-sage' :
		'bg-neutral-300'
	);

	const borderClass = $derived(
		domain.status === 'autonomous' ? 'border-l-4 border-l-rose bg-rose/5' :
		''
	);

	const delta = $derived(domain.score - domain.previousScore);

	// Curve: cubic bezier from previousScore → currentScore
	const H = 44;
	const PAD = 3;
	const curveLine = $derived(() => {
		const lo = Math.min(domain.previousScore, domain.score) - PAD;
		const hi = Math.max(domain.previousScore, domain.score) + PAD;
		const range = hi - lo || 1;
		const y = (v: number) => H - ((v - lo) / range) * H;
		const y0 = y(domain.previousScore);
		const y1 = y(domain.score);
		return `M0,${y0.toFixed(1)} C40,${y0.toFixed(1)} 60,${y1.toFixed(1)} 100,${y1.toFixed(1)}`;
	});

	const highlightColor = $derived(
		domain.status === 'autonomous' ? 'var(--color-rose)' :
		domain.status === 'elevated' ? 'var(--color-sage)' :
		'var(--color-primary)'
	);

	// Highlighted segment: proportional to delta vs largest delta across all domains
	// Longest highlight = fastest-growing domain
	const highlightPct = $derived((delta / maxDelta) * 60);
	const clipX = $derived(100 - highlightPct);
</script>

<a
	href="{base}/delegation/{domain.id}"
	class="group relative flex flex-col p-4 hairline-b hover:bg-surface transition-colors cursor-pointer {borderClass}"
>
	<div class="flex justify-between items-start mb-2">
		<div class="flex flex-col">
			<span class="text-lg font-bold tracking-tight uppercase">{domain.name}</span>
			<span class="font-mono text-xs text-neutral-500 mt-1">{domain.fullName}</span>
		</div>
		<div class="flex flex-col items-end">
			<div class="flex items-center gap-2">
				<span class="text-2xl font-bold font-mono tracking-tighter tabular-nums {domain.status === 'autonomous' ? 'text-rose' : ''}">{domain.score}</span>
				<div class="w-2 h-2 {dotColor} mb-1"></div>
			</div>
			<span class="text-[10px] font-bold uppercase tracking-wider {statusColor}">{domain.status}</span>
		</div>
	</div>

	<!-- Bezier curve: base in grey, trailing segment highlighted for delta -->
	<div class="relative h-11 w-full">
		<svg class="w-full h-full" viewBox="0 0 100 {H}" preserveAspectRatio="none">
			<defs>
				<clipPath id="delta-{domain.id}">
					<rect x={clipX} y="0" width={highlightPct} height={H} />
				</clipPath>
			</defs>
			<!-- Base curve in light grey -->
			<path
				d={curveLine()}
				fill="none"
				stroke="var(--color-curve-base)"
				stroke-width="1.5"
				vector-effect="non-scaling-stroke"
			/>
			<!-- Delta segment: same path, clipped to trailing portion -->
			<path
				d={curveLine()}
				fill="none"
				stroke={highlightColor}
				stroke-width="2"
				vector-effect="non-scaling-stroke"
				clip-path="url(#delta-{domain.id})"
			/>
		</svg>
	</div>
	<div class="flex justify-between">
		<span class="text-[10px] font-mono text-neutral-400">{domain.previousScore} prior</span>
		<span class="text-[10px] font-mono text-neutral-400">{formatDelta(delta)} vs prior</span>
	</div>
</a>
