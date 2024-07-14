package http

import (
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain/DTO"
	"encoding/json"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService auth.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	config := utils.ConfigSetup()
	appCfg := config.App
	refreshTokenExp := appCfg.RefreshTokenExp
	accessTokenExp := appCfg.AccessTokenExp
	var req DTO.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.AuthService.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenExp) * time.Minute),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Duration(refreshTokenExp) * time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	config := utils.ConfigSetup()
	appCfg := config.App
	refreshTokenExp := appCfg.RefreshTokenExp
	accessTokenExp := appCfg.AccessTokenExp
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token missing", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := h.AuthService.RefreshToken(refreshTokenCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenExp) * time.Minute),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Duration(refreshTokenExp) * time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}
