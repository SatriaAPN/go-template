package middleware

import (
	dto "go-template/dto/general"
	dtohttp "go-template/dto/http"
	errorapp "go-template/share/general/error"
	"go-template/share/general/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		if err != nil {
			logger := util.GetLogger()
			switch e := err.Err.(type) {
			case *errors.Error:
				stackTrace := e.ErrorStack()

				logger.Errorf(dtohttp.NewErrorLoggerData("error", c.Writer.Header().Get(dto.RequestIdKey), stackTrace))

				rCode := http.StatusInternalServerError
				rMessage := e.Error()

				checkErrorStruct(e, &rCode, &rMessage)

				c.AbortWithStatusJSON(rCode, gin.H{
					"error": rMessage,
				})
			case *time.ParseError:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			default:
				logger.Errorf(dtohttp.NewErrorLoggerData("error", c.Writer.Header().Get(dto.RequestIdKey), "unhandled"))

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Something Wrong has Happened",
				})
			}
		}
	}
}

func checkErrorStruct(e error, rCode *int, s *string) {
	switch {
	case errors.Is(e, errorapp.ErrEmailAlreadyExist):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrEmailIsNotValid):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrWrongPassword):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrEmailNotFound):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrMinimumPasswordLength):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrForgetPasswordTokenLength):
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrResetCodeNotFound):
		*rCode = http.StatusBadRequest
		*rCode = http.StatusBadRequest
	case errors.Is(e, errorapp.ErrMaximumPasswordLength):
		*rCode = http.StatusBadRequest
	default:
		*s = "Something Wrong has Happened"
	}
}
