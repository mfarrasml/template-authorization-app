package usecase

import (
	"context"
	"errors"

	"github.com/mfarrasml/template-authorization-app/repository"
	"github.com/mfarrasml/template-authorization-app/util"
)

type UserUsecase interface {
	UserLogin(ctx context.Context, email string, password string) (*string, error)
}

type UserUcImpl struct {
	userRepo     repository.UserRepository
	passwordUtil util.PasswordUtil
	tokenUtil    util.TokenUtil
}

func NewUserUcImpl(userRepo repository.UserRepository, passwordUtil util.PasswordUtil, tokenUtil util.TokenUtil) *UserUcImpl {
	return &UserUcImpl{
		userRepo:     userRepo,
		passwordUtil: passwordUtil,
		tokenUtil:    tokenUtil,
	}
}

func (u *UserUcImpl) UserLogin(ctx context.Context, email string, password string) (*string, error) {
	user, err := u.userRepo.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = u.passwordUtil.ComparePwdWithHash(password, []byte(user.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	token, err := u.tokenUtil.NewAuthToken(email)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
