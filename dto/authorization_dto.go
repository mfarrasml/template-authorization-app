package dto

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthTokenResponse(accToken string, expiresIn int, refToken string) AuthTokenResponse {
	return AuthTokenResponse{
		AccessToken:  accToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn * 60,
		RefreshToken: refToken,
	}
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
