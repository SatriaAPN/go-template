package middleware

import (
	dtohttp "go-template/dto/http"
	"go-template/share/util"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := util.GetLogger()

		requestData := dtohttp.NewHttpRequestLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "request")
		logger.Infof(requestData)

		c.Next()

		responseData := dtohttp.NewHttpResponseLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "response", c.Writer.Status())
		logger.Infof(responseData)
	}
}
