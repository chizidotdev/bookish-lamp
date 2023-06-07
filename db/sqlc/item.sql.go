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
    user_id, title, buying_price, selling_price, quantity
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, title, buying_price, selling_price, quantity, user_id, created_at
`

type CreateItemParams struct {
	UserID       uuid.UUID `json:"user_id"`
	Title        string    `json:"title"`
	BuyingPrice  float32   `json:"buying_price"`
	SellingPrice float32   `json:"selling_price"`
	Quantity     int64     `json:"quantity"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.UserID,
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
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items WHERE (id = $1 AND user_id = $2)
`

type DeleteItemParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) DeleteItem(ctx context.Context, arg DeleteItemParams) error {
	_, err := q.db.ExecContext(ctx, deleteItem, arg.ID, arg.UserID)
	return err
}

const getItem = `-- name: GetItem :one
SELECT id, title, buying_price, selling_price, quantity, user_id, created_at FROM items
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
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getItemForUpdate = `-- name: GetItemForUpdate :one
SELECT id, title, buying_price, selling_price, quantity, user_id, created_at FROM items
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
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listItems = `-- name: ListItems :many
SELECT id, title, buying_price, selling_price, quantity, user_id, created_at FROM items
WHERE user_id = $1
ORDER BY created_at
`

func (q *Queries) ListItems(ctx context.Context, userID uuid.UUID) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems, userID)
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
			&i.UserID,
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
SET title = $2,
buying_price = $3,
selling_price = $4,
quantity = $5
WHERE (id = $1 AND user_id = $6)
RETURNING id, title, buying_price, selling_price, quantity, user_id, created_at
`

type UpdateItemParams struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	BuyingPrice  float32   `json:"buying_price"`
	SellingPrice float32   `json:"selling_price"`
	Quantity     int64     `json:"quantity"`
	UserID       uuid.UUID `json:"user_id"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.ID,
		arg.Title,
		arg.BuyingPrice,
		arg.SellingPrice,
		arg.Quantity,
		arg.UserID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.BuyingPrice,
		&i.SellingPrice,
		&i.Quantity,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
