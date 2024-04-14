package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Error(apperror.ErrNoRoute)
	}
}
