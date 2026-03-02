<script lang="ts">
	import '../app.css';
	import Header from '$lib/components/Header.svelte';
	import BottomNav from '$lib/components/BottomNav.svelte';
	import PageShell from '$lib/components/PageShell.svelte';
	import { initDB, getInitError } from '$lib/db/boot';
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet } = $props();

	let dbFailed = $state(false);

	// Eagerly start DB init so it's warming up while the page renders.
	// Capture failure so we can show a global error instead of infinite spinner.
	initDB().catch(() => { dbFailed = true; });
</script>

<svelte:head>
	<title>AI Delegation Curve</title>
	<meta name="description" content="The AI Delegation Curve tracks what percentage of consequential decisions are delegated to AI systems. A composite index across 9 domains including content moderation, algorithmic trading, code generation, medical diagnostics, credit scoring, legal AI, hiring automation, and education. Measuring AI autonomy, AI governance, and AI decision-making influence." />
	<meta name="keywords" content="AI delegation, AI autonomy index, AI decision making, AI governance metrics, AI adoption tracker, AI influence measurement, artificial intelligence oversight, algorithmic decision making, AI transparency, AI accountability, autonomous AI systems, AI safety metrics, AI risk assessment, machine learning adoption, AI automation index, content moderation AI, algorithmic trading, AI code generation, AI medical diagnosis, AI credit scoring, AI hiring, AI legal, AI education" />
</svelte:head>

<div class="min-h-screen flex flex-col">
	<Header />
	<PageShell>
		{#if dbFailed}
			<div class="px-4 pt-8 pb-6 flex flex-col items-center justify-center min-h-[60vh]">
				<span class="material-symbols-outlined text-3xl text-neutral-300">warning</span>
				<p class="text-xs font-bold uppercase tracking-widest text-neutral-400 mt-3">Unable to load data engine</p>
				<p class="text-xs text-neutral-400 mt-1 font-mono">{getInitError()?.message ?? 'Unknown error'}</p>
				<button
					class="mt-4 px-4 py-2 text-xs font-bold uppercase tracking-wider bg-black text-white hover:bg-neutral-800 transition-colors"
					onclick={() => location.reload()}
				>Reload page</button>
			</div>
		{:else}
			{@render children()}
		{/if}
	</PageShell>
	<BottomNav />
</div>
