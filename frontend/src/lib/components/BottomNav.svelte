<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';

	const tabs = [
		{ href: `${base}/`, icon: 'dashboard', label: 'Home' },
		{ href: `${base}/delegation`, icon: 'list', label: 'Domains' },
		{ href: `${base}/about`, icon: 'info', label: 'About' }
	];


	function isActive(href: string): boolean {
		if (href === `${base}/`) return page.url.pathname === `${base}/` || page.url.pathname === base;
		return page.url.pathname.startsWith(href);
	}
</script>

<!-- Mobile only — desktop uses header nav -->
<nav class="md:hidden fixed bottom-0 w-full max-w-md left-1/2 -translate-x-1/2 bg-white/95 backdrop-blur-sm hairline-t h-14 flex z-40">
	{#each tabs as tab}
		<a
			href={tab.href}
			class="flex flex-1 flex-col items-center justify-center transition-colors
				{isActive(tab.href) ? 'bg-black text-white' : 'text-neutral-400 hover:bg-surface hover:text-black'}"
		>
			<span class="material-symbols-outlined text-xl">{tab.icon}</span>
		</a>
	{/each}
</nav>
