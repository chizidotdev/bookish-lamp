package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/chizidotdev/copia/db/mock"
	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetItem(t *testing.T) {
	item := randomItem()

	controller := gomock.NewController(t)
	defer controller.Finish()

	store := mockdb.NewMockStore(controller)

	store.EXPECT().
		CreateItem(gomock.Any(), gomock.Eq(item.ID)).
		Times(1).
		Return(item, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/items/%s", item.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomItem() db.Item {
	return db.Item{
		ID:           uuid.New(),
		Title:        utils.RandomString(6),
		BuyingPrice:  utils.RandomMoney(),
		SellingPrice: utils.RandomMoney(),
		Quantity:     utils.RandomInt(1, 10),
	}
}
