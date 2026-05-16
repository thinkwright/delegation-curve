<script lang="ts">
	import '../app.css';
	import Header from '$lib/components/Header.svelte';
	import BottomNav from '$lib/components/BottomNav.svelte';
	import PageShell from '$lib/components/PageShell.svelte';
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet } = $props();
</script>

<svelte:head>
	<title>AI Delegation Curve</title>
	<meta name="description" content="The AI Delegation Curve is a composite index estimating AI influence and delegated workflow share across 9 consequential domains including content moderation, algorithmic trading, code generation, medical diagnostics, credit scoring, legal AI, hiring automation, and education." />
	<meta name="keywords" content="AI delegation, AI autonomy index, AI decision making, AI governance metrics, AI adoption tracker, AI influence measurement, artificial intelligence oversight, algorithmic decision making, AI transparency, AI accountability, autonomous AI systems, AI safety metrics, AI risk assessment, machine learning adoption, AI automation index, content moderation AI, algorithmic trading, AI code generation, AI medical diagnosis, AI credit scoring, AI hiring, AI legal, AI education" />
</svelte:head>

<div class="min-h-screen flex flex-col">
	<Header />
	<PageShell>
		{@render children()}
	</PageShell>
	<BottomNav />
</div>
