package http_transport

import (
	"github.com/Jungle20m/electricity/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	productBusiness "github.com/Jungle20m/electricity/internal/modules/product/business"
	productStorage "github.com/Jungle20m/electricity/internal/modules/product/storage"
)

// @Summary get product by id
// @Description
// @Tags products
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} common.SuccessResponse{data=model.Product}
// @Failure 400 {object} common.ErrorResponse
// @Router /products/{id} [get]
func GetProductByID(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		storage := productStorage.NewMysqlStorage(appCtx.GetMysql().DB)
		business := productBusiness.NewGetProductBusiness(storage, appCtx)

		order, err := business.GetProductByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, err))
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(order))
	}
}
