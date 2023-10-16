package middleware

import (
	dtohttp "go-template/dto/http"
	"go-template/share/general/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := util.GetLogger()
		tn := time.Now()

		requestData := dtohttp.NewHttpRequestLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "request")
		logger.Infof(requestData)

		c.Next()

		tp := time.Since(tn)
		responseData := dtohttp.NewHttpResponseLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "response", c.Writer.Status(), tp)
		logger.Infof(responseData)
	}
}
