package sale

import (
	"fmt"
	"github.com/chizidotdev/copia/internal/dto"
	"github.com/chizidotdev/copia/internal/repository"
	"github.com/chizidotdev/copia/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *saleService) GetSaleByID(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserJWT)
	saleID := uuid.MustParse(ctx.Param("saleID"))

	sale, err := s.Store.GetSale(ctx, repository.GetSaleParams{
		ID:     saleID,
		UserID: user.ID,
	})
	if err != nil {
		errMessage := fmt.Errorf("sale not found: %v", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(errMessage.Error()))
		return
	}

	ctx.JSON(http.StatusOK, sale)
}
