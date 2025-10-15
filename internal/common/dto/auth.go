package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,lte=100"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken" validate:"required"`
}
