package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name         string  `json:"name"`
	BuyingPrice  float64 `json:"buying_price"`
	SellingPrice float64 `json:"selling_price"`
	Quantity     int     `json:"quantity"`
}

type Sale struct {
	gorm.Model
	ItemID    string  `json:"item_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"created_at"`
}
