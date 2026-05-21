package router

import (
	_ "github.com/joaofilippe/subclub/docs"
	"github.com/joaofilippe/subclub/internal/infra/server"
	"github.com/joaofilippe/subclub/internal/web"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server   *server.Server
	handlers *web.Handlers
}

func NewRouter(server *server.Server, handlers *web.Handlers) *Router {
	return &Router{
		server:   server,
		handlers: handlers,
	}
}

func (r *Router) RegisterRoutes() {
	echo := r.server.GetEcho()

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	clientGroup := echo.Group("/clients")
	clientGroup.POST("", r.handlers.Client.Create)
	clientGroup.GET("", r.handlers.Client.List)
	clientGroup.GET("/:id", r.handlers.Client.Get)
	clientGroup.PUT("/:id", r.handlers.Client.Update)
	clientGroup.DELETE("/:id", r.handlers.Client.Delete)

	planGroup := echo.Group("/plans")
	planGroup.POST("", r.handlers.Plan.Create)
	planGroup.GET("", r.handlers.Plan.List)
	planGroup.GET("/:id", r.handlers.Plan.Get)
	planGroup.PUT("/:id", r.handlers.Plan.Update)
	planGroup.DELETE("/:id", r.handlers.Plan.Delete)

	subGroup := echo.Group("/subscriptions")
	subGroup.POST("", r.handlers.Subscription.Create)
	subGroup.GET("", r.handlers.Subscription.List)
	subGroup.GET("/:id", r.handlers.Subscription.Get)
	subGroup.PUT("/:id", r.handlers.Subscription.Update)
	subGroup.DELETE("/:id", r.handlers.Subscription.Delete)

	productGroup := echo.Group("/products")
	productGroup.POST("", r.handlers.Product.Create)
	productGroup.GET("", r.handlers.Product.List)
	productGroup.GET("/:id", r.handlers.Product.Get)
	productGroup.PUT("/:id", r.handlers.Product.Update)
	productGroup.DELETE("/:id", r.handlers.Product.Delete)

	userGroup := echo.Group("/users")
	userGroup.POST("", r.handlers.User.Create)
}

func (r *Router) Start(port string) error {
	return r.server.Start(port)
}
