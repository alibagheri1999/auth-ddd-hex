package http

import (
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService auth.AuthService
}

var emptyAccessTokenCookies = &http.Cookie{
	Name:     "access_token",
	Value:    "",
	HttpOnly: true,
}
var emptyRefreshTokenCookies = &http.Cookie{
	Name:     "refresh_token",
	Value:    "",
	HttpOnly: true,
}

func (h *AuthHandler) Login(c echo.Context) error {
	config := utils.ConfigSetup()
	appCfg := config.App
	refreshTokenExp := appCfg.RefreshTokenExp
	accessTokenExp := appCfg.AccessTokenExp
	var req DTO.LoginRequest
	var res DTO.LoginResponse

	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	accessToken, refreshToken, err := h.AuthService.Authenticate(c, req.Email, req.Password)
	if err != nil {
		c.SetCookie(emptyAccessTokenCookies)
		c.SetCookie(emptyRefreshTokenCookies)
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenExp) * time.Minute),
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Duration(refreshTokenExp) * time.Hour),
		HttpOnly: true,
	})

	res.Message = "login"
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	var res DTO.LoginResponse
	config := utils.ConfigSetup()
	appCfg := config.App
	refreshTokenExp := appCfg.RefreshTokenExp
	accessTokenExp := appCfg.AccessTokenExp
	refreshTokenCookie, err := c.Cookie("refresh_token")
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	accessToken, refreshToken, err := h.AuthService.RefreshToken(c, refreshTokenCookie.Value)
	if err != nil {
		c.SetCookie(emptyAccessTokenCookies)
		c.SetCookie(emptyRefreshTokenCookies)
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenExp) * time.Minute),
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Duration(refreshTokenExp) * time.Hour),
		HttpOnly: true,
	})

	res.Message = "refreshed"
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Generate2FACode(c echo.Context) error {
	var req DTO.GenerateCodeRequest
	var res DTO.GenerateCodeResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		res.Code = ""
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	code, err := h.AuthService.Generate2FACode(c, req.Email)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	res.Message = "generated"
	res.Code = code
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Validate2FACode(c echo.Context) error {
	var req DTO.ValidateCodeRequest
	var res DTO.ValidateCodeResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	config := utils.ConfigSetup()
	appCfg := config.App
	refreshTokenExp := appCfg.RefreshTokenExp
	accessTokenExp := appCfg.AccessTokenExp

	err := h.AuthService.Validate2FACode(c, req.Email, req.Code)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}
	accessToken, refreshToken, err := h.AuthService.GenerateTokens(c, req.Email)
	if err != nil {
		c.SetCookie(emptyAccessTokenCookies)
		c.SetCookie(emptyRefreshTokenCookies)
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenExp) * time.Minute),
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Duration(refreshTokenExp) * time.Hour),
		HttpOnly: true,
	})

	res.Message = "login"
	return c.JSON(http.StatusOK, res)
}
