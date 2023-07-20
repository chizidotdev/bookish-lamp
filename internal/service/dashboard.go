package service

import (
	"context"
	"fmt"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/repository/sqlx"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/google/uuid"
	"math"
)

type DashboardService interface {
	GetDashboard(ctx context.Context, userID uuid.UUID) (datastruct.DashboardResponse, error)
}

type dashboardService struct {
	Store *sqlx.Store
}

func NewDashboardService(store *sqlx.Store) DashboardService {
	return &dashboardService{
		Store: store,
	}
}

func (d *dashboardService) GetDashboard(ctx context.Context, userID uuid.UUID) (datastruct.DashboardResponse, error) {
	inventory, err := d.Store.GetInventoryStats(ctx, userID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get inventory stats: %w", err)
		return datastruct.DashboardResponse{}, errMsg
	}

	salesPerformance, err := d.getSalesPerformance(ctx, userID)
	if err != nil {
		return datastruct.DashboardResponse{}, err
	}

	priceSoldByDate, err := d.getPriceSoldByDate(ctx, userID)
	if err != nil {
		return datastruct.DashboardResponse{}, err
	}

	priceSoldByWeek, err := d.getPriceSoldByWeek(ctx, userID)
	if err != nil {
		return datastruct.DashboardResponse{}, err
	}

	dashboard := datastruct.DashboardResponse{
		TotalItems:       inventory.TotalItems,
		LowStockItems:    inventory.LowStockItems,
		RecentSales:      inventory.RecentSales,
		PendingOrders:    inventory.PendingOrders,
		SalesPerformance: salesPerformance,
		PriceSoldByDate:  priceSoldByDate,
		PriceSoldByWeek:  priceSoldByWeek,
	}

	return dashboard, nil
}

func (d *dashboardService) getSalesPerformance(ctx context.Context, userID uuid.UUID) (float64, error) {
	currWeekSale, err := d.Store.CurrentWeekSales(ctx, userID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get current week sales: %w", err)
		return 0, errMsg
	}

	lastWeekSales, err := d.Store.LastWeekSales(ctx, userID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get last week sales: %w", err)
		return 0, errMsg
	}

	var salesPerformance float64
	if lastWeekSales == 0 {
		if currWeekSale == 0 {
			salesPerformance = 0
		} else {
			salesPerformance = 100
		}
	} else {
		diff := utils.CalcPercentageDiff(lastWeekSales, currWeekSale)
		salesPerformance = math.Floor(diff)
	}

	return salesPerformance, nil
}

func (d *dashboardService) getPriceSoldByDate(ctx context.Context, userID uuid.UUID) ([]sqlx.PriceSoldByDateRow, error) {
	priceSoldByDate, err := d.Store.PriceSoldByDate(ctx, userID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get price sold by date: %w", err)
		return nil, errMsg
	}

	return priceSoldByDate, nil
}

func (d *dashboardService) getPriceSoldByWeek(ctx context.Context, userID uuid.UUID) ([]sqlx.PriceSoldByWeekRow, error) {
	priceSoldByWeek, err := d.Store.PriceSoldByWeek(ctx, userID)
	if err != nil {
		errMsg := fmt.Errorf("failed to get price sold by week: %w", err)
		return nil, errMsg
	}

	return priceSoldByWeek, nil
}
