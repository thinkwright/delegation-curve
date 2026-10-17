<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import BottomNav from '$lib/components/BottomNav.svelte';
	import PageShell from '$lib/components/PageShell.svelte';
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet } = $props();
</script>


<div class="min-h-screen flex flex-col">
	<Header />
	<PageShell>
		{@render children()}
	</PageShell>
	<BottomNav />
</div>
