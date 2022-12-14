package http_transport

import (
	"github.com/Jungle20m/electricity/common"
	grabBusiness "github.com/Jungle20m/electricity/internal/modules/grab/business"
	grabModel "github.com/Jungle20m/electricity/internal/modules/grab/model"
	gabStorage "github.com/Jungle20m/electricity/internal/modules/grab/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateService(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData grabModel.ServiceCreate

		if err := c.ShouldBind(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		storage := gabStorage.NewStorage(appCtx.GetMysql().DB)
		business := grabBusiness.NewCreateServiceBusiness(appCtx, storage)

		err := business.CreateNewService(c.Request.Context(), requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "create service",
		})
	}
}
