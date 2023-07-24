package repository

import (
	"context"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/google/uuid"
	"time"
)

func (s *Store) CreateSale(_ context.Context, arg datastruct.CreateSaleParams) (Sale, error) {
	sale := Sale{
		ItemID:       arg.ItemID,
		UserEmail:    arg.UserEmail,
		QuantitySold: arg.QuantitySold,
		SalePrice:    arg.SalePrice,
		CustomerName: arg.CustomerName,
		SaleDate:     arg.SaleDate,
	}
	result := s.DB.Create(&sale)
	return sale, result.Error
}

func (s *Store) DeleteSale(_ context.Context, arg datastruct.DeleteSaleParams) error {
	result := s.DB.Delete(&Sale{}, "id = ? AND user_email = ?", arg.ID, arg.UserEmail)
	return result.Error
}

func (s *Store) GetSale(_ context.Context, arg datastruct.GetSaleParams) (Sale, error) {
	var sale Sale
	result := s.DB.First(&sale, "id = ? AND user_email = ?", arg.ID, arg.UserEmail)
	return sale, result.Error
}

type ListSalesRow struct {
	ID           uuid.UUID `json:"id"`
	ItemID       uuid.UUID `json:"item_id"`
	UserID       uuid.UUID `json:"user_id"`
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	SaleDate     time.Time `json:"sale_date"`
	CustomerName string    `json:"customer_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Title        string    `json:"title"`
}

func (s *Store) ListSales(_ context.Context, arg datastruct.ListSalesParams) ([]ListSalesRow, error) {
	var sales []ListSalesRow
	result := s.DB.Table("sales s").
		Joins("JOIN items i ON s.item_id = i.id").
		Select("s.id, s.item_id, s.user_email, s.quantity_sold, s.sale_price, s.sale_date, s.customer_name, s.created_at, s.updated_at, i.title").
		Where("s.item_id = ? AND s.user_email = ?", arg.ItemID, arg.UserEmail).
		Order("s.sale_date DESC").
		Scan(&sales)

	return sales, result.Error
}

func (s *Store) ListSalesByUser(_ context.Context, userEmail string) ([]ListSalesRow, error) {
	var sales []ListSalesRow
	result := s.DB.Table("sales s").
		Joins("JOIN items i ON s.item_id = i.id").
		Select("s.id, s.item_id, s.user_email, s.quantity_sold, s.sale_price, s.sale_date, s.customer_name, s.created_at, s.updated_at, i.title").
		Where("s.user_email = ?", userEmail).
		Order("s.sale_date DESC").
		Scan(&sales)

	return sales, result.Error
}

func (s *Store) UpdateSale(_ context.Context, arg datastruct.UpdateSaleParams) (Sale, error) {
	var sale Sale
	if err := s.DB.First(&sale, "id = ? AND user_email = ?", arg.ID, arg.UserEmail).Error; err != nil {
		return sale, err
	}
	sale.QuantitySold = arg.QuantitySold
	sale.SalePrice = arg.SalePrice
	sale.CustomerName = arg.CustomerName
	sale.SaleDate = arg.SaleDate

	err := s.DB.Save(&sale).Error
	return sale, err
}
