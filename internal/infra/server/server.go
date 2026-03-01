package server

import (
	"net/http"

	"github.com/joaofilippe/subclub/internal/infra/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Server representa o servidor web da aplicação
type Server struct {
	echo *echo.Echo
}

// NewServer inicializa e configura um novo servidor
func NewServer() *Server {
	e := echo.New()

	// Middleware globais
	e.Use(middleware.ConfigureLogger())
	e.Use(echoMiddleware.Recover())

	// Rotas base
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	return &Server{
		echo: e,
	}
}

// Start inicia o servidor na porta especificada
func (s *Server) Start(port string) error {
	return s.echo.Start(port)
}
