package transportation

import (
	"fmt"
	"github.com/Jungle20m/electricity/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSomething(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		msql := appCtx.GetMysql()

		fmt.Printf("connection: %+v\n", msql.DB)

		c.JSON(http.StatusOK, gin.H{
			"data": "success",
		})
	}
}
