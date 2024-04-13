package usecase

import (
	"context"

	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/entity"
	"github.com/mfarrasml/template-authorization-app/repository"
	"github.com/mfarrasml/template-authorization-app/util"
)

type UserUsecase interface {
	UserLogin(ctx context.Context, email string, password string) (*string, error)
	GetOneById(ctx context.Context, id int, authId int) (*entity.User, error)
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
	if err == apperror.ErrUserNotFound {
		return nil, err
	}
	if err != nil {
		return nil, apperror.ErrInternalServer
	}

	err = u.passwordUtil.ComparePwdWithHash(password, []byte(user.Password))
	if err != nil {
		return nil, apperror.ErrWrongPassword
	}

	token, err := u.tokenUtil.NewAuthToken(user.Id, user.Email)
	if err != nil {
		return nil, apperror.ErrAccessToken
	}
	return &token, nil
}

func (u *UserUcImpl) GetOneById(ctx context.Context, id int, authId int) (*entity.User, error) {
	if id != authId {
		return nil, apperror.ErrForbidden
	}

	user, err := u.userRepo.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
