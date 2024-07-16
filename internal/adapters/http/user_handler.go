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
	var req DTO.UserRequest
	var res DTO.UserResponse
	if err := c.Bind(&req); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if err := h.UserService.CreateUser(c, req.Name, req.Email, req.Password); err != nil {
		res.Message = err.Error()
		return echo.NewHTTPError(http.StatusInternalServerError, res)
	}
	res.Message = "created"
	return c.JSON(http.StatusCreated, res)
}
