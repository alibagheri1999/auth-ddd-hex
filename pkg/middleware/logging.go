package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		method := c.Request().Method
		requestURI := c.Request().RequestURI
		remoteAddr := c.Request().RemoteAddr
		bodyLength := c.Request().ContentLength
		err := next(c)
		statusCode := c.Response().Status
		logrus.Infof("%s %d %s %s %s %d bytes", time.Now().Format("2006-01-02 15:04:05"), statusCode, method, requestURI, remoteAddr, bodyLength)
		logger.SetOutput(os.Stdout)
		return err
	}
}
