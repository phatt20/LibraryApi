package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/phatt20/LibraryApi/config"
	"github.com/phatt20/LibraryApi/databases"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Databases
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEhoServer(conf *config.Config, db databases.Databases) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})
	return server
}

func (s *echoServer) Start() {
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	// Prevent application from crashing
	s.app.Use(middleware.Recover())

	s.app.Use(middleware.Logger())
	s.app.Use(timeOutMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)

	s.app.GET("/v1/health", s.healthCheck)

	s.initBookRouter()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Start(url); err != nil && err != http.ErrBodyReadAfterClose {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) gracefullyShutdown(quitch chan os.Signal) {
	ctx := context.Background()
	<-quitch
	s.app.Logger.Info("Shutting down server...")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}
func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout:      timeout * time.Second,
	})
}

func getBodyLimitMiddleware(bodylimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodylimit)
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}
