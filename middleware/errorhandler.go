package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/dto"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, err := range ctx.Errors {
			switch err.Err {
			case apperror.ErrEmailNotFound:
				fallthrough
			case apperror.ErrWrongPassword:
				ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
					Msg: "wrong email or password",
				})
			case apperror.ErrNoRoute:
				ctx.AbortWithStatusJSON(http.StatusNotFound, dto.Response{
					Msg: apperror.ErrNoRoute.Error(),
				})
			case apperror.ErrNoMethod:
				ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, dto.Response{
					Msg: apperror.ErrNoMethod.Error(),
				})
			default:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
					Msg: "server error",
				})
			}
		}
	}
}
