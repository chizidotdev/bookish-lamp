package service

import (
	"context"
	"github.com/chizidotdev/copia/internal/datastruct"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/google/uuid"
)

type ItemService interface {
	ListItems(ctx context.Context, userID string) ([]repository.Item, error)
	CreateItem(ctx context.Context, req datastruct.CreateItemParams) (repository.Item, error)
	UpdateItem(ctx context.Context, req datastruct.UpdateItemParams) (repository.Item, error)
	GetItemByID(ctx context.Context, itemID uuid.UUID) (repository.Item, error)
	DeleteItem(ctx context.Context, req datastruct.DeleteItemParams) error
}

type itemService struct {
	Store *repository.Store
}

func NewItemService(store *repository.Store) ItemService {
	return &itemService{
		Store: store,
	}
}

func (i *itemService) CreateItem(ctx context.Context, req datastruct.CreateItemParams) (repository.Item, error) {
	newItem, err := i.Store.CreateItem(ctx, req)
	if err != nil {
		return repository.Item{}, err
	}

	return newItem, nil
}

func (i *itemService) ListItems(ctx context.Context, userEmail string) ([]repository.Item, error) {
	items, err := i.Store.ListItems(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemService) UpdateItem(ctx context.Context, req datastruct.UpdateItemParams) (repository.Item, error) {
	item, err := i.Store.UpdateItem(ctx, req)
	if err != nil {
		return repository.Item{}, err
	}

	return item, nil
}

func (i *itemService) GetItemByID(ctx context.Context, itemID uuid.UUID) (repository.Item, error) {
	item, err := i.Store.GetItem(ctx, itemID)
	if err != nil {
		return repository.Item{}, err
	}

	return item, nil
}

func (i *itemService) DeleteItem(ctx context.Context, req datastruct.DeleteItemParams) error {
	err := i.Store.DeleteItem(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
