export type ItemBase = {
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
};

type Item = ItemBase & {
    ID: string;
    CreatedAt: string;
};

const BASE_URL = import.meta.env.VITE_API_URL;

export const getItems = async (): Promise<Item[]> => {
    const response = await fetch(`${BASE_URL}/items`);
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

export const deleteItem = async (id: string): Promise<string> => {
    const response = await fetch(`${BASE_URL}/items/${id}`, {
        method: 'DELETE',
    });
    const data = await response.json();
    return data;
};
