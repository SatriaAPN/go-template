package middleware

import (
	dto "go-template/dto/general"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewString()
		c.Header(dto.RequestIdKey, uuid)
		c.Next()
	}
}
