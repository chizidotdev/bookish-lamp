-- name: GetInventoryStats :one
SELECT
  (SELECT COUNT(*) FROM items WHERE items.user_id = $1) AS total_items,
  (SELECT items.id FROM items WHERE items.user_id = $1 AND items.quantity <= 5) AS low_stock_items,
  (SELECT COUNT(*) FROM sales WHERE sales.user_id = $1 LIMIT 10) AS recent_sales,
  -- (SELECT SUM(sales.quantity_sold) FROM sales WHERE sales.user_id = $1) AS sales_performance,
  (SELECT COUNT(*) FROM orders WHERE orders.user_id = $1 AND orders.status = 'Pending') AS pending_orders;
  -- (SELECT SUM(items.selling_price * items.quantity) FROM items WHERE items.user_id = $1) AS inventory_value;
