// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateItem(ctx context.Context, arg CreateItemParams) (Item, error)
	CreateSale(ctx context.Context, arg CreateSaleParams) (Sale, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteItem(ctx context.Context, arg DeleteItemParams) error
	DeleteSale(ctx context.Context, arg DeleteSaleParams) error
	GetItem(ctx context.Context, id uuid.UUID) (Item, error)
	GetItemForUpdate(ctx context.Context, id uuid.UUID) (Item, error)
	GetSale(ctx context.Context, id uuid.UUID) (Sale, error)
	GetSaleForUpdate(ctx context.Context, id uuid.UUID) (Sale, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListItems(ctx context.Context, userID uuid.UUID) ([]Item, error)
	ListSales(ctx context.Context, itemID uuid.UUID) ([]Sale, error)
	ListUsers(ctx context.Context) ([]User, error)
	// UPDATE dashboard
	// SET items_to_ship = (
	//     SELECT COUNT(*) FROM items
	//     WHERE items.user_id = dashboard.user_id
	//     AND items.status = 'ready_to_ship'
	// );
	UpdateDashboardItemsToShip(ctx context.Context) error
	UpdateDashboardLowStockItems(ctx context.Context) error
	// UPDATE dashboard
	// SET notifications = (
	//   SELECT COUNT(*) FROM notifications
	//   WHERE notifications.user_id = dashboard.user_id
	// );
	UpdateDashboardNotifications(ctx context.Context) error
	UpdateDashboardPendingOrders(ctx context.Context) error
	UpdateDashboardSalesPerformance(ctx context.Context) error
	UpdateDashboardTotalItems(ctx context.Context) error
	UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error)
	UpdateItemQuantity(ctx context.Context, arg UpdateItemQuantityParams) (Item, error)
	UpdateSale(ctx context.Context, arg UpdateSaleParams) (Sale, error)
}

var _ Querier = (*Queries)(nil)
