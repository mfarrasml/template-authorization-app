package dto

import "github.com/mfarrasml/template-authorization-app/entity"

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthTokenResponse(accToken string, refToken string) AuthTokenResponse {
	return AuthTokenResponse{
		AccessToken:  accToken,
		RefreshToken: refToken,
	}
}

type GetOneUserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewGetOneUserResponse(user entity.User) GetOneUserResponse {
	return GetOneUserResponse{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}
