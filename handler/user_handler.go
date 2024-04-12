package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfarrasml/template-authorization-app/dto"
	"github.com/mfarrasml/template-authorization-app/usecase"
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

	token, err := h.userUc.UserLogin(ctx, req.Email, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "ok",
		Data: dto.NewUserLoginResponse(*token),
	})
}
