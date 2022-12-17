package httpserver

import (
	"github.com/Jungle20m/electricity/common"
	_ "github.com/Jungle20m/electricity/docs"
	"github.com/Jungle20m/electricity/internal/httpserver/middleware"
	orderTransport "github.com/Jungle20m/electricity/internal/modules/order/http-transport"
	productTransport "github.com/Jungle20m/electricity/internal/modules/product/http-transport"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
	"path/filepath"
)

// @host localhost:8000
// @BasePath /example/v1
// @query.collection.format multi
func NewHandler(appCtx common.AppContext) *gin.Engine {
	config := appCtx.GetConfig()

	// Init handler
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	// Recovery middleware
	gin.DisableConsoleColor()
	f, _ := os.Create(filepath.Join(config.Log.Folder, config.Log.ApiLogFile))
	gin.DefaultWriter = io.MultiWriter(f)
	handler.Use(gin.Recovery())

	// Logger middleware
	handler.Use(gin.Logger())

	// Tracing middleware
	handler.Use(middleware.RequestID(nil))

	// Router
	v1 := handler.Group("/example/v1")
	{
		// Order
		order := v1.Group("/orders")
		order.GET("/:id", orderTransport.GetOrderByID(appCtx))
		//order.GET("/:order_code/customer", orderTransport.GetOrderByID(appCtx))

		// Product
		product := v1.Group("/products")
		product.GET("/:id", productTransport.GetProductByID(appCtx))

		// Swagger
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return handler
}
