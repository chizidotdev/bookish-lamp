import type { Item, ItemBase } from '~lib/models';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const getItems = async (): Promise<Item[]> => {
    const response = await fetch(`${BASE_URL}/items?page_id=1&page_size=5`);
    const data = await response.json();
    return data;
};

export const getItemById = async (id: string): Promise<Item> => {
    const response = await fetch(`${BASE_URL}/items/${id}`);
    const data = await response.json();
    return data;
};

export const addItem = async (item: ItemBase): Promise<Item> => {
    const response = await fetch(`${BASE_URL}/items`, {
        method: 'POST',
        body: JSON.stringify(item),
    });
    const data = await response.json();
    return data;
};

export const updateItem = async (item: ItemBase & { id: string }): Promise<Item> => {
    const { id, ...rest } = item;
    const response = await fetch(`${BASE_URL}/items/${id}`, {
        method: 'PUT',
        body: JSON.stringify(rest),
    });
    const data = await response.json();
    return data;
};

export const deleteItem = async (id: string): Promise<string> => {
    const response = await fetch(`${BASE_URL}/items/${id}`, {
        method: 'DELETE',
    });
    const data = await response.json();
    return data;
};
