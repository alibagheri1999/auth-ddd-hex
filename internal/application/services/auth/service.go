package auth

type AuthService interface {
	Authenticate(email, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, string, error)
}
