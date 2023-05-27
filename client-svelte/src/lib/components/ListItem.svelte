<script lang="ts">
	import { enhance, type SubmitFunction } from '$app/forms';
	import { page } from '$app/stores';
	import '$lib/app.css';
	import type { Item } from '../../routes/+page';

	export let item: Item;

	const { ID, title, buying_price, selling_price, quantity, CreatedAt } = item;
</script>

<div class="card bg-base-200 shadow-md">
	<div class="card-body flex-row items-center justify-between gap-1 mx-0 py-5">
		<div class="flex-1 flex flex-col cursor-pointer">
			<div class="indicator">
				<span
					class={`indicator-item indicator-start badge badge-sm ${
						!quantity ? 'badge-error' : 'badge-secondary'
					}`}
				>
					{quantity}
				</span>
				<h2 class="card-title text-lg">{title}</h2>
			</div>
			<div class="flex gap-2 mt-1">
				<div class="badge">₦{buying_price}</div>
				<div class="badge badge-accent">
					₦{selling_price}
				</div>
			</div>
		</div>
		<div class="card-actions flex-row gap-5">
			<!-- <Link to={`/items/edit/${ID}`}> -->
			<!--     <FaEdit class="text-secondary" /> -->
			<!-- </Link> -->
			<!-- <Link to={`/items/delete/${ID}`}> -->
			<!--     <FaTrash class="text-error" /> -->
			<!-- </Link> -->
		</div>
	</div>
</div>
