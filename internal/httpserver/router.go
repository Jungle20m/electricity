package httpserver

import (
	"github.com/Jungle20m/electricity/component"
	_ "github.com/Jungle20m/electricity/docs"
	grabHttpTransport "github.com/Jungle20m/electricity/internal/module/grab/http-transport"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @host localhost:8000
// @BasePath /
// @query.collection.format multi
func NewHandler(appCtx component.AppContext) *gin.Engine {
	// Init handler
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	// Middleware
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

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

	// Swagger
	handler.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return handler
}

//// response represents response data
//type Response struct {
//	Code         int         `json:"code"`
//	Status       string      `json:"status"`
//	ErrorCode    string      `json:"error_code"`
//	ErrorMessage string      `json:"error_message"`
//	Message      string      `json:"message,omitempty"`
//	Data         interface{} `json:"data,omitempty"`
//}
//
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
//
//// @Summary get a todo item by ID
//// @ID get-todo-by-id
//// @Produce json
//// @Param id path string true "todo ID"
//// @Success 200 {object} todo
//// @Failure 404 {object} message
//// @Router /todo/{id} [get]
