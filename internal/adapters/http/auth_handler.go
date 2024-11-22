package http

import (
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	AuthService auth.AuthService
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req DTO.LoginRequest
	var res DTO.LoginResponse
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if !utils.IsValidEmail(req.Email) {
		res.Message = "Enter a valid email"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	accessToken, refreshToken, err := h.AuthService.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}
	res = DTO.LoginResponse{
		Message:      "login",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	ctx := c.Request().Context()
	var res DTO.LoginResponse
	refreshTokenCookie := c.Request().Header.Get("refresh_token")
	accessToken, refreshToken, err := h.AuthService.RefreshToken(ctx, refreshTokenCookie)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	res = DTO.LoginResponse{
		Message:      "refreshed",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Generate2FACode(c echo.Context) error {
	ctx := c.Request().Context()
	var req DTO.GenerateCodeRequest
	var res DTO.GenerateCodeResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		res.Code = ""
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if !utils.IsValidEmail(req.Email) {
		res.Message = "Enter a valid email"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	code, err := h.AuthService.Generate2FACode(ctx, req.Email)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}

	res.Message = "generated"
	res.Code = code
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Validate2FACode(c echo.Context) error {
	ctx := c.Request().Context()
	var req DTO.ValidateCodeRequest
	var res DTO.ValidateCodeResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if !utils.IsValidEmail(req.Email) {
		res.Message = "Enter a valid email"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	err := h.AuthService.Validate2FACode(ctx, req.Email, req.Code)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}
	accessToken, refreshToken, err := h.AuthService.Validate2FA(ctx, req.Email)
	if err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusUnauthorized, res)
	}
	res = DTO.ValidateCodeResponse{
		Message:      "login",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusOK, res)
}
