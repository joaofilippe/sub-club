package api

import (
	_ "github.com/joaofilippe/subclub/docs"
	"github.com/joaofilippe/subclub/internal/adapter/api/client"
	"github.com/joaofilippe/subclub/internal/adapter/api/plan"
	"github.com/joaofilippe/subclub/internal/adapter/api/product"
	"github.com/joaofilippe/subclub/internal/adapter/api/subscription"
	userhandler "github.com/joaofilippe/subclub/internal/adapter/api/user"
	"github.com/joaofilippe/subclub/internal/application"
	"github.com/joaofilippe/subclub/internal/infra/server"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type API struct {
	server      *server.Server
	application *application.Application
}

func New(server *server.Server, app *application.Application) *API {
	return &API{
		server:      server,
		application: app,
	}
}

func (a *API) RegisterRoutes() {
	clientHandler := client.NewClientHandler(a.application.ClientService)
	planHandler := plan.NewPlanHandler(a.application.PlanService)
	productHandler := product.NewProductHandler(a.application.ProductService)
	subHandler := subscription.NewSubscriptionHandler(a.application.SubscriptionService)
	userHandler := userhandler.NewUserHandler(a.application.UserService)

	router := a.server.GetEcho()

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	clientGroup := router.Group("/clients")
	clientGroup.POST("", clientHandler.Create)
	clientGroup.GET("", clientHandler.List)
	clientGroup.GET("/:id", clientHandler.Get)
	clientGroup.PUT("/:id", clientHandler.Update)
	clientGroup.DELETE("/:id", clientHandler.Delete)

	planGroup := router.Group("/plans")
	planGroup.POST("", planHandler.Create)
	planGroup.GET("", planHandler.List)
	planGroup.GET("/:id", planHandler.Get)
	planGroup.PUT("/:id", planHandler.Update)
	planGroup.DELETE("/:id", planHandler.Delete)

	subGroup := router.Group("/subscriptions")
	subGroup.POST("", subHandler.Create)
	subGroup.GET("", subHandler.List)
	subGroup.GET("/:id", subHandler.Get)
	subGroup.PUT("/:id", subHandler.Update)
	subGroup.DELETE("/:id", subHandler.Delete)

	productGroup := router.Group("/products")
	productGroup.POST("", productHandler.Create)
	productGroup.GET("", productHandler.List)
	productGroup.GET("/:id", productHandler.Get)
	productGroup.PUT("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)

	userGroup := router.Group("/users")
	userGroup.POST("", userHandler.Create)
}

func (a *API) Start(port string) error {
	return a.server.Start(port)
}
