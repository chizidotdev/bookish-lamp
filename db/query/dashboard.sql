-- name: UpdateDashboardTotalItems :exec
UPDATE dashboard
SET total_items = (
    SELECT COUNT(*) FROM items 
    WHERE items.user_id = dashboard.user_id
);

-- name: UpdateDashboardLowStockItems :exec
UPDATE dashboard
SET low_stock_items = (
    SELECT COUNT(*) FROM items 
    WHERE items.user_id = dashboard.user_id
    AND items.stock <= 5
);

-- name: UpdateDashboardItemsToShip :exec
-- UPDATE dashboard
-- SET items_to_ship = (
--     SELECT COUNT(*) FROM items 
--     WHERE items.user_id = dashboard.user_id
--     AND items.status = 'ready_to_ship'
-- );

-- name: UpdateDashboardRecentSales :exec
UPDATE dashboard
SET recent_sales = (
  SELECT COUNT(*) FROM sales
  WHERE sales.user_id = dashboard.user_id
  AND sales.date >= now() - interval '7 days' -- Adjust the timeframe as needed
);

-- name: UpdateDashboardSalesPerformance :exec
UPDATE dashboard
SET sales_performance = (
  SELECT SUM(sales.amount) FROM sales
  WHERE sales.user_id = dashboard.user_id
);

-- name: UpdateDashboardPendingOrders :exec
UPDATE dashboard
SET pending_orders = (
  SELECT COUNT(*) FROM orders
  WHERE orders.user_id = dashboard.user_id
    AND orders.status = 'Pending'
);

-- name: UpdateDashboardNotifications :exec
-- UPDATE dashboard
-- SET notifications = (
--   SELECT COUNT(*) FROM notifications
--   WHERE notifications.user_id = dashboard.user_id
-- );

-- name: UpdateDashboardInventoryValue :exec
UPDATE dashboard
SET inventory_value = (
  SELECT SUM(items.price * items.quantity) FROM items
  WHERE items.user_id = dashboard.user_id
);
