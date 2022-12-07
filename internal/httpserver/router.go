package httpserver

import (
	"fmt"
	"github.com/Jungle20m/electricity/component"
	grabHttpTransport "github.com/Jungle20m/electricity/internal/module/grab/http-transport"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewHandler(appCtx component.AppContext) *gin.Engine {
	// Init handler
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	// Middleware
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	handler.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

		v2.GET("/something", getTodoByID)

	}

	return handler
}

// todo represents data about a task in the todo list
type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
func getTodoByID(c *gin.Context) {
	ID := c.Param("id")
	fmt.Println(ID)
	c.JSON(http.StatusNotFound, "success")
}
