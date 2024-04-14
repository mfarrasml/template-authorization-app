package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/config"
	"github.com/mfarrasml/template-authorization-app/constant"
	"github.com/mfarrasml/template-authorization-app/dto"
	"github.com/mfarrasml/template-authorization-app/usecase"
	"github.com/mfarrasml/template-authorization-app/util"
)

type UserHandler struct {
	userUc usecase.UserUsecase
	config config.Config
}

func NewUserHandler(userUc usecase.UserUsecase, config config.Config) *UserHandler {
	return &UserHandler{
		userUc: userUc,
		config: config,
	}
}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	req := dto.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	accToken, refToken, err := h.userUc.UserLogin(ctx, req.Email, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "ok",
		Data: dto.NewAuthTokenResponse(*accToken, h.config.JwtAccTknExpiry(), *refToken),
	})
}

func (h *UserHandler) RefreshTokens(ctx *gin.Context) {
	req := dto.RefreshTokenRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	accToken, refToken, err := h.userUc.GetTokensByRefToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "ok",
		Data: dto.NewAuthTokenResponse(accToken, h.config.JwtAccTknExpiry(), refToken),
	})
}

func (h *UserHandler) GetOneAuthorizedUser(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Error(apperror.ErrInvalidUserId)
		return
	}

	claims, ok := ctx.Value(constant.AccessTokenClaims).(*util.UserAuthClaims)
	if !ok {
		ctx.Error(apperror.ErrParsingAccessToken)
		return
	}

	user, err := h.userUc.GetOneById(ctx, id, claims.UserId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "ok",
		Data: dto.NewGetOneUserResponse(*user),
	})
}
