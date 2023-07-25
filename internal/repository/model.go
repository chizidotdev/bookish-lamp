package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID      `gorm:"primaryKey;column:id;"`
	CreatedAt time.Time      `gorm:"column:created_at;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;"`
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
	Title        string  `gorm:"not null;column:title;"`
	BuyingPrice  float32 `gorm:"not null;column:buying_price;"`
	SellingPrice float32 `gorm:"not null;column:selling_price;"`
	Quantity     int64   `gorm:"not null;column:quantity;"`
	Sales        []Sale  `gorm:"foreignKey:ItemID;"`
	UserEmail    string  `gorm:"not null;column:user_email;"`
}

type Sale struct {
	Base
	QuantitySold int64     `gorm:"not null;column:quantity_sold;"`
	SalePrice    float32   `gorm:"not null;column:sale_price;"`
	SaleDate     time.Time `gorm:"not null;column:sale_date;"`
	CustomerName string    `gorm:"not null;column:customer_name;"`
	ItemID       uuid.UUID `gorm:"not null;column:item_id;"`
	UserEmail    string    `gorm:"not null;column:user_email;"`
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
