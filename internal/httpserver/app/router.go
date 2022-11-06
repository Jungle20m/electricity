package app

import (
	"github.com/Jungle20m/electricity/component"
	httpTransport "github.com/Jungle20m/electricity/internal/module/order/http-transport"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, appCtx component.AppContext) {
	// Middleware
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger

	// Router
	h := handler.Group("/base")
	{
		h.GET("/something", httpTransport.GetSomething(appCtx))
	}
}
