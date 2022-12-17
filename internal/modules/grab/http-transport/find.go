package http_transport

import (
	"github.com/Jungle20m/electricity/common"
	orderBusiness "github.com/Jungle20m/electricity/internal/modules/grab/business"
	orderStorage "github.com/Jungle20m/electricity/internal/modules/grab/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response represents request response with a response
type Response struct {
	Code         int         `json:"code"`
	Status       string      `json:"status"`
	ErrorCode    string      `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func GetCustomerOrders(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := appCtx.GetLogger()

		customerCode := c.Param("customer_code")

		mysqlDB := appCtx.GetMysql().DB
		storage := orderStorage.NewStorage(mysqlDB)
		business := orderBusiness.NewGetOrderBusiness(appCtx, storage)

		business.GetCustomerOrders(c.Request.Context())

		logger.Info(customerCode)

		c.JSON(http.StatusOK, gin.H{
			"data": "customer orders",
		})
	}
}

func GetCustomerOrder(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "customer order",
		})
	}
}

func GetElectricianOrders(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "electrician orders",
		})
	}
}

func GetElectricianOrder(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "electrician order",
		})
	}
}
