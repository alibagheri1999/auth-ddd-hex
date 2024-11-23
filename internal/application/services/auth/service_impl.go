package auth

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/domain/DTO"
	"DDD-HEX/internal/ports/cache"
	"DDD-HEX/internal/ports/repository"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authServiceImpl struct {
	authRepository repository.AuthRepository
	userService    user.UserService
	cacheRepo      cache.CacheRepository
	appConfig      config.AppConfig
	checkHashFunc  utils.CheckHashFunc
}

func NewAuthService(authRepo repository.AuthRepository, userService user.UserService, cacheRepo cache.CacheRepository, appConfig config.AppConfig, checkHashFunc utils.CheckHashFunc) AuthService {
	return &authServiceImpl{authRepository: authRepo, userService: userService, appConfig: appConfig, cacheRepo: cacheRepo, checkHashFunc: checkHashFunc}
}

func (s *authServiceImpl) Authenticate(c context.Context, email, password string) (string, string, error) {
	failedCount, err := HandleFailLogin(c, email, s.cacheRepo)
	if err != nil {
		return "", "", err
	}
	accessTokenExp := s.appConfig.AccessTokenExp
	user, err := s.userService.FindUserByEmail(c, email)
	validPass := s.checkHashFunc(password, user.Password.String)
	if err != nil || !validPass {
		failedCount += 1
		if sErr := s.cacheRepo.SetFailedCount(c, email, failedCount); sErr != nil {
			return "", "", sErr
		}
		if err == nil && !validPass {
			err = errors.New("wrong username or password")
		}
		return "", "", err
	}
	if user.Status == "deactivated" {
		return "", "", errors.New("user is deactivated")
	}
	tokens := GetUserTokens(c, email, s.cacheRepo)
	if tokens != nil {
		return tokens.AccessToken, tokens.RefreshToken, nil
	}
	accessToken, refreshToken, err := s.generateTokens(c, user)
	if err != nil {
		return "", "", err
	}

	auth := domain.AuthEntity{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(),
	}

	if err := s.authRepository.Save(c, auth); err != nil {
		return "", "", err
	}
	tokens = &DTO.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TTL:          uint32(s.appConfig.AccessTokenExp),
	}
	SetUserTokens(c, email, s.cacheRepo, *tokens)
	return accessToken, refreshToken, nil
}

func (s *authServiceImpl) RefreshToken(c context.Context, refreshToken string) (string, string, error) {
	var auth domain.UserAuthEntity
	var err error
	if auth, err = s.authRepository.FindByRefreshToken(c, refreshToken); err != nil {
		return "", "", err
	}
	if auth.Status == "deactivated" {
		return "", "", errors.New("user is deactivated")
	}
	user := &domain.UserEntity{
		Rule:  auth.Rule,
		Email: auth.Email,
		ID:    auth.UserID,
	}
	tokens := GetUserTokens(c, auth.Email, s.cacheRepo)
	if tokens != nil {
		return tokens.AccessToken, tokens.RefreshToken, nil
	}
	accessToken, newRefreshToken, err := s.generateTokens(c, user)
	if err != nil {
		return "", "", err
	}

	auth.AccessToken = accessToken
	auth.RefreshToken = newRefreshToken
	auth.Expires = time.Now().Add(15 * time.Minute).Unix()
	newAuth := domain.AuthEntity{
		UserID:       auth.UserID,
		Expires:      auth.Expires,
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}
	if err := s.authRepository.Save(c, newAuth); err != nil {
		return "", "", err
	}
	tokens = &DTO.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TTL:          uint32(s.appConfig.AccessTokenExp),
	}
	SetUserTokens(c, auth.Email, s.cacheRepo, *tokens)
	return accessToken, newRefreshToken, nil
}

func (s *authServiceImpl) generateTokens(_ context.Context, user *domain.UserEntity) (string, string, error) {
	refreshTokenExp := s.appConfig.RefreshTokenExp
	accessTokenExp := s.appConfig.AccessTokenExp
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Rule,
		"exp":     time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(), // 15 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Rule,
		"exp":     time.Now().Add(time.Duration(refreshTokenExp) * time.Hour).Unix(), // 7 weeks
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.appConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *authServiceImpl) Validate2FA(c context.Context, email string) (string, string, error) {
	user, err := s.userService.FindUserByEmail(c, email)
	if err != nil {
		return "", "", err
	}
	if user.Status == "deactivated" {
		return "", "", errors.New("user is deactivated")
	}
	tokens := GetUserTokens(c, email, s.cacheRepo)
	if tokens != nil {
		return tokens.AccessToken, tokens.RefreshToken, nil
	}
	accessTokenExp := s.appConfig.AccessTokenExp
	accessToken, refreshToken, err := s.generateTokens(c, user)
	auth := domain.AuthEntity{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      time.Now().Add(time.Duration(accessTokenExp) * time.Minute).Unix(),
	}

	if err := s.authRepository.Save(c, auth); err != nil {
		return "", "", err
	}
	tokens = &DTO.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TTL:          uint32(s.appConfig.AccessTokenExp),
	}
	SetUserTokens(c, email, s.cacheRepo, *tokens)
	return accessToken, refreshToken, nil
}

func (s *authServiceImpl) Generate2FACode(c context.Context, username string) (string, error) {
	code := GenerateRandomCode(6)
	if err := s.cacheRepo.Set2FA(c, username, code); err != nil {
		return "", err
	}
	// Send code via email (implementation omitted)
	return code, nil
}

func (s *authServiceImpl) Validate2FACode(c context.Context, username, code string) error {
	storedCode, err := s.cacheRepo.Get2FA(c, username)
	if err != nil || storedCode != code {
		return errors.New("invalid 2FA code")
	}
	return nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
