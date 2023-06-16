package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type saleRequest struct {
	QuantitySold int64     `json:"quantity_sold"`
	SalePrice    float32   `json:"sale_price"`
	CustomerName string    `json:"customer_name"`
	SaleDate     time.Time `json:"sale_date"`
}

func (server *Server) createSale(ctx *gin.Context) {
	var req saleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	args := db.CreateSaleParams{
		ItemID:       uuid.MustParse(ctx.Param("id")),
		QuantitySold: req.QuantitySold,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		SaleDate:     req.SaleDate,
	}

	user := ctx.MustGet("user").(userJWT)
	itemArg := db.UpdateItemQuantityParams{
		ID:       args.ItemID,
		Quantity: -args.QuantitySold,
		UserID:   user.ID,
	}

	var sale db.Sale
	err := server.store.ExecTx(ctx, func(query *db.Queries) error {
		var err error
		sale, err = query.CreateSale(ctx, args)
		if err != nil {
			return err
		}
		_, err = query.UpdateItemQuantity(ctx, itemArg)
		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) updateSale(ctx *gin.Context) {
	var req saleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	itemID := uuid.MustParse(ctx.Param("id"))
	saleID := uuid.MustParse(ctx.Param("saleID"))

	initialSale, err := server.store.GetSale(ctx, db.GetSaleParams{
		ID:     saleID,
		ItemID: itemID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	quantityDiff := req.QuantitySold - initialSale.QuantitySold

	args := db.UpdateSaleParams{
		ID:           saleID,
		QuantitySold: req.QuantitySold,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		SaleDate:     req.SaleDate,
		ItemID:       itemID,
	}

	user := ctx.MustGet("user").(userJWT)
	itemArgs := db.UpdateItemQuantityParams{
		ID:       itemID,
		Quantity: -quantityDiff,
		UserID:   user.ID,
	}

	var sale db.Sale
	err = server.store.ExecTx(ctx, func(query *db.Queries) error {
		var err error
		sale, err = query.UpdateSale(ctx, args)
		if err != nil {
			return fmt.Errorf("error updating sale: %w", err)
		}
		_, err = query.UpdateItemQuantity(ctx, itemArgs)
		if err != nil {
			return fmt.Errorf("error updating item quantity: %w", err)
		}

		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) listSales(ctx *gin.Context) {
	ItemID := uuid.MustParse(ctx.Param("id"))

	sales, err := server.store.ListSales(ctx, ItemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sales)
}

func (server *Server) getSale(ctx *gin.Context) {
	itemID := uuid.MustParse(ctx.Param("id"))
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := server.store.GetSale(ctx, db.GetSaleParams{
		ID:     saleID,
		ItemID: itemID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (server *Server) deleteSale(ctx *gin.Context) {
	itemID := uuid.MustParse(ctx.Param("id"))
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := server.store.GetSale(ctx, db.GetSaleParams{
		ID:     saleID,
		ItemID: itemID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	args := db.DeleteSaleParams{
		ID:     saleID,
		ItemID: itemID,
	}

	user := ctx.MustGet("user").(userJWT)
	itemArgs := db.UpdateItemQuantityParams{
		ID:       itemID,
		Quantity: sale.QuantitySold,
		UserID:   user.ID,
	}

	err = server.store.ExecTx(ctx, func(query *db.Queries) error {
		var err error
		err = query.DeleteSale(ctx, args)
		if err != nil {
			return err
		}
		_, err = query.UpdateItemQuantity(ctx, itemArgs)
		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, "Sale deleted successfully")
}
