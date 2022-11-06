package app

import (
	"github.com/Jungle20m/electricity/component"
	"github.com/Jungle20m/electricity/internal/httpserver/module/order/transportation"
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
		h.GET("/something", transportation.GetSomething(appCtx))
	}
}
