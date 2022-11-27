package httpserver

import (
	"github.com/Jungle20m/electricity/component"
	grabHttpTransport "github.com/Jungle20m/electricity/internal/module/grab/http-transport"
	"github.com/gin-gonic/gin"
)

func NewHandler(appCtx component.AppContext) *gin.Engine {
	// Init handler
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

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

	return handler
}
