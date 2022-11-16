package app

import (
	"github.com/Jungle20m/electricity/component"
	orderHttpTransport "github.com/Jungle20m/electricity/internal/module/order/http-transport"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, appCtx component.AppContext) {
	// Middleware
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger

	// Router
	v2 := handler.Group("/grab-electric")
	{
		// Orders
		order := v2.Group("/orders")
		order.GET("/customer/:customer_code", orderHttpTransport.GetCustomerOrders(appCtx))
		order.GET("/:order_code/customer", orderHttpTransport.GetCustomerOrder(appCtx))
	}
}
