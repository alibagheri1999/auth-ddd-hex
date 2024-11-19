package server

import (
	middlewares "DDD-HEX/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// NewRouter initiate new route to using api and middlewares
func NewRouter() *echo.Echo {

	router := echo.New()
	router.HidePort = true
	router.HideBanner = true

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "object not found"})
	}

	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, echo.Map{"message": "method not allowed"})
	}

	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.SecurityHeaders)
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	router.Use(middlewares.CORS())
	//router.Use(middleware.Logger())
	return router
}
