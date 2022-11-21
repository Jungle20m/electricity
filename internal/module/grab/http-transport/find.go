package http_transport

import (
	"fmt"
	"github.com/Jungle20m/electricity/component"
	orderBusiness "github.com/Jungle20m/electricity/internal/module/grab/business"
	orderStorage "github.com/Jungle20m/electricity/internal/module/grab/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomerOrders(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerCode := c.Param("customer_code")

		mysqlDB := appCtx.GetMysql().DB
		storage := orderStorage.NewStorage(mysqlDB)
		business := orderBusiness.NewGetOrderBusiness(appCtx, storage)

		business.GetCustomerOrders(c.Request.Context())
		fmt.Println(customerCode)

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
