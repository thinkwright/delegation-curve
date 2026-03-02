<script lang="ts">
	let { status, label }: { status: string; label?: string } = $props();

	const displayLabel = $derived(label ?? status);

	const classes: Record<string, string> = {
		nominal: 'bg-transparent border-black text-black',
		elevated: 'bg-sage/20 border-sage text-sage',
		autonomous: 'bg-rose/20 border-rose text-rose',
		healthy: 'bg-sage/20 border-sage text-sage',
		caution: 'bg-amber-100 border-amber-500 text-amber-700',
		risk: 'bg-rose/20 border-rose text-rose',
		widening: 'bg-rose/20 border-rose text-rose',
		stable: 'bg-sage/20 border-sage text-sage'
	};
</script>

<span class="inline-block px-2 py-0.5 text-[10px] font-bold uppercase tracking-wider border {classes[status] ?? classes.nominal}">
	{displayLabel}
</span>
