import { BASE_URL } from '~lib/constants';
import type { Item, ItemBase } from '~lib/types';

export const getItems = async (): Promise<Item[]> => {
    const response = await fetch(`${BASE_URL}/items`, {
        credentials: 'include',
    });
    const data = await response.json();
    return data;
};

export const getItemById = async (id: string): Promise<Item | undefined> => {
    if (!id) return;
    const response = await fetch(`${BASE_URL}/items/${id}`);
    const data = await response.json();
    return data;
};

export const newItem = async (item: ItemBase): Promise<Item> => {
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
    console.log('response', response);
    const data = await response.json();
    return data;
};
