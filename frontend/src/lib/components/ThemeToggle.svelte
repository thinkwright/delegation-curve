<script lang="ts">
	let dark = $state(false);

	$effect(() => {
		dark = localStorage.getItem('theme') === 'dark';
	});

	function toggle() {
		dark = !dark;
		localStorage.setItem('theme', dark ? 'dark' : 'light');
		document.documentElement.classList.toggle('dark', dark);
	}
</script>

<button
	onclick={toggle}
	class="p-1.5 text-neutral-400 hover:text-black transition-colors"
	aria-label={dark ? 'Switch to light mode' : 'Switch to dark mode'}
>
	<span class="material-symbols-outlined text-lg">
		{dark ? 'light_mode' : 'dark_mode'}
	</span>
</button>
