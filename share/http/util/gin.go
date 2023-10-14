package utilhttp

import (
	"errors"
	"go-template/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthDataFromGinContext(c *gin.Context) (dto.AuthData, error) {
	ad := dto.AuthData{}
	uid, found := c.Get("user")

	if !found {
		return ad, errors.New("jwt token invalid")
	}

	ad2 := uid.(dto.AuthData)

	return ad2, nil
}

func GetJobIdFromPathParam(c *gin.Context) (uint, error) {
	s := c.Param("jobId")
	bookId, err := strconv.Atoi(s)

	return uint(bookId), err
}

func GetQueryParam(c *gin.Context, query string) string {
	return strings.ToLower(c.Query(query))
}
