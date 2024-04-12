package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/handler"
	"github.com/mfarrasml/template-authorization-app/middleware"
)

type HandlerOpt struct {
	userHandler *handler.UserHandler
}

func NewRouter(opt HandlerOpt) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(gin.Recovery(), gin.Logger())
	router.Use(middleware.ErrorHandler())

	router.POST("/auth/login", opt.userHandler.UserLogin)

	return router
}
