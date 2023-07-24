package repository

import (
	"context"
	"fmt"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Queries interface {
	CreateItem(ctx context.Context, arg datastruct.CreateItemParams) (Item, error)
	DeleteItem(ctx context.Context, arg datastruct.DeleteItemParams) error
	GetItem(ctx context.Context, id uuid.UUID) (Item, error)
	ListItems(ctx context.Context, userEmail string) ([]Item, error)
	UpdateItem(ctx context.Context, arg datastruct.UpdateItemParams) (Item, error)
	UpdateItemQuantity(ctx context.Context, arg datastruct.UpdateItemQuantityParams) (Item, error)

	CreateSale(ctx context.Context, arg datastruct.CreateSaleParams) (Sale, error)
	DeleteSale(ctx context.Context, arg datastruct.DeleteSaleParams) error
	GetSale(ctx context.Context, arg datastruct.GetSaleParams) (Sale, error)
	ListSales(ctx context.Context, arg datastruct.ListSalesParams) ([]ListSalesRow, error)
	ListSalesByUser(ctx context.Context, userEmail string) ([]ListSalesRow, error)
	UpdateSale(ctx context.Context, arg datastruct.UpdateSaleParams) (Sale, error)

	GetInventoryStats(ctx context.Context, userEmail string) (GetInventoryStatsRow, error)
	PriceSoldByDate(ctx context.Context, userEmail string) ([]datastruct.PriceSoldByDateRow, error)
	PriceSoldByWeek(ctx context.Context, userEmail string) ([]datastruct.PriceSoldByWeekRow, error)
	CurrentWeekSales(ctx context.Context, userEmail string) (int32, error)
	LastWeekSales(ctx context.Context, userEmail string) (int32, error)

	//CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	//GetUser(ctx context.Context, email string) (User, error)
	//ListUsers(ctx context.Context) ([]User, error)
}

var _ Queries = (*Store)(nil)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	err := db.AutoMigrate(&UserProfile{}, &Item{}, &Sale{})
	if err != nil {
		log.Panic("Cannot migrate db:", err)
	}

	return &Store{
		DB: db,
	}
}

func (s *Store) WithTx(tx *gorm.DB) *Store {
	return &Store{
		DB: tx,
	}
}

// ExecTx executes a function within a transaction.
func (s *Store) ExecTx(ctx context.Context, fn func(*Store) error) error {
	tx := s.DB.Begin()

	qtx := s.WithTx(tx)

	err := fn(qtx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit().Error
}
