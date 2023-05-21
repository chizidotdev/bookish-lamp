-- name: CreateItem :one
INSERT INTO items (
    title, buying_price, selling_price, quantity
) VALUES (
    $1, $2, $3, $4
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
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateItem :one
UPDATE items 
SET title = $2, buying_price = $3, selling_price = $4, quantity = $5
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1;

