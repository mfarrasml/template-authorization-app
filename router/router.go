package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/handler"
	"github.com/mfarrasml/template-authorization-app/middleware"
	"github.com/mfarrasml/template-authorization-app/util"
)

type HandlerOpt struct {
	userHandler *handler.UserHandler
	tokenUtil   util.TokenUtil
}

func NewRouter(opt HandlerOpt) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	router.Use(gin.Recovery(), gin.Logger())
	router.Use(middleware.ErrorHandler())

	router.POST("/auth/login", opt.userHandler.UserLogin)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizationMiddleware(opt.tokenUtil))
	{
		authorized.GET("/users/:id", opt.userHandler.GetOneAuthorizedUser)
	}

	router.NoRoute(handler.NoRouteHandler())
	router.NoMethod(handler.NoMethodHandler())

	return router
}
