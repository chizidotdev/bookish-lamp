package dto

import "time"

type CreateSaleRequest struct {
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	CustomerName string    `json:"customer_name"`
	SaleDate     time.Time `json:"sale_date"`
}

type UpdateSaleRequest struct {
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	CustomerName string    `json:"customer_name"`
	SaleDate     time.Time `json:"sale_date"`
}
