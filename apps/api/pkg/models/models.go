package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID           string  `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	BuyingPrice  float64 `json:"buying_price"`
	SellingPrice float64 `json:"selling_price"`
}

type Sale struct {
	gorm.Model
	ID        string  `json:"id" gorm:"primaryKey"`
	ItemID    string  `json:"item_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"created_at"`
}
