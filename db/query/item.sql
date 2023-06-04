-- name: CreateItem :one
INSERT INTO items (
    user_id, title, buying_price, selling_price, quantity
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetItem :one
SELECT * FROM items
WHERE id = $1 LIMIT 1;

-- name: GetItemForUpdate :one
SELECT * FROM items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListItems :many
SELECT * FROM items
WHERE user_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateItem :one
UPDATE items 
SET title = $2, buying_price = $3, selling_price = $4, quantity = $5
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1;

