<script lang="ts">
	import { items } from '$stores/items';
	import ListItem from '$lib/components/ListItem.svelte';
</script>

<h1>Items</h1>
<ul class="flex flex-col gap-3">
	{#each $items as item}
		<li>
			<ListItem {item} />
		</li>
	{/each}
</ul>