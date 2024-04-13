package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/constant"
	"github.com/mfarrasml/template-authorization-app/dto"
	"github.com/mfarrasml/template-authorization-app/usecase"
	"github.com/mfarrasml/template-authorization-app/util"
)

type UserHandler struct {
	userUc usecase.UserUsecase
}

func NewUserHandler(userUc usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUc: userUc,
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
		Data: dto.NewUserLoginResponse(*accToken, *refToken),
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
