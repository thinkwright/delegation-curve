<script lang="ts">
	import type { Snippet } from 'svelte';
	import Footer from './Footer.svelte';

	let { children }: { children: Snippet } = $props();
</script>

<main class="flex-1 flex flex-col w-full max-w-md md:max-w-2xl mx-auto border-x border-grid min-h-screen bg-white relative">
	{@render children()}
	<div class="mt-auto">
		<Footer />
	</div>
	<!-- Bottom nav clearance — mobile only, desktop has no bottom nav -->
	<div class="h-16 md:h-4 w-full"></div>
</main>
