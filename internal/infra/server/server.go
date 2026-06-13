package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/joaofilippe/subclub/internal/infra/database"
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

func NewServer(handlers *web.Handlers, tenantManager *database.TenantClientManager, jwtSecret []byte, isDev bool) *Server {
	logger := slog.Default()
	e := echo.New()

	s := &Server{
		echo:   e,
		logger: logger,
		router: newRouter(e, handlers, tenantManager, jwtSecret),
	}

	e.Use(middleware.ConfigureLogger(s.logger))
	e.Use(echoMiddleware.Recover())

	corsConfig := echoMiddleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}
	if isDev {
		corsConfig.AllowOriginFunc = func(origin string) (bool, error) {
			return strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "https://localhost"), nil
		}
	}
	e.Use(echoMiddleware.CORSWithConfig(corsConfig))

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
	if port != "" && port[0] != ':' {
		port = ":" + port
	}
	return s.echo.Start(port)
}
