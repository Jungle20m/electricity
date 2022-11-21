package http_transport

import (
	"github.com/Jungle20m/electricity/component"
	grabBusiness "github.com/Jungle20m/electricity/internal/module/grab/business"
	grabModel "github.com/Jungle20m/electricity/internal/module/grab/model"
	gabStorage "github.com/Jungle20m/electricity/internal/module/grab/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateService(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData grabModel.ServiceCreate

		if err := c.ShouldBind(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		storage := gabStorage.NewStorage(appCtx.GetMysql().DB)
		business := grabBusiness.NewCreateServiceBusiness(appCtx, storage)

		business.CreateNewService(c.Request.Context(), requestData)

		c.JSON(http.StatusOK, gin.H{
			"data": "create service",
		})
	}
}
