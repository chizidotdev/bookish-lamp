package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(db *gorm.DB) error {
	base.ID = uuid.New()

	return nil
}

type UserProfile struct {
	Email     string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Items     []Item `gorm:"foreignKey:UserEmail;"`
	Sales     []Sale `gorm:"foreignKey:UserEmail;"`
}

type Item struct {
	Base
	Title        string  `gorm:"not null;"`
	BuyingPrice  float32 `gorm:"not null;"`
	SellingPrice float32 `gorm:"not null;"`
	Quantity     int64   `gorm:"not null;"`
	Sales        []Sale
	UserEmail    string
}

type Sale struct {
	Base
	QuantitySold int64     `gorm:"not null;"`
	SalePrice    float32   `gorm:"not null;"`
	SaleDate     time.Time `gorm:"not null;"`
	CustomerName string    `gorm:"not null;"`
	ItemID       uuid.UUID
	UserEmail    string
}

//type Order struct {
//	Base
//	OrderDate   time.Time `json:"order_date"`
//	TotalAmount float32   `json:"total_amount"`
//	Status      string    `json:"status"`
//	UserEmail   string
//}

//type OrderItem struct {
//	Base
//	OrderID   uuid.UUID `json:"order_id"`
//	ItemID    uuid.UUID `json:"item_id"`
//	Quantity  int64     `json:"quantity"`
//	Price     float32   `json:"price"`
//}

//type Customer struct {
//	Base
//	OrderID   uuid.UUID `json:"order_id"`
//	Name      string    `json:"name"`
//	Email     string    `json:"email"`
//	Phone     string    `json:"phone"`
//	Address   string    `json:"address"`
//}
