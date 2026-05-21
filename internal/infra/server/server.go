package server

import (
	"net/http"

	_ "github.com/joaofilippe/subclub/docs"
	"github.com/joaofilippe/subclub/internal/infra/middleware"
	"github.com/joaofilippe/subclub/internal/web"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(handlers *web.Handlers) *Server {
	e := echo.New()

	e.Use(middleware.ConfigureLogger())
	e.Use(echoMiddleware.Recover())

	s := &Server{echo: e}
	s.registerRoutes(handlers)

	return s
}

func (s *Server) registerRoutes(h *web.Handlers) {
	s.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	clientGroup := s.echo.Group("/clients")
	clientGroup.POST("", h.Client.Create)
	clientGroup.GET("", h.Client.List)
	clientGroup.GET("/:id", h.Client.Get)
	clientGroup.PUT("/:id", h.Client.Update)
	clientGroup.DELETE("/:id", h.Client.Delete)

	planGroup := s.echo.Group("/plans")
	planGroup.POST("", h.Plan.Create)
	planGroup.GET("", h.Plan.List)
	planGroup.GET("/:id", h.Plan.Get)
	planGroup.PUT("/:id", h.Plan.Update)
	planGroup.DELETE("/:id", h.Plan.Delete)

	subGroup := s.echo.Group("/subscriptions")
	subGroup.POST("", h.Subscription.Create)
	subGroup.GET("", h.Subscription.List)
	subGroup.GET("/:id", h.Subscription.Get)
	subGroup.PUT("/:id", h.Subscription.Update)
	subGroup.DELETE("/:id", h.Subscription.Delete)

	productGroup := s.echo.Group("/products")
	productGroup.POST("", h.Product.Create)
	productGroup.GET("", h.Product.List)
	productGroup.GET("/:id", h.Product.Get)
	productGroup.PUT("/:id", h.Product.Update)
	productGroup.DELETE("/:id", h.Product.Delete)

	userGroup := s.echo.Group("/users")
	userGroup.POST("", h.User.Create)
}

func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}

func (s *Server) Start(port string) error {
	return s.echo.Start(port)
}
