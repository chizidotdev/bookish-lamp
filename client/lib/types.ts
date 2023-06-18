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

export type SaleBase = {
    quantity_sold: number;
    sale_price: number;
    sale_date: string;
    customer_name: string;
};

export type Sale = SaleBase & {
    id: string;
    item_id: string;
    user_id: string;
    created_at: string;
    updated_at: string;
};
