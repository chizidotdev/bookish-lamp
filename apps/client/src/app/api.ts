export type ItemBase = {
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
};

export type Item = ItemBase & {
    ID: string;
    CreatedAt: string;
};

export type SaleBase = {
    quantity_sold: number;
    unit_price: number;
    sale_date: string;
};

type Sale = SaleBase & {
    CreatedAt: string;
    UpdatedAt: string;
    ID: string;
    ItemID: string;
};

const BASE_URL = import.meta.env.VITE_API_URL;

export const getItems = async (): Promise<Item[]> => {
    const response = await fetch(`${BASE_URL}/items`);
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

export const updateItem = async (
    item: ItemBase & { id: string }
): Promise<Item> => {
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

export const getItemSales = async (id: string): Promise<Sale[]> => {
    const response = await fetch(`${BASE_URL}/items/${id}/sales`);
    const data = await response.json();
    return data;
};

export const addSale = async (
    item: SaleBase & { id: string }
): Promise<Item> => {
    const { id, ...rest } = item;
    const response = await fetch(`${BASE_URL}/items/${id}/sales`, {
        method: 'POST',
        body: JSON.stringify(rest),
    });
    const data = await response.json();
    return data;
};
