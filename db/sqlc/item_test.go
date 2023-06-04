package db

import (
	"context"
	"testing"
	"time"

	"github.com/chizidotdev/copia/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomItem(t *testing.T, userId uuid.UUID) Item {
	arg := CreateItemParams{
		UserID:       userId,
		Title:        utils.RandomTitle(),
		BuyingPrice:  utils.RandomMoney(),
		SellingPrice: utils.RandomMoney(),
		Quantity:     utils.RandomQuantity(),
	}

	item, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Title, item.Title)
	require.Equal(t, arg.BuyingPrice, item.BuyingPrice)
	require.Equal(t, arg.SellingPrice, item.SellingPrice)
	require.Equal(t, arg.Quantity, item.Quantity)

	require.NotZero(t, item.CreatedAt)

	return item
}

func TestCreateItem(t *testing.T) {
	user := createRandomUser(t)
	createRandomItem(t, user.ID)
}

func TestGetItem(t *testing.T) {
	user := createRandomUser(t)
	createdItem := createRandomItem(t, user.ID)
	item, err := testQueries.GetItem(context.Background(), createdItem.ID)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, createdItem.ID, item.ID)
	require.Equal(t, createdItem.Title, item.Title)
	require.Equal(t, createdItem.BuyingPrice, item.BuyingPrice)
	require.Equal(t, createdItem.SellingPrice, item.SellingPrice)
	require.Equal(t, createdItem.Quantity, item.Quantity)

	require.WithinDuration(t, createdItem.CreatedAt, item.CreatedAt, time.Second)
}

func TestUpdateItem(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	arg := UpdateItemParams{
		ID:       item.ID,
		Quantity: utils.RandomQuantity(),
	}

	updatedItem, err := testQueries.UpdateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedItem)

	require.Equal(t, arg.ID, updatedItem.ID)
	require.Equal(t, arg.Quantity, updatedItem.Quantity)
}

func TestDeleteItem(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	err := testQueries.DeleteItem(context.Background(), item.ID)
	require.NoError(t, err)

	deletedItem, err := testQueries.GetItem(context.Background(), item.ID)
	require.Error(t, err)
	require.Empty(t, deletedItem)
}

func TestListItems(t *testing.T) {
	user := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomItem(t, user.ID)
	}

	arg := ListItemsParams{
		UserID: user.ID,
		Limit:  4,
		Offset: 0,
	}

	items, err := testQueries.ListItems(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, items, 4)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}
