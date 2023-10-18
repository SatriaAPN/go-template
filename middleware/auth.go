package middleware

import (
	dtohttp "go-template/dto/http"
	"go-template/share/general/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusBadRequest, dtohttp.Response{Message: "Invalid Token"})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		ud, err := util.GetAuthTokenGenerator().Decode(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		c.Set("user", ud)

		c.Next()
	}
}
