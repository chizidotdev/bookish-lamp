package api

import (
	"database/sql"
	"net/http"

	db "github.com/chizidotdev/copia/db/sqlc"
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateItemParams(req)
	item, err := server.store.CreateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, item)
}

type getItemRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

func (server *Server) getItem(ctx *gin.Context) {
	var req getItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := server.store.GetItem(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListItemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	items, err := server.store.ListItems(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, items)
}

type updateItemRequest struct {
	ID           uuid.UUID `uri:"id" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	BuyingPrice  float32   `json:"buying_price" binding:"required"`
	SellingPrice float32   `json:"selling_price" binding:"required"`
	Quantity     int64     `json:"quantity" binding:"required,min=0"`
}

func (server *Server) updateItem(ctx *gin.Context) {
	var req updateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateItemParams(req)
	item, err := server.store.UpdateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, item)
}

type deleteItemRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

func (server *Server) deleteItem(ctx *gin.Context) {
	var req deleteItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteItem(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Item deleted")
}
