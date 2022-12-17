package middleware

import (
	"fmt"
	"github.com/Jungle20m/electricity/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"strings"
)

const (
	CredentialKeyName = "credential"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

type credential struct {
	WorkingSiteID string `json:"working_site_id"`
	WorkerSiteID  string `json:"worker_site_id"`
}

func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
				"message": err,
			})
			c.Abort()
			return
		}

		token := strings.Split(h.Token, "Bearer ")
		if len(token) < 2 {
			err := fmt.Errorf("must provide Authorization header with format `Bearer Token`")
			c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, err))
			c.Abort()
			return
		}

		bearerToken := token[1]
		data, err := jwt.ParseSigned(bearerToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, err))
			c.Abort()
			return
		}

		credential := credential{}
		err = data.UnsafeClaimsWithoutVerification(&credential)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewErrorResponse(http.StatusBadRequest, err))
			c.Abort()
			return
		}

		c.Set(CredentialKeyName, credential)
		c.Next()
	}
}

func GetCredential(ctx *gin.Context) (*credential, error) {
	cred, exist := ctx.Get(CredentialKeyName)
	if exist != true {
		return nil, fmt.Errorf("bearer token not found")
	}
	output := cred.(credential)
	return &output, nil
}
