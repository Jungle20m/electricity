package app

import (
	"github.com/Jungle20m/electricity/component"
	grabHttpTransport "github.com/Jungle20m/electricity/internal/module/grab/http-transport"
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
		order.GET("/customer/:customer_code", grabHttpTransport.GetCustomerOrders(appCtx))
		order.GET("/:order_code/customer", grabHttpTransport.GetCustomerOrder(appCtx))

		// Services
		service := v2.Group("/services")
		service.POST("/", grabHttpTransport.CreateService(appCtx))
	}
}
