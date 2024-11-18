package DTO

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GenerateCodeRequest struct {
	Email string `json:"email"`
}

type GenerateCodeResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ValidateCodeResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidateCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
