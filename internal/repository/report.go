package repository

import (
	"context"
	"github.com/chizidotdev/copia/internal/datastruct"
	"gorm.io/gorm/clause"
)

type GetInventoryStatsRow struct {
	TotalItems    int64 `json:"total_items"`
	LowStockItems int64 `json:"low_stock_items"`
	RecentSales   int64 `json:"recent_sales"`
	//PendingOrders int64 `json:"pending_orders"`
}

func (s *Store) GetInventoryStats(_ context.Context, userEmail string) (GetInventoryStatsRow, error) {
	var result GetInventoryStatsRow
	err := s.DB.Model(&Item{}).
		Select("(SELECT COUNT(*) FROM items WHERE items.user_email = ?) AS total_items", userEmail).
		Select("(SELECT COUNT(*) FROM items WHERE items.user_email = ? AND items.quantity <= 5) AS low_stock_items", userEmail).
		Select("(SELECT COUNT(*) FROM sales WHERE sales.user_email = ? LIMIT 10) AS recent_sales", userEmail).
		//Select("(SELECT COUNT(*) FROM orders WHERE orders.user_email = ? AND orders.status = 'Pending') AS pending_orders", userEmail).
		Scan(&result).Error

	return result, err
}

func (s *Store) PriceSoldByWeek(_ context.Context, userEmail string) ([]datastruct.PriceSoldByWeekRow, error) {
	var results []datastruct.PriceSoldByWeekRow
	err := s.DB.Model(&Sale{}).
		Select("DATE_TRUNC('week', sale_date) AS date, SUM(sale_price) AS total_sale_price").
		Where("user_email = ?", userEmail).
		Group("DATE_TRUNC('week', sale_date)").
		Order(clause.Expr{SQL: "DATE_TRUNC('week', sale_date)"}).
		Scan(&results).Error

	return results, err
}

func (s *Store) PriceSoldByDate(_ context.Context, userEmail string) ([]datastruct.PriceSoldByDateRow, error) {
	var results []datastruct.PriceSoldByDateRow
	err := s.DB.Model(&Sale{}).
		Select("DATE_TRUNC('day', sale_date) AS date, SUM(sale_price) AS total_sale_price").
		Where("user_email = ?", userEmail).
		Group("DATE_TRUNC('day', sale_date)").
		Order(clause.Expr{SQL: "DATE_TRUNC('day', sale_date)"}).
		Scan(&results).Error

	return results, err
}

//const currentWeekSales = `-- name: CurrentWeekSales :one
//SELECT CAST(COALESCE(SUM(quantity_sold), 0) AS INTEGER) AS total_quantity_sold
//FROM sales
//WHERE sale_date >= DATE_TRUNC('week', CURRENT_DATE)
//  AND sale_date < DATE_TRUNC('week', CURRENT_DATE) + INTERVAL '1 week'
//  AND user_id = $1
//`

func (s *Store) CurrentWeekSales(_ context.Context, userEmail string) (int32, error) {
	return 0, nil
}

//const lastWeekSales = `-- name: LastWeekSales :one
//SELECT CAST(COALESCE(SUM(quantity_sold), 0) AS INTEGER) AS total_quantity_sold
//FROM sales
//WHERE sale_date >= DATE_TRUNC('week', CURRENT_DATE) - INTERVAL '1 week'
//  AND sale_date
//    < DATE_TRUNC('week'
//    , CURRENT_DATE)
//  AND user_id = $1
//`

func (s *Store) LastWeekSales(_ context.Context, userEmail string) (int32, error) {
	return 0, nil
}
