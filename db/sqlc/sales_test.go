package db

import (
	"context"
	"testing"
	"time"

	"github.com/chizidotdev/copia/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSale(t *testing.T, itemID uuid.UUID) Sale {
	arg := CreateSaleParams{
		ItemID:       itemID,
		QuantitySold: utils.RandomQuantity(),
		SalePrice:    utils.RandomMoney(),
		CustomerName: utils.RandomTitle(),
		SaleDate:     time.Now(),
	}

	sale, err := testQueries.CreateSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sale)

	require.Equal(t, arg.ItemID, sale.ItemID)
	require.Equal(t, arg.QuantitySold, sale.QuantitySold)
	require.Equal(t, arg.SalePrice, sale.SalePrice)
	require.Equal(t, arg.CustomerName, sale.CustomerName)

	require.NotZero(t, sale.CreatedAt)
	require.NotZero(t, sale.UpdatedAt)

	return sale
}

func TestCreateSale(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	createRandomSale(t, item.ID)
}

func TestGetSale(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	createdSale := createRandomSale(t, item.ID)
	args := GetSaleParams{
		ID:     createdSale.ID,
		UserID: user.ID,
	}
	sale, err := testQueries.GetSale(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, sale)

	require.Equal(t, createdSale.ID, sale.ID)
	require.Equal(t, createdSale.ItemID, sale.ItemID)
	require.Equal(t, createdSale.QuantitySold, sale.QuantitySold)
	require.Equal(t, createdSale.SalePrice, sale.SalePrice)
	require.Equal(t, createdSale.CustomerName, sale.CustomerName)

	require.WithinDuration(t, createdSale.CreatedAt, sale.CreatedAt, time.Second)
	require.WithinDuration(t, createdSale.UpdatedAt, sale.UpdatedAt, time.Second)
}

func TestUpdateSale(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	sale := createRandomSale(t, item.ID)
	arg := UpdateSaleParams{
		ID:           sale.ID,
		QuantitySold: utils.RandomQuantity(),
		SalePrice:    utils.RandomMoney(),
		CustomerName: utils.RandomTitle(),
		SaleDate:     time.Now(),
		UserID:       user.ID,
	}

	updatedSale, err := testQueries.UpdateSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSale)

	require.Equal(t, arg.ID, updatedSale.ID)
	require.Equal(t, arg.QuantitySold, updatedSale.QuantitySold)
	require.Equal(t, arg.SalePrice, updatedSale.SalePrice)
	require.Equal(t, arg.CustomerName, updatedSale.CustomerName)

	require.WithinDuration(t, arg.SaleDate, updatedSale.SaleDate, time.Second)
	require.WithinDuration(t, sale.CreatedAt, updatedSale.CreatedAt, time.Second)
	require.WithinDuration(t, updatedSale.UpdatedAt, updatedSale.UpdatedAt, time.Second)
}

func TestDeleteSale(t *testing.T) {
	user := createRandomUser(t)
	item := createRandomItem(t, user.ID)
	sale := createRandomSale(t, item.ID)
	err := testQueries.DeleteSale(context.Background(), DeleteSaleParams{
		ID:     sale.ID,
		UserID: user.ID,
	})
	require.NoError(t, err)

	args := GetSaleParams{
		ID:     sale.ID,
		UserID: user.ID,
	}
	deletedSale, err := testQueries.GetSale(context.Background(), args)
	require.Error(t, err)
	require.Empty(t, deletedSale)
}
