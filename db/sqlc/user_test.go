package db

import (
	"context"
	"testing"
	"time"

	"github.com/chizidotdev/copia/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.Password)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func GetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), createdUser.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Email, user.Email)
	require.Equal(t, createdUser.Password, user.Password)

	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}
