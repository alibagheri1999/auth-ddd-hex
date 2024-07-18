package DTO

type CreateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}
