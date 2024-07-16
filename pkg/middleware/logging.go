package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		f, err := os.OpenFile("logs/logstash.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening log file:", err)
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags)
		logger.Printf("%s %s %s", c.Request().Method, c.Request().RequestURI, c.Request().RemoteAddr)
		return next(c)
	}
}
