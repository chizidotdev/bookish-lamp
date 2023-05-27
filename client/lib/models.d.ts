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
