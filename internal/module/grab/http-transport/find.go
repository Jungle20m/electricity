package http_transport

import (
	"github.com/Jungle20m/electricity/component"
	orderBusiness "github.com/Jungle20m/electricity/internal/module/grab/business"
	orderStorage "github.com/Jungle20m/electricity/internal/module/grab/storage"
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

//// todo represents data about a task in the todo list
//type todo struct {
//	ID   string `json:"id"`
//	Task string `json:"task"`
//}
//
//// message represents request response with a message
//type message struct {
//	Message string `json:"message"`
//}

// @Summary get a list order of customer
// @ID get-orders
// @Produce json
// @Param customer_code path string true "todo ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /grab-electric/orders/customer/{customer_code} [get]
func GetCustomerOrders(appCtx component.AppContext) gin.HandlerFunc {
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

func GetCustomerOrder(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "customer order",
		})
	}
}

func GetElectricianOrders(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "electrician orders",
		})
	}
}

func GetElectricianOrder(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//msql := appCtx.GetMysql()
		c.JSON(http.StatusOK, gin.H{
			"data": "electrician order",
		})
	}
}
