package datastruct

import (
	"github.com/google/uuid"
	"time"
)

type CreateSaleParams struct {
	ItemID       uuid.UUID `json:"item_id"`
	UserEmail    string    `json:"user_email"`
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	CustomerName string    `json:"customer_name"`
	SaleDate     time.Time `json:"sale_date"`
}

type DeleteSaleParams struct {
	ID        uuid.UUID `json:"id"`
	UserEmail string    `json:"user_email"`
}

type GetSaleParams struct {
	ID        uuid.UUID `json:"id"`
	UserEmail string    `json:"user_email"`
}
type ListSalesParams struct {
	ItemID    uuid.UUID `json:"item_id"`
	UserEmail string    `json:"user_email"`
}

type UpdateSaleParams struct {
	ID           uuid.UUID `json:"id"`
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	CustomerName string    `json:"customer_name"`
	SaleDate     time.Time `json:"sale_date"`
	UserEmail    string    `json:"user_email"`
}
