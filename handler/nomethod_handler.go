package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Error(apperror.ErrNoMethod)
	}
}
