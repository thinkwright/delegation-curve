<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';
	import ThemeToggle from './ThemeToggle.svelte';

	const segments = $derived(page.url.pathname.replace(base, '').split('/').filter(Boolean));
	const isDetail = $derived(segments.length > 1);

	const tabs = [
		{ href: `${base}/delegation`, label: 'Domains' },
		{ href: `${base}/data`, label: 'Data' },
		{ href: `${base}/about`, label: 'About' }
	];

	function isActive(href: string): boolean {
		if (href === `${base}/`) return page.url.pathname === `${base}/` || page.url.pathname === base;
		return page.url.pathname.startsWith(href);
	}
</script>

<header class="sticky top-0 z-50 bg-white hairline-b">
	<div class="flex h-14 items-center justify-between px-4 max-w-2xl mx-auto">
		<!-- Logo — always visible -->
		<div class="flex items-center gap-2 min-w-0">
			<a href="https://thinkwright.ai" class="shrink-0 text-lg font-normal text-neutral-400 hover:text-black transition-colors whitespace-nowrap" style="letter-spacing: -0.02em;">thinkwright /</a>
			<a href="{base}/" class="text-lg font-bold text-black hover:opacity-60 transition-opacity whitespace-nowrap" style="letter-spacing: -0.02em;">
				AI Delegation Curve
			</a>

			{#if isDetail}
				<span class="text-neutral-300">/</span>
				<a href="{base}/{segments.slice(0, -1).join('/')}" class="text-xs font-bold uppercase tracking-widest text-neutral-400 hover:text-black transition-colors">
					{segments[0]}
				</a>
			{/if}
		</div>

		<div class="flex items-center gap-1 shrink-0">
			<!-- Desktop nav — hidden on mobile (bottom nav handles it) -->
			<nav class="hidden md:flex items-center gap-1">
				{#each tabs as tab}
					<a
						href={tab.href}
						class="px-3 py-1.5 text-xs font-bold uppercase tracking-widest transition-colors
							{isActive(tab.href) ? 'bg-black text-white' : 'text-neutral-400 hover:text-black'}"
					>
						{tab.label}
					</a>
				{/each}
			</nav>
			<ThemeToggle />
		</div>
	</div>
</header>
