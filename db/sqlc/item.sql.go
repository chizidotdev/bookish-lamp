// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: item.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
    title, buying_price, selling_price, quantity
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, title, buying_price, selling_price, quantity, created_at
`

type CreateItemParams struct {
	Title        string  `json:"title"`
	BuyingPrice  float32 `json:"buying_price"`
	SellingPrice float32 `json:"selling_price"`
	Quantity     int64   `json:"quantity"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Title,
		arg.BuyingPrice,
		arg.SellingPrice,
		arg.Quantity,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.BuyingPrice,
		&i.SellingPrice,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const getItem = `-- name: GetItem :one
SELECT id, title, buying_price, selling_price, quantity, created_at FROM items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetItem(ctx context.Context, id uuid.UUID) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.BuyingPrice,
		&i.SellingPrice,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const getItemForUpdate = `-- name: GetItemForUpdate :one
SELECT id, title, buying_price, selling_price, quantity, created_at FROM items
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetItemForUpdate(ctx context.Context, id uuid.UUID) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemForUpdate, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.BuyingPrice,
		&i.SellingPrice,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const listItems = `-- name: ListItems :many
SELECT id, title, buying_price, selling_price, quantity, created_at FROM items
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListItemsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListItems(ctx context.Context, arg ListItemsParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.BuyingPrice,
			&i.SellingPrice,
			&i.Quantity,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :one
UPDATE items 
SET title = $2, buying_price = $3, selling_price = $4, quantity = $5
WHERE id = $1
RETURNING id, title, buying_price, selling_price, quantity, created_at
`

type UpdateItemParams struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	BuyingPrice  float32   `json:"buying_price"`
	SellingPrice float32   `json:"selling_price"`
	Quantity     int64     `json:"quantity"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.ID,
		arg.Title,
		arg.BuyingPrice,
		arg.SellingPrice,
		arg.Quantity,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.BuyingPrice,
		&i.SellingPrice,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}
