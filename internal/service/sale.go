package service

import (
	"context"
	"fmt"
	"github.com/chizidotdev/copia/internal/repository/sqlx"
	"github.com/google/uuid"
)

type SaleService interface {
	ListSalesByUser(ctx context.Context, userID uuid.UUID) ([]sqlx.ListSalesByUserIdRow, error)
	ListSalesByItem(ctx context.Context, req sqlx.ListSalesParams) ([]sqlx.ListSalesRow, error)
	CreateSale(ctx context.Context, req sqlx.CreateSaleParams) (sqlx.Sale, error)
	UpdateSale(ctx context.Context, req sqlx.UpdateSaleParams) (sqlx.Sale, error)
	GetSaleByID(ctx context.Context, params sqlx.GetSaleParams) (sqlx.Sale, error)
	DeleteSale(ctx context.Context, req sqlx.DeleteSaleParams) error
}

type saleService struct {
	Store *sqlx.Store
}

func NewSaleService(store *sqlx.Store) SaleService {
	return &saleService{
		Store: store,
	}
}

func (s *saleService) ListSalesByUser(ctx context.Context, userID uuid.UUID) ([]sqlx.ListSalesByUserIdRow, error) {
	sales, err := s.Store.ListSalesByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *saleService) ListSalesByItem(ctx context.Context, req sqlx.ListSalesParams) ([]sqlx.ListSalesRow, error) {
	_, err := s.Store.GetItem(ctx, req.ItemID)
	if err != nil {
		return nil, err
	}

	sales, err := s.Store.ListSales(ctx, req)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *saleService) CreateSale(ctx context.Context, req sqlx.CreateSaleParams) (sqlx.Sale, error) {
	var sale sqlx.Sale
	err := s.Store.ExecTx(ctx, func(query *sqlx.Queries) error {
		var err error
		sale, err = query.CreateSale(ctx, req)
		if err != nil {
			return err
		}
		_, err = query.UpdateItemQuantity(ctx, sqlx.UpdateItemQuantityParams{
			ID:       req.ItemID,
			Quantity: -req.QuantitySold,
			UserID:   req.UserID,
		})
		return err
	})
	if err != nil {
		return sqlx.Sale{}, err
	}

	return sale, nil
}

func (s *saleService) UpdateSale(ctx context.Context, req sqlx.UpdateSaleParams) (sqlx.Sale, error) {
	initialSale, err := s.GetSaleByID(ctx, sqlx.GetSaleParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		return sqlx.Sale{}, errMessage
	}

	quantityDiff := req.QuantitySold - initialSale.QuantitySold

	args := sqlx.UpdateSaleParams{
		ID:           req.ID,
		QuantitySold: req.QuantitySold,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		SaleDate:     req.SaleDate,
		UserID:       req.UserID,
	}

	itemArgs := sqlx.UpdateItemQuantityParams{
		ID:       initialSale.ItemID,
		Quantity: -quantityDiff,
		UserID:   req.UserID,
	}

	var sale sqlx.Sale
	err = s.Store.ExecTx(ctx, func(query *sqlx.Queries) error {
		var err error
		sale, err = query.UpdateSale(ctx, args)
		if err != nil {
			return fmt.Errorf("error updating sale: %w", err)
		}
		_, err = query.UpdateItemQuantity(ctx, itemArgs)
		if err != nil {
			return fmt.Errorf("error updating item quantity: %w", err)
		}
		return err
	})
	if err != nil {
		return sqlx.Sale{}, err
	}

	return sale, nil
}

func (s *saleService) GetSaleByID(ctx context.Context, req sqlx.GetSaleParams) (sqlx.Sale, error) {
	sale, err := s.Store.GetSale(ctx, req)
	if err != nil {
		return sqlx.Sale{}, err
	}

	return sale, nil
}

func (s *saleService) DeleteSale(ctx context.Context, req sqlx.DeleteSaleParams) error {
	sale, err := s.GetSaleByID(ctx, sqlx.GetSaleParams{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return fmt.Errorf("sale not found: %v", err)
	}

	err = s.Store.ExecTx(ctx, func(query *sqlx.Queries) error {
		err := query.DeleteSale(ctx, sqlx.DeleteSaleParams{
			ID:     req.ID,
			UserID: req.UserID,
		})
		if err != nil {
			return fmt.Errorf("an error occurred while deleting sale: %v", err)
		}

		_, err = query.UpdateItemQuantity(ctx, sqlx.UpdateItemQuantityParams{
			ID:       sale.ItemID,
			Quantity: sale.QuantitySold,
			UserID:   sale.UserID,
		})
		if err != nil {
			return fmt.Errorf("an error occurred while updating sale item quantity: %v", err)
		}

		return err
	})
	if err != nil {
		return err
	}

	return nil
}
