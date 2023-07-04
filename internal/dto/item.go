package dto

type CreateItemRequest struct {
	Title        string  `json:"title" binding:"required"`
	BuyingPrice  float32 `json:"buying_price" binding:"required"`
	SellingPrice float32 `json:"selling_price" binding:"required"`
	Quantity     int64   `json:"quantity" binding:"required"`
}

type UpdateItemRequest struct {
	Title        string  `json:"title" binding:"required"`
	BuyingPrice  float32 `json:"buying_price" binding:"required"`
	SellingPrice float32 `json:"selling_price" binding:"required"`
	Quantity     int64   `json:"quantity" binding:"required,min=0"`
}
