package repository

import (
	"context"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/google/uuid"
)

func (s *Store) CreateItem(_ context.Context, arg datastruct.CreateItemParams) (Item, error) {
	item := Item{
		UserEmail:    arg.UserEmail,
		Title:        arg.Title,
		BuyingPrice:  arg.BuyingPrice,
		SellingPrice: arg.SellingPrice,
		Quantity:     arg.Quantity,
	}
	result := s.DB.Create(&item)
	return item, result.Error
}

func (s *Store) DeleteItem(_ context.Context, arg datastruct.DeleteItemParams) error {
	result := s.DB.Delete(&Item{}, "id = ? AND user_email = ?", arg.ID, arg.UserEmail)
	return result.Error
}

func (s *Store) GetItem(_ context.Context, id uuid.UUID) (Item, error) {
	var item Item
	result := s.DB.First(&item, "id = ?", id)
	return item, result.Error
}

func (s *Store) ListItems(_ context.Context, userEmail string) ([]Item, error) {
	var items []Item
	result := s.DB.Find(&items, "user_email = ?", userEmail)
	return items, result.Error
}

func (s *Store) UpdateItem(_ context.Context, arg datastruct.UpdateItemParams) (Item, error) {
	var item Item
	if err := s.DB.First(&item, "id = ? AND user_email = ?", arg.ID, arg.UserEmail).Error; err != nil {
		return item, err
	}
	item.UserEmail = arg.UserEmail
	item.Title = arg.Title
	item.BuyingPrice = arg.BuyingPrice
	item.SellingPrice = arg.SellingPrice
	item.Quantity = arg.Quantity

	err := s.DB.Save(&item).Error
	return item, err
}

func (s *Store) UpdateItemQuantity(_ context.Context, arg datastruct.UpdateItemQuantityParams) (Item, error) {
	var item Item
	if err := s.DB.First(&item, "id = ? AND user_email = ?", arg.ID, arg.UserEmail).Error; err != nil {
		return item, err
	}

	item.Quantity = arg.Quantity
	err := s.DB.Save(&item).Error

	return item, err
}
