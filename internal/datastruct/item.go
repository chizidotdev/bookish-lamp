package datastruct

import "github.com/google/uuid"

type CreateItemParams struct {
	UserEmail    string  `json:"user_email"`
	Title        string  `json:"title"`
	BuyingPrice  float32 `json:"buying_price"`
	SellingPrice float32 `json:"selling_price"`
	Quantity     int64   `json:"quantity"`
}

type DeleteItemParams struct {
	ID        uuid.UUID `json:"id"`
	UserEmail string    `json:"user_email"`
}

type UpdateItemParams struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	BuyingPrice  float32   `json:"buying_price"`
	SellingPrice float32   `json:"selling_price"`
	Quantity     int64     `json:"quantity"`
	UserEmail    string    `json:"user_email"`
}

type UpdateItemQuantityParams struct {
	ID        uuid.UUID `json:"id"`
	Quantity  int64     `json:"quantity"`
	UserEmail string    `json:"user_email"`
}
