package api

import (
	"database/sql"
	"net/http"

	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createItemRequest struct {
	Title        string  `json:"title" binding:"required"`
	BuyingPrice  float32 `json:"buying_price" binding:"required"`
	SellingPrice float32 `json:"selling_price" binding:"required"`
	Quantity     int64   `json:"quantity" binding:"required"`
}

func (server *Server) createItem(ctx *gin.Context) {
	var req createItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(userJWT)

	arg := db.CreateItemParams{
		Title:        req.Title,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		UserID:       user.ID,
	}

	item, err := server.store.CreateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (server *Server) getItem(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	item, err := server.store.GetItem(ctx, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, item)
}

type listItemRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=15"`
}

func (server *Server) listItems(ctx *gin.Context) {
	var req listItemRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(userJWT)

	arg := db.ListItemsParams{
		UserID: user.ID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	items, err := server.store.ListItems(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, items)
}

type updateItemRequest struct {
	Title        string  `json:"title" binding:"required"`
	BuyingPrice  float32 `json:"buying_price" binding:"required"`
	SellingPrice float32 `json:"selling_price" binding:"required"`
	Quantity     int64   `json:"quantity" binding:"required,min=0"`
}

func (server *Server) updateItem(ctx *gin.Context) {
	var req updateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	idParam := ctx.Params.ByName("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(userJWT)

	arg := db.UpdateItemParams{
		ID:           itemID,
		Title:        req.Title,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Quantity:     req.Quantity,
		UserID:       user.ID,
	}

	item, err := server.store.UpdateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
	}

	ctx.JSON(http.StatusOK, item)
}

func (server *Server) deleteItem(ctx *gin.Context) {
	idParam := ctx.Params.ByName("id")
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	user := ctx.MustGet("user").(userJWT)

	arg := db.DeleteItemParams{
		ID:     itemID,
		UserID: user.ID,
	}

	err = server.store.DeleteItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "Item deleted")
}
