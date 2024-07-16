package auth

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

type authServiceImpl struct {
	authRepository repository.AuthRepository
	userService    user.UserService
	cacheRepo      repository.CacheRepository
	appConfig      config.AppConfig
}

func NewAuthService(authRepo repository.AuthRepository, userService user.UserService, cacheRepo repository.CacheRepository, appConfig config.AppConfig) AuthService {
	return &authServiceImpl{authRepository: authRepo, userService: userService, appConfig: appConfig, cacheRepo: cacheRepo}
}

func (s *authServiceImpl) Authenticate(c echo.Context, email, password string) (string, string, error) {
	failedCount, err := HandleFailLogin(email, s.cacheRepo)
	if err != nil {
		return "", "", err
	}
	accessTokenExp := s.appConfig.AccessTokenExp
	user, err := s.userService.FindUserByEmail(c, email)
	if err != nil || !utils.CheckHash(password, user.Password) { // Assume this function checks the password hash
		failedCount += 1
		if sErr := s.cacheRepo.SetFailedCount(email, failedCount); sErr != nil {
			return "", "", sErr
		}
		return "", "", err
	}

	accessToken, refreshToken, err := s.generateTokens(c, user.ID)
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

func (s *authServiceImpl) RefreshToken(c echo.Context, refreshToken string) (string, string, error) {
	auth, err := s.authRepository.FindByRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, newRefreshToken, err := s.generateTokens(c, auth.UserID)
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

func (s *authServiceImpl) generateTokens(c echo.Context, userID string) (string, string, error) {
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

func (s *authServiceImpl) GenerateTokens(c echo.Context, email string) (string, string, error) {
	user, err := s.userService.FindUserByEmail(c, email)
	refreshTokenExp := s.appConfig.RefreshTokenExp
	accessTokenExp := s.appConfig.AccessTokenExp
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(), // 15 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Duration(refreshTokenExp) * time.Hour).Unix(), // 7 weeks
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *authServiceImpl) Generate2FACode(c echo.Context, username string) (string, error) {
	code := GenerateRandomCode(6)
	if err := s.cacheRepo.Set2FA(username, code); err != nil {
		return "", err
	}
	// Send code via email (implementation omitted)
	return code, nil
}

func (s *authServiceImpl) Validate2FACode(c echo.Context, username, code string) error {
	storedCode, err := s.cacheRepo.Get2FA(username)
	if err != nil || storedCode != code {
		return errors.New("invalid 2FA code")
	}
	return nil
}
