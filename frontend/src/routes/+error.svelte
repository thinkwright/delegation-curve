<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';
</script>

<svelte:head>
	<title>{page.status} — AI Delegation Curve</title>
	<meta name="robots" content="noindex" />
</svelte:head>

<div class="min-h-screen flex flex-col items-center justify-center px-4 text-center">
	<p class="text-[10px] font-bold uppercase tracking-widest text-neutral-500 mb-2">{page.status}</p>
	<h1 class="text-3xl font-black tracking-tight uppercase leading-none mb-3">
		{page.status === 404 ? 'Page not found' : 'Something went wrong'}
	</h1>
	<p class="text-xs text-neutral-500 mb-6 max-w-xs">{page.error?.message ?? ''}</p>
	<a href="{base}/" class="text-sm font-bold uppercase tracking-wide underline">Back to the curve</a>
</div>
