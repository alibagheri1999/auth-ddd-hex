package server

import (
	"DDD-HEX/internal/ports/clients"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// NewServer initiate new instance of http server
func NewServer(router *echo.Echo, port int, gracefulShutdown time.Duration) *Server {
	return &Server{
		addr:             fmt.Sprintf("%s:%v", os.Getenv("HOST_IP"), port),
		router:           router,
		gracefulShutdown: gracefulShutdown,
	}
}

type Server struct {
	addr             string
	router           *echo.Echo
	gracefulShutdown time.Duration
}

// StartListening force server to start listening on a port
func (s *Server) StartListening(cache clients.Cache, db clients.Database) {
	logrus.Info(fmt.Sprintf("server start to listen on %s\n", s.addr))
	go func() {
		if err := s.router.Start(s.addr); err != nil && err != http.ErrServerClosed {
			logrus.Error("Err server start", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	<-quit

	logrus.Info(fmt.Sprintf("server shutting down in %s...\n", s.gracefulShutdown))
	c, cancel := context.WithTimeout(context.Background(), s.gracefulShutdown)
	defer cancel()
	if err := s.router.Shutdown(c); err != nil {
		logrus.Info("Err server shutdown", err)
	}
	db.Close()
	cache.Close()
	<-c.Done()
	logrus.Info("Good Luck!")
}
