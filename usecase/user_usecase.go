package usecase

import (
	"context"

	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/entity"
	"github.com/mfarrasml/template-authorization-app/repository"
	"github.com/mfarrasml/template-authorization-app/util"
)

type UserUsecase interface {
	UserLogin(ctx context.Context, email string, password string) (*string, *string, error)
	GetOneById(ctx context.Context, id int, authId int) (*entity.User, error)
}

type UserUcImpl struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	passwordUtil     util.PasswordUtil
	tokenUtil        util.TokenUtil
}

type UserUcImplOpt struct {
	UserRepo         repository.UserRepository
	RefreshTokenRepo repository.RefreshTokenRepository
	PasswordUtil     util.PasswordUtil
	TokenUtil        util.TokenUtil
}

func NewUserUcImpl(opt UserUcImplOpt) *UserUcImpl {
	return &UserUcImpl{
		userRepo:         opt.UserRepo,
		refreshTokenRepo: opt.RefreshTokenRepo,
		passwordUtil:     opt.PasswordUtil,
		tokenUtil:        opt.TokenUtil,
	}
}

func (u *UserUcImpl) UserLogin(ctx context.Context, email string, password string) (*string, *string, error) {
	user, err := u.userRepo.FindOneByEmail(ctx, email)
	if err == apperror.ErrUserNotFound {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, apperror.ErrInternalServer
	}

	err = u.passwordUtil.ComparePwdWithHash(password, []byte(user.Password))
	if err != nil {
		return nil, nil, apperror.ErrWrongPassword
	}

	// generate new access token
	accToken, err := u.tokenUtil.NewAuthToken(user.Id, user.Email)
	if err != nil {
		return nil, nil, apperror.ErrAccessToken
	}

	// also generate new refresh token every login
	refToken, jti, err := u.tokenUtil.NewRefreshToken(user.Id)
	if err != nil {
		return nil, nil, apperror.ErrRefreshToken
	}

	// send refresh token's jti to db
	err = u.refreshTokenRepo.CreateOne(ctx, jti)
	if err != nil {
		return nil, nil, err
	}

	return &accToken, &refToken, nil
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
