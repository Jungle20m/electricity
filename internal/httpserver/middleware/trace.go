package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

const (
	RequestIDKey = "X-Request-ID"
)

// generator a function type that returns string.
type generator func() string

var (
	random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func uuid(len int) string {
	bytes := make([]byte, len)
	random.Read(bytes)
	return base64.StdEncoding.EncodeToString(bytes)[:len]
}

// RequestID is a middleware that injects a 'RequestID' into the context and header of each request.
func RequestID(gen generator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		if gen != nil {
			requestID = gen()
		} else {
			requestID = uuid(16)
		}
		c.Request.Header.Set(RequestIDKey, requestID)
		c.Set(RequestIDKey, requestID)
		c.Next()
	}
}

//GetDefaultLogFormatterWithRequestID returns gin.LogFormatter with 'RequestID'
func GetDefaultLogFormatterWithRequestID() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN-debug] %s [%s] - [%s] \"%s %s %s %d %s\" %s %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Request.Header.Get(RequestIDKey),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}

// GetRequestIDFromContext returns 'RequestID' from the given context if present.
func GetRequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(RequestIDKey); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}
	return ""
}

// GetRequestIDFromHeaders returns 'RequestID' from the headers if present.
func GetRequestIDFromHeaders(c *gin.Context) string {
	return c.Request.Header.Get(RequestIDKey)
}
