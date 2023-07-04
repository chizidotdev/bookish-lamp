-- name: CreateSale :one
INSERT INTO sales (item_id, user_id, quantity_sold, sale_price, customer_name, sale_date)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetSale :one
SELECT *
FROM sales
WHERE (id = $1 AND user_id = $2) LIMIT 1;

-- name: GetSaleForUpdate :one
SELECT *
FROM sales
WHERE (id = $1 AND user_id = $2) LIMIT 1 FOR NO KEY
UPDATE;

-- name: ListSales :many
SELECT s.*, i.title
FROM sales s
         JOIN items i ON s.item_id = i.id
WHERE s.item_id = $1
  AND s.user_id = $2
ORDER BY s.sale_date DESC;

-- name: ListSalesByUserId :many
SELECT s.*, i.title
FROM sales s
         JOIN items i ON s.item_id = i.id
WHERE s.user_id = $1
ORDER BY s.sale_date DESC;

-- name: UpdateSale :one
UPDATE sales
SET quantity_sold = $2,
    sale_price    = $3,
    customer_name = $4,
    sale_date     = $5
WHERE (id = $1 AND user_id = $6) RETURNING *;

-- name: DeleteSale :exec
DELETE
FROM sales
WHERE (id = $1 AND user_id = $2);

-- name: CurrentWeekSales :one
SELECT CAST(COALESCE(SUM(quantity_sold), 0) AS INTEGER) AS total_quantity_sold
FROM sales
WHERE sale_date >= DATE_TRUNC('week', CURRENT_DATE)
  AND sale_date < DATE_TRUNC('week', CURRENT_DATE) + INTERVAL '1 week'
  AND user_id = $1;

-- name: LastWeekSales :one
SELECT CAST(COALESCE(SUM(quantity_sold), 0) AS INTEGER) AS total_quantity_sold
FROM sales
WHERE sale_date >= DATE_TRUNC('week', CURRENT_DATE) - INTERVAL '1 week'
  AND sale_date
    < DATE_TRUNC('week'
    , CURRENT_DATE)
  AND user_id = $1;

-- name: PriceSoldByDate :many
SELECT DATE_TRUNC('day', sale_date) AS date,
    SUM(sale_price) AS total_sale_price
FROM
    sales
WHERE
    user_id = $1
GROUP BY
    DATE_TRUNC('day', sale_date)
ORDER BY
    DATE_TRUNC('day', sale_date);

-- name: PriceSoldByWeek :many
SELECT DATE_TRUNC('week', sale_date) AS date,
    SUM(sale_price) AS total_sale_price
FROM
    sales
WHERE
    user_id = $1
GROUP BY
    DATE_TRUNC('week', sale_date)
ORDER BY
    DATE_TRUNC('week', sale_date);
