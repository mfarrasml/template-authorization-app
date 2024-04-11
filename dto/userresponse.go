package dto

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func NewUserLoginResponse(token string) UserLoginResponse {
	return UserLoginResponse{
		Token: token,
	}
}
