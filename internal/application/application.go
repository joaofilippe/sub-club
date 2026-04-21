package application

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/internal/adapter/api/client"
	"github.com/joaofilippe/subclub/internal/adapter/api/plan"
	"github.com/joaofilippe/subclub/internal/adapter/api/product"
	"github.com/joaofilippe/subclub/internal/adapter/api/subscription"
	"github.com/joaofilippe/subclub/internal/adapter/repository"
	services "github.com/joaofilippe/subclub/internal/adapter/service"
	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/server"

	entsql "entgo.io/ent/dialect/sql"
)

// Application orquestra os componentes principais do sistema
type Application struct {
	dbConnection *database.Connection
	server       *server.Server
	entClient    *ent.Client
	userService  user.Service // Add generic domain user.Service
}

// New inicializa uma nova Application com suas dependências globais
func New(server *server.Server, db *database.Connection) *Application {
	return &Application{
		dbConnection: db,
		server:       server,
	}
}

func (a *Application) InitServices() error {
	ctx := context.Background()

	// 1. Initialize Ent Client
	drv := entsql.OpenDB("postgres", a.dbConnection.GetDB().DB)
	a.entClient = ent.NewClient(ent.Driver(drv))
	if err := a.entClient.Schema.Create(ctx); err != nil {
		log.Printf("Failed creating schema resources: %v", err)
	}

	// 2. Repositories
	clientRepo := repository.NewClientEntRepository(a.entClient)
	planRepo := repository.NewPlanEntRepository(a.entClient)
	productRepo := repository.NewProductEntRepository(a.entClient)
	subRepo := repository.NewSubscriptionEntRepository(a.entClient)
	userRepo := repository.NewUserEntRepository(a.entClient)

	// 3. Services
	clientService := services.NewClientService(clientRepo)
	planService := services.NewPlanService(planRepo)
	productService := services.NewProductService(productRepo)
	subService := services.NewSubscriptionService(subRepo)
	userService := services.NewUserService(userRepo)
	a.userService = userService // Stored in application struct just in case

	// 4. Handlers
	clientHandler := client.NewClientHandler(clientService)
	planHandler := plan.NewPlanHandler(planService)
	productHandler := product.NewProductHandler(productService)
	subHandler := subscription.NewSubscriptionHandler(subService)

	// 5. Register Routes
	router := a.server.GetEcho()

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

	return nil
}

func (a *Application) Start(port string) error {
	return a.server.Start(port)
}
