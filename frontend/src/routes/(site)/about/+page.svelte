<script lang="ts">
	import WeightExplorer from '$lib/components/WeightExplorer.svelte';
	import Seo from '$lib/components/Seo.svelte';
	import { base } from '$app/paths';
	import { formatIsoDate } from '$lib/utils/format';
	import { breadcrumbs, ogImageUrl } from '$lib/site';

	let { data } = $props();

	const domains = $derived(data.meta.domains ?? []);
	const publishedAt = $derived(data.meta.delegation.publishedAt ?? data.meta.delegation.lastUpdated);
</script>

<Seo
	title="Methodology — How We Measure AI Influence — AI Delegation Curve"
	description="Methodology for the AI Delegation Curve: how we estimate AI influence and delegated workflow share. Indicator normalization, evidence grades, domain weighting, source quality, and curve maintenance across 9 domains."
	path="/about"
	image={ogImageUrl(data.meta.delegation.runId)}
	modifiedAt={publishedAt}
	jsonLd={breadcrumbs([{ name: 'Methodology', path: '/about' }])}
/>


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
		Each domain score is built from 2-5 normalized indicators sourced from transparency reports, regulatory filings, surveys, public datasets, and research.
	</p>
</div>

<!-- Curve Maintenance -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Curve Maintenance</p>
	<p class="text-sm leading-relaxed mb-4">
		The public curve is kept on one measuring stick over time. When sources or scoring rules improve, earlier public points may be restated before showing headline movement.
	</p>
	<p class="text-sm leading-relaxed">
		Older published points are retained in the dataset for audit, but the main chart uses the current measurement series.
	</p>
	<p class="text-sm leading-relaxed mt-4">
		Every restatement and correction, with its effect on the composite, is recorded in the public <a href="https://github.com/thinkwright/delegation-curve/blob/main/docs/methodology/restatement-policy.md" target="_blank" rel="noopener" class="underline hover:text-neutral-500 transition-colors">restatement and correction log</a>.
	</p>
</div>

<!-- Delegation Curve Methodology -->
<div class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">The Delegation Curve</p>

	<p class="text-sm leading-relaxed mb-4">
		The composite score is a weighted average across {domains.length} decision domains. Each domain score (0–100) estimates AI influence in that sector using direct delegation measures where available and proxy signals where direct measurement is still incomplete.
	</p>

	<p class="text-xs font-bold uppercase tracking-widest text-neutral-500 mb-2">Normalization</p>
	<p class="text-sm leading-relaxed mb-4">
		Raw metrics are normalized to 0–100 before weighting. Most inputs are direct percentages; count and reach metrics use fixed caps or log scaling so they can be compared with percentage-based indicators without implying they are literal decision shares.
	</p>

	<p class="text-xs font-bold uppercase tracking-widest text-neutral-500 mb-2">Status Tiers</p>
	<div class="space-y-2 mb-4">
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-black">Nominal</span>
			<span class="text-sm text-neutral-600">Score below 40</span>
		</div>
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-sage bg-sage/20 text-sage">Elevated</span>
			<span class="text-sm text-neutral-600">Score 40 to 74.9</span>
		</div>
		<div class="flex items-center gap-2">
			<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border border-rose bg-rose/20 text-rose">Autonomous</span>
			<span class="text-sm text-neutral-600">Score 75 and above</span>
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
			<p class="text-xs font-bold uppercase mb-2">Tier 1 — Direct Workflow Evidence</p>
			<p class="text-sm text-neutral-600 leading-relaxed">Platform transparency reports, operational datasets, regulatory or filing disclosures, and first-party telemetry. Highest confidence when definitions are stable.</p>
		</div>
		<div>
			<p class="text-xs font-bold uppercase mb-2">Tier 2 — Strong Proxies</p>
			<p class="text-sm text-neutral-600 leading-relaxed">Official databases, industry surveys, product-market denominators, and adoption reports. Used when direct delegated-decision data is not yet available.</p>
		</div>
		<div>
			<p class="text-xs font-bold uppercase mb-2">Tier 3 — Context & Guardrails</p>
			<p class="text-sm text-neutral-600 leading-relaxed">Public benchmarks, research studies, quality reports, and reliability evidence. Used to interpret confidence and scope, not as standalone proof of social delegation.</p>
		</div>
	</div>
</div>

<!-- Evidence Grades -->
<div id="evidence-grades" class="px-4 py-6 hairline-b">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-3">Evidence Grades</p>
	<p class="text-sm leading-relaxed mb-4">Every scored indicator carries a grade for how directly its source measures delegation, a confidence level, and the period it measures. Grades and sources are shown on each domain page.</p>
	<div class="space-y-3">
		<div class="flex gap-3"><span class="font-mono font-bold shrink-0 w-4">A</span><p class="text-sm text-neutral-600 leading-relaxed">Direct observed delegated decisions, task resolutions, operational outcomes, regulatory disclosures, or platform telemetry.</p></div>
		<div class="flex gap-3"><span class="font-mono font-bold shrink-0 w-4">B</span><p class="text-sm text-neutral-600 leading-relaxed">Workflow penetration, operational deployment, official databases, or strong representative surveys.</p></div>
		<div class="flex gap-3"><span class="font-mono font-bold shrink-0 w-4">C</span><p class="text-sm text-neutral-600 leading-relaxed">AI-attributable output, value, or productivity share, or narrower proxy evidence. Used with caveats.</p></div>
		<div class="flex gap-3"><span class="font-mono font-bold shrink-0 w-4">D</span><p class="text-sm text-neutral-600 leading-relaxed">Adoption, sentiment, capability benchmarks, or broad context. Scored only through an explicit, documented conversion.</p></div>
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
			<p class="text-sm leading-relaxed"><strong>Source confidence</strong> — uncertainty is documented, not hidden. Survey-based and proxy-heavy domains are treated more cautiously.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">04</span>
			<p class="text-sm leading-relaxed"><strong>Continuous series</strong> — public prior points stay on the current measuring stick; archived publication points remain available for audit.</p>
		</div>
		<div class="flex gap-3">
			<span class="text-base font-mono font-bold tabular-nums shrink-0 text-neutral-300">05</span>
			<p class="text-sm leading-relaxed"><strong>Method before interpretation</strong> — sources, weights, and formulas are separated from narrative conclusions.</p>
		</div>
	</div>
</div>

<div class="px-4 py-6">
	<p class="font-mono text-xs text-neutral-400 uppercase">AI Delegation Curve, Last updated {formatIsoDate(publishedAt)}</p>
</div>
