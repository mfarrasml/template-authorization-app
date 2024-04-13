package dto

type AccessTokenRequest struct {
	Token string `json:"access_token" binding:"required"`
}
