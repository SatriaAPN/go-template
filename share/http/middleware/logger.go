package middleware

import (
	dto "go-template/dto/general"
	dtohttp "go-template/dto/http"
	"go-template/share/general/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := util.GetLogger()
		tn := time.Now()

		requestData := dtohttp.NewHttpRequestLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get(dto.RequestIdKey), "request")
		logger.Infof(requestData)

		c.Next()

		tp := time.Since(tn)
		responseData := dtohttp.NewHttpResponseLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get(dto.RequestIdKey), "response", c.Writer.Status(), tp)
		logger.Infof(responseData)
	}
}
