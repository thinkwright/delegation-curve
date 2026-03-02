<script lang="ts">
	import { getDomains } from '$lib/db/client';
	import type { DelegationDomain } from '$lib/data/types';
	import WeightExplorer from '$lib/components/WeightExplorer.svelte';
	import { base } from '$app/paths';

	let domains = $state<DelegationDomain[]>([]);

	$effect(() => {
		getDomains().then(d => { domains = d; }).catch(() => {});
	});
</script>

<svelte:head>
	<title>Methodology — How We Measure AI Decision-Making — AI Delegation Curve</title>
	<meta name="description" content="Methodology for the AI Delegation Curve: how we measure AI autonomy and AI decision-making influence. Min-max normalization, domain weighting, three-tier data sources (transparency reports, regulatory filings, public benchmarks), and credibility principles. Open methodology for tracking AI governance and AI adoption across 9 domains." />
	<link rel="canonical" href="https://curve.thinkwright.ai/about" />
	<meta property="og:title" content="Methodology — How We Measure AI Decision-Making" />
	<meta property="og:description" content="Open methodology for measuring AI delegation: normalization, weighting, data source tiers, and credibility principles." />
	<meta property="og:url" content="https://curve.thinkwright.ai/about" />
	<meta property="og:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
	<meta property="og:image:width" content="1200" />
	<meta property="og:image:height" content="630" />
	<meta property="og:type" content="website" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:image" content="https://curve.thinkwright.ai/og-image.png?v=2" />
</svelte:head>

<div class="px-4 py-8 pb-4">
	<h2 class="text-4xl font-black tracking-tight uppercase leading-none mb-1">Method&shy;ology</h2>
	<p class="font-mono text-xs text-neutral-500 uppercase tracking-widest">How we measure</p>
</div>

<!-- What This Is -->
<div class="px-4 py-6 hairline-b hairline-t">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">What This Is</p>
	<p class="text-sm leading-relaxed mb-4">
		The Keeling Curve has tracked atmospheric CO₂ since 1958. It is a single, continuous measurement that makes the invisible visible. The Delegation Curve seeks to do the same for AI's growing role in consequential decision-making throughout society. One composite score (0–100), updated regularly, so the shape of the curve is the story.
	</p>
	<p class="text-sm leading-relaxed mb-4">
		{domains.length} domains are tracked: content moderation, algorithmic trading, code generation, customer support, credit decisioning, medical diagnostics, legal AI, hiring, and education.
	</p>
	<p class="text-sm leading-relaxed">
		Each domain score is built from 3–4 normalized indicators sourced from transparency reports, regulatory filings, surveys, and public benchmarks.
	</p>
</div>

<!-- Delegation Curve Methodology -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">The Delegation Curve</p>

	<p class="text-sm leading-relaxed mb-4">
		The composite score is a weighted average across {domains.length} decision domains. Each domain score (0–100) represents the percentage of decisions in that sector made or substantially influenced by AI systems.
	</p>

	<p class="text-xs font-bold uppercase tracking-widest text-neutral-500 mb-2">Normalization</p>
	<p class="text-sm leading-relaxed mb-4">
		Raw metrics are normalized to 0–100 using min-max scaling against a historical baseline floor and theoretical automation ceiling. For example, content moderation automated detection was ~24% in 2017 (floor), theoretical ceiling is 100%.
	</p>

	<p class="text-xs font-bold uppercase tracking-widest text-neutral-500 mb-2">Status Tiers</p>
	<div class="space-y-2 mb-4">
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-black">Nominal</span>
			<span class="text-sm text-neutral-600">&lt; 30% AI influence</span>
		</div>
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-sage bg-sage/20 text-sage">Elevated</span>
			<span class="text-sm text-neutral-600">30–60% AI influence</span>
		</div>
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-rose bg-rose/20 text-rose">Autonomous</span>
			<span class="text-sm text-neutral-600">&gt; 60% AI influence</span>
		</div>
	</div>

	<p class="text-xs font-bold uppercase tracking-widest text-neutral-500 mb-2">Domain Weights</p>
	<p class="text-sm leading-relaxed mb-4">
		The composite is a weighted average. Does the weighting matter? Try different lenses — the contribution bars show each domain's weighted impact, sorted by magnitude.
	</p>
	{#if domains.length > 0}
		<WeightExplorer {domains} />
	{/if}
</div>

<!-- Data Sources Overview -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Data Source Tiers</p>

	<div class="space-y-4">
		<div>
			<p class="text-xs font-bold uppercase mb-2">Tier 1 — Direct Measurement</p>
			<p class="text-sm text-neutral-600 leading-relaxed">Platform transparency reports, exchange data, IDE telemetry. Updated quarterly or more frequently. Highest confidence.</p>
		</div>
		<div>
			<p class="text-xs font-bold uppercase mb-2">Tier 2 — Strong Proxies</p>
			<p class="text-sm text-neutral-600 leading-relaxed">FDA databases, industry surveys (Gartner, McKinsey), regulatory filings. Updated annually. Medium confidence.</p>
		</div>
		<div>
			<p class="text-xs font-bold uppercase mb-2">Tier 3 — Directional Signal</p>
			<p class="text-sm text-neutral-600 leading-relaxed">Google Trends, web traffic, app rankings, VC funding data. Updated continuously. Used for interpolation between anchor data points.</p>
		</div>
	</div>
</div>

<!-- Credibility -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Credibility Principles</p>

	<div class="space-y-3">
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">01</span>
			<p class="text-sm leading-relaxed"><strong>Published methodology</strong> — every weight, formula, and source is version-controlled and public.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">02</span>
			<p class="text-sm leading-relaxed"><strong>Raw data access</strong> — underlying indicators are <a href="{base}/data" class="underline hover:text-neutral-500 transition-colors">downloadable</a>.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">03</span>
			<p class="text-sm leading-relaxed"><strong>Confidence bands</strong> — uncertainty is shown, not hidden. Survey-based domains get wider bands.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">04</span>
			<p class="text-sm leading-relaxed"><strong>Methodology versioning</strong> — when weights change, both old and new curves are shown.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">05</span>
			<p class="text-sm leading-relaxed"><strong>No editorializing</strong> — the number is the number. The curve does the talking.</p>
		</div>
	</div>
</div>

<div class="px-4 py-6">
	<p class="font-mono text-xs text-neutral-400 uppercase">AI Delegation Curve, Last updated 28 February 2026</p>
</div>
