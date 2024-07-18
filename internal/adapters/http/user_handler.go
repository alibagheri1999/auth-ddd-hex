package http

import (
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserService user.UserService
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req DTO.CreateUserRequest
	var res DTO.CreateUserResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if err := h.UserService.CreateUser(c, req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusInternalServerError, res)
	}
	res.Message = "created"
	return c.JSON(http.StatusCreated, res)
}
