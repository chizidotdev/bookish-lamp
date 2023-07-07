package datastruct

import "github.com/chizidotdev/copia/internal/repository"

type DashboardResponse struct {
	TotalItems       int64                           `json:"total_items"`
	LowStockItems    int64                           `json:"low_stock_items"`
	RecentSales      int64                           `json:"recent_sales"`
	PendingOrders    int64                           `json:"pending_orders"`
	SalesPerformance float64                         `json:"sales_performance"`
	PriceSoldByDate  []repository.PriceSoldByDateRow `json:"price_sold_by_date"`
	PriceSoldByWeek  []repository.PriceSoldByWeekRow `json:"price_sold_by_week"`
}
