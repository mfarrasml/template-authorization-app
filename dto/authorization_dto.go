package dto

type AccessTokenRequest struct {
	Token string `json:"access_token" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
