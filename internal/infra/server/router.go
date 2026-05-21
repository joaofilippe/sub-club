package server

import (
	_ "github.com/joaofilippe/subclub/docs"
	"github.com/joaofilippe/subclub/internal/web"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type router struct {
	echo     *echo.Echo
	handlers *web.Handlers
}

func newRouter(e *echo.Echo, handlers *web.Handlers) *router {
	return &router{echo: e, handlers: handlers}
}

func (r *router) registerRoutes() {
	r.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	clientGroup := r.echo.Group("/clients")
	clientGroup.POST("", r.handlers.Client.Create)
	clientGroup.GET("", r.handlers.Client.List)
	clientGroup.GET("/:id", r.handlers.Client.Get)
	clientGroup.PUT("/:id", r.handlers.Client.Update)
	clientGroup.DELETE("/:id", r.handlers.Client.Delete)

	planGroup := r.echo.Group("/plans")
	planGroup.POST("", r.handlers.Plan.Create)
	planGroup.GET("", r.handlers.Plan.List)
	planGroup.GET("/:id", r.handlers.Plan.Get)
	planGroup.PUT("/:id", r.handlers.Plan.Update)
	planGroup.DELETE("/:id", r.handlers.Plan.Delete)

	subGroup := r.echo.Group("/subscriptions")
	subGroup.POST("", r.handlers.Subscription.Create)
	subGroup.GET("", r.handlers.Subscription.List)
	subGroup.GET("/:id", r.handlers.Subscription.Get)
	subGroup.PUT("/:id", r.handlers.Subscription.Update)
	subGroup.DELETE("/:id", r.handlers.Subscription.Delete)

	productGroup := r.echo.Group("/products")
	productGroup.POST("", r.handlers.Product.Create)
	productGroup.GET("", r.handlers.Product.List)
	productGroup.GET("/:id", r.handlers.Product.Get)
	productGroup.PUT("/:id", r.handlers.Product.Update)
	productGroup.DELETE("/:id", r.handlers.Product.Delete)

	userGroup := r.echo.Group("/users")
	userGroup.POST("", r.handlers.User.Create)
}
