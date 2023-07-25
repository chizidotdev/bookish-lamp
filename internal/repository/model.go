package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(db *gorm.DB) error {
	base.ID = uuid.New()

	return nil
}

// type UserProfile struct {
// 	Email     string `gorm:"primaryKey"`
// 	FirstName string
// 	LastName  string
// 	Items     []Item `gorm:"foreignKey:UserEmail;"`
// 	Sales     []Sale `gorm:"foreignKey:UserEmail;"`
// }

type Item struct {
	Base
	Title        string  `gorm:"not null" json:"title"`
	BuyingPrice  float32 `gorm:"not null" json:"buying_price"`
	SellingPrice float32 `gorm:"not null" json:"selling_price"`
	Quantity     int64   `gorm:"not null" json:"quantity"`
	Sales        []Sale  `gorm:"foreignKey:ItemID" json:"sales"`
	UserEmail    string  `gorm:"not null" json:"user_email"`
}

type Sale struct {
	Base
	QuantitySold int64     `gorm:"not null" json:"quantity_sold"`
	SalePrice    float32   `gorm:"not null" json:"sale_price"`
	SaleDate     time.Time `gorm:"not null" json:"sale_date"`
	CustomerName string    `gorm:"not null" json:"customer_name"`
	ItemID       uuid.UUID `gorm:"not null" json:"item_id"`
	UserEmail    string    `gorm:"not null" json:"user_email"`
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
