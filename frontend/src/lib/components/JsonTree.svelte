<script lang="ts">
	import JsonTree from './JsonTree.svelte';

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	type JsonValue = any;

	let {
		data,
		key = '',
		depth = 0,
		root = true
	}: {
		data: JsonValue;
		key?: string;
		depth?: number;
		root?: boolean;
	} = $props();

	let collapsed = $state(false);
	$effect(() => { collapsed = depth > 1; });

	const isObject = $derived(data !== null && typeof data === 'object' && !Array.isArray(data));
	const isArray = $derived(Array.isArray(data));
	const isExpandable = $derived(isObject || isArray);

	const entries = $derived(
		isObject ? Object.entries(data as Record<string, JsonValue>) :
		isArray ? (data as JsonValue[]).map((v, i) => [String(i), v] as [string, JsonValue]) :
		[]
	);

	const preview = $derived(
		isArray ? `[${(data as JsonValue[]).length}]` :
		isObject ? `{${Object.keys(data as Record<string, JsonValue>).length}}` :
		''
	);

	const valueClass = $derived(
		typeof data === 'string' ? 'text-sage' :
		typeof data === 'number' ? 'text-rose' :
		typeof data === 'boolean' ? 'text-neutral-500 italic' :
		data === null ? 'text-neutral-400 italic' :
		''
	);

	function formatValue(v: JsonValue): string {
		if (typeof v === 'string') return `"${v}"`;
		if (v === null) return 'null';
		return String(v);
	}

	function toggle() {
		collapsed = !collapsed;
	}
</script>

{#if isExpandable}
	<div class="json-node" class:root style="--depth: {depth}">
		<button
			class="json-toggle"
			onclick={toggle}
			aria-expanded={!collapsed}
		>
			<span class="json-chevron" class:collapsed>{#if collapsed}+{:else}&minus;{/if}</span>
			{#if key}<span class="json-key">{key}</span><span class="json-colon">:</span>{/if}
			{#if collapsed}
				<span class="json-preview">{preview}</span>
			{:else}
				<span class="json-bracket">{isArray ? '[' : '{'}</span>
			{/if}
		</button>

		{#if !collapsed}
			<div class="json-children">
				{#each entries as [childKey, childValue], i}
					<div class="json-entry">
						<JsonTree data={childValue} key={isArray ? '' : childKey} depth={depth + 1} root={false} />
						{#if i < entries.length - 1}<span class="json-comma">,</span>{/if}
					</div>
				{/each}
			</div>
			<span class="json-bracket-close">{isArray ? ']' : '}'}</span>
		{/if}
	</div>
{:else}
	<span class="json-leaf">
		{#if key}<span class="json-key">{key}</span><span class="json-colon">:</span>{/if}
		<span class={valueClass}>{formatValue(data)}</span>
	</span>
{/if}

<style>
	.json-node {
		font-family: 'JetBrains Mono', monospace;
		font-size: 12px;
		line-height: 1.7;
	}

	.json-node.root {
		padding: 0;
	}

	.json-toggle {
		display: inline-flex;
		align-items: baseline;
		gap: 0.25rem;
		background: none;
		border: none;
		cursor: pointer;
		padding: 0;
		font-family: inherit;
		font-size: inherit;
		line-height: inherit;
		color: inherit;
		text-align: left;
	}

	.json-toggle:hover .json-chevron {
		color: var(--color-rose);
	}

	.json-chevron {
		display: inline-block;
		width: 1ch;
		text-align: center;
		color: var(--color-primary);
		opacity: 0.3;
		font-weight: bold;
		user-select: none;
	}

	.json-key {
		color: var(--color-primary);
		font-weight: 600;
	}

	.json-colon {
		color: var(--color-primary);
		opacity: 0.3;
		margin-right: 0.5ch;
	}

	.json-preview {
		color: var(--color-primary);
		opacity: 0.3;
		font-style: italic;
	}

	.json-bracket,
	.json-bracket-close {
		color: var(--color-primary);
		opacity: 0.3;
	}

	.json-children {
		padding-left: 1.5rem;
		border-left: 1px solid var(--color-grid);
		margin-left: 0.4ch;
	}

	.json-entry {
		display: flex;
		align-items: baseline;
	}

	.json-comma {
		color: var(--color-primary);
		opacity: 0.3;
	}

	.json-leaf {
		display: inline-flex;
		align-items: baseline;
		gap: 0.25rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 12px;
		line-height: 1.7;
		padding-left: calc(1ch + 0.25rem);
	}
</style>
