import { writable } from "svelte/store";

export type ItemBase = {
	title: string;
	buying_price: number;
	selling_price: number;
	quantity: number;
};

export type Item = ItemBase & {
	ID: string;
	CreatedAt: string;
};

const API_URL = 'http://localhost:8080';

export const items = writable<Item[]>([])

async function getItems(size: number) {
	const response = await fetch(`${API_URL}/items?page_id=1&page_size=${size}`);
	const data = (await response.json()) as Item[];

    items.set(data)
}

getItems(15)
