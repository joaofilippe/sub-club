package server

import (
	"log/slog"
	"net/http"

	"github.com/joaofilippe/subclub/internal/infra/middleware"
	"github.com/joaofilippe/subclub/internal/web"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	logger *slog.Logger
	router *router
}

func NewServer(handlers *web.Handlers) *Server {
	logger := slog.Default()
	e := echo.New()

	s := &Server{
		echo:   e,
		logger: logger,
		router: newRouter(e, handlers),
	}

	e.Use(middleware.ConfigureLogger(s.logger))
	e.Use(echoMiddleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	s.router.registerRoutes()

	return s
}

func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}

func (s *Server) Start(port string) error {
	return s.echo.Start(port)
}
