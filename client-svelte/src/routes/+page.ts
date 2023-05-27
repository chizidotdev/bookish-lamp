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

export async function load() {
    const response = await fetch(`http://localhost:8080/items?page_id=1&page_size=15`);
    const posts = await response.json() as Item[]

    return {
        posts
    };
}
