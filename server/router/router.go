package router

import (
	"go-template/handler"
	"go-template/share/http/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	GetRouter() *gin.Engine
}

type router struct {
	router      *gin.Engine
	userHandler handler.UserHandler
}

type RouterConfig struct {
	UserHandler handler.UserHandler
}

func NewRouter(config RouterConfig) Router {
	return &router{
		router:      gin.New(),
		userHandler: config.UserHandler,
	}
}

func (rb *router) GetRouter() *gin.Engine {
	rb.buildMiddleware()
	rb.buildEndpointHandler()

	return rb.router
}

func (rb *router) buildEndpointHandler() {
	rb.router.POST("/users", rb.userHandler.CreateUser)
	rb.router.POST("/login", rb.userHandler.Login)
	rb.router.GET("/profile", middleware.JwtDecoding(), rb.userHandler.GetProfile)
	rb.router.POST("/forget-password", rb.userHandler.ForgetPassword)
	rb.router.POST("/reset-password", rb.userHandler.ResetPassword)

}

func (rb *router) buildMiddleware() {
	rb.router.Use(middleware.SetRequestId())

	rb.router.Use(middleware.Logger())

	rb.router.Use(middleware.GlobalErrorHandler())

	rb.router.Use(middleware.HttpRequestTimeout())
}
