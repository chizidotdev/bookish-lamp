export type NextPage = {
    params: { slug: string };
    searchParams: { [key: string]: string | undefined };
};

export type ItemBase = {
    title: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
};

export type Item = ItemBase & {
    id: string;
    user_id: string;
    created_at: string;
};

export type User = {
    id: string;
    email: string;
};

export type Inventory = {
    total_items: number;
    low_stock_items: number;
    recent_sales: number;
    pending_orders: number;
    sales_performance: number;
};
