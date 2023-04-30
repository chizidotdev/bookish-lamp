package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string    `json:"name"`
	BuyingPrice  float64   `json:"buying_price"`
	SellingPrice float64   `json:"selling_price"`
	Quantity     int       `json:"quantity"`
	Sales        []Sale    `json:"sales" gorm:"foreignKey:ItemID"`
}

type Sale struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ItemID       uuid.UUID `gorm:"type:uuid;index"`
	QuantitySold int       `json:"quantity_sold"`
	UnitPrice    float64   `json:"unit_price"`
	SaleDate     time.Time `json:"sale_date"`
}
