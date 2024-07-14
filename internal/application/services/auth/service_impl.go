package auth

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authServiceImpl struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
	appConfig      config.AppConfig
}

func NewAuthService(authRepo repository.AuthRepository, userRepo repository.UserRepository, appConfig config.AppConfig) AuthService {
	return &authServiceImpl{authRepository: authRepo, userRepository: userRepo, appConfig: appConfig}
}

func (s *authServiceImpl) Authenticate(email, password string) (string, string, error) {
	accessTokenExp := s.appConfig.AccessTokenExp
	user, err := s.userRepository.FindByEmail(email)
	if err != nil || !utils.CheckHash(password, user.Password) { // Assume this function checks the password hash
		return "", "", err
	}

	accessToken, refreshToken, err := s.generateTokens(user.ID)
	if err != nil {
		return "", "", err
	}

	auth := domain.Auth{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(),
	}

	if err := s.authRepository.Save(auth); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authServiceImpl) RefreshToken(refreshToken string) (string, string, error) {
	auth, err := s.authRepository.FindByRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, newRefreshToken, err := s.generateTokens(auth.UserID)
	if err != nil {
		return "", "", err
	}

	auth.AccessToken = accessToken
	auth.RefreshToken = newRefreshToken
	auth.Expires = time.Now().Add(15 * time.Minute).Unix()

	if err := s.authRepository.Save(auth); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *authServiceImpl) generateTokens(userID string) (string, string, error) {
	refreshTokenExp := s.appConfig.RefreshTokenExp
	accessTokenExp := s.appConfig.AccessTokenExp
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(), // 15 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(refreshTokenExp) * time.Hour).Unix(), // 7 weeks
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
