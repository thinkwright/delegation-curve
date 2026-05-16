<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';
	import ThemeToggle from './ThemeToggle.svelte';

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
	<div class="flex h-14 items-center justify-between gap-3 px-4 max-w-2xl mx-auto md:gap-6">
		<!-- Logo — always visible -->
		<div class="flex min-w-0 flex-1 items-center gap-2 overflow-hidden">
			<a href="https://thinkwright.ai" class="shrink-0 text-lg font-normal text-neutral-400 hover:text-black transition-colors whitespace-nowrap" style="letter-spacing: -0.02em;">thinkwright /</a>
			<a href="{base}/" class="min-w-0 truncate text-lg font-bold text-black hover:opacity-60 transition-opacity whitespace-nowrap" style="letter-spacing: -0.02em;">
				AI Delegation Curve
			</a>
		</div>

		<div class="flex items-center gap-1 shrink-0">
			<!-- Desktop nav — hidden on mobile (bottom nav handles it) -->
			<nav class="hidden md:flex items-center gap-5">
				{#each tabs as tab}
					<a
						href={tab.href}
						aria-current={isActive(tab.href) ? 'page' : undefined}
						class="inline-flex h-8 items-center border-b px-0.5 text-[11px] font-bold uppercase tracking-widest transition-colors
							{isActive(tab.href) ? 'border-black text-black' : 'border-transparent text-neutral-400 hover:text-black'}"
					>
						{tab.label}
					</a>
				{/each}
			</nav>
			<ThemeToggle />
		</div>
	</div>
</header>
