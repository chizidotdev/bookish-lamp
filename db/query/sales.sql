-- name: CreateSale :one
INSERT INTO sales (
    item_id, quantity_sold, sale_price, customer_name, sale_date 
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSale :one
SELECT * FROM sales
WHERE (id = $1 AND item_id = $2)
LIMIT 1;

-- name: GetSaleForUpdate :one
SELECT * FROM sales
WHERE (id = $1 AND item_id = $2)
LIMIT 1 FOR NO KEY UPDATE;

-- name: ListSales :many
SELECT * FROM sales
WHERE item_id = $1
ORDER BY sale_date DESC;

-- name: UpdateSale :one
UPDATE sales 
SET quantity_sold = $2,
sale_price = $3,
customer_name = $4,
sale_date = $5
WHERE (id = $1 AND item_id = $6)
RETURNING *;

-- name: DeleteSale :exec
DELETE FROM sales WHERE (id = $1 AND item_id = $2);

