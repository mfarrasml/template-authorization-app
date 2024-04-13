package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/constant"
	"github.com/mfarrasml/template-authorization-app/dto"
	"github.com/mfarrasml/template-authorization-app/util"
)

func AuthorizationMiddleware(tokenUtil util.TokenUtil) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.AccessTokenRequest{}
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.Error(apperror.ErrBadRequest)
			ctx.Abort()
			return
		}

		tokenList := strings.Fields(req.Token)
		if len(tokenList) < 2 || tokenList[0] != "Bearer" {
			ctx.Error(apperror.ErrInvalidAccessToken)
			ctx.Abort()
			return
		}
		token := tokenList[1]

		claims, err := tokenUtil.ParseAuthToken(token)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set(constant.AccessTokenClaims, claims)

		ctx.Next()
	}
}
