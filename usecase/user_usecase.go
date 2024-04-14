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
	GetTokensByRefToken(ctx context.Context, refToken string) (newAccToken string, newRefToken string, err error)
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

	accToken, refToken, jti, err := generateNewTokenPair(user.Id, user.Email, u.tokenUtil)
	if err != nil {
		return nil, nil, err
	}

	// send refresh token's jti to db
	err = u.refreshTokenRepo.CreateOne(ctx, jti, user.Id)
	if err != nil {
		return nil, nil, err
	}

	return &accToken, &refToken, nil
}

func (u *UserUcImpl) GetTokensByRefToken(ctx context.Context, refToken string) (newAccToken string, newRefToken string, err error) {
	claims, err := u.tokenUtil.ParseRefreshToken(refToken)
	if err != nil {
		return
	}

	refTokenRcrd, err := u.refreshTokenRepo.FindOneByUserId(ctx, claims.UserId)
	if err != nil {
		return
	}

	// check if user's refresh token is the newest (valid) refresh token for said user
	if refTokenRcrd.Jti != claims.ID {
		err = apperror.ErrInvalidRefreshToken
		return
	}

	user, err := u.userRepo.FindOneById(ctx, refTokenRcrd.UserId)
	if err != nil {
		err = apperror.ErrUserNotFound
		return
	}

	newAccToken, newRefToken, jti, err := generateNewTokenPair(user.Id, user.Email, u.tokenUtil)
	if err != nil {
		return "", "", err
	}

	// send refresh token's jti to db
	err = u.refreshTokenRepo.CreateOne(ctx, jti, user.Id)
	if err != nil {
		return "", "", err
	}

	err = nil
	return
}

func generateNewTokenPair(userId int, email string, tokenUtil util.TokenUtil) (accToken string, refToken string, refTokenJti string, err error) {
	// generate new access token
	accToken, err = tokenUtil.NewAuthToken(userId, email)
	if err != nil {
		return "", "", "", apperror.ErrAccessToken
	}

	// also generate new refresh token every login
	refToken, refTokenJti, err = tokenUtil.NewRefreshToken(userId)
	if err != nil {
		return "", "", "", apperror.ErrRefreshToken
	}

	err = nil
	return
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
