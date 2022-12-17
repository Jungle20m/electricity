package http_transport

import (
	"github.com/Jungle20m/electricity/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	orderBusiness "github.com/Jungle20m/electricity/internal/modules/order/business"
	orderStorage "github.com/Jungle20m/electricity/internal/modules/order/storage"
)

// @Summary get order by id
// @Description
// @Tags orders
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} common.SuccessResponse{data=model.Order}
// @Failure 400 {object} common.ErrorResponse
// @Router /orders/{id} [get]
func GetOrderByID(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		storage := orderStorage.NewMysqlStorage(appCtx.GetMysql().DB)
		business := orderBusiness.NewGetOrderBusiness(storage, appCtx)

		order, err := business.GetOrderByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, err))
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(order))
	}
}
