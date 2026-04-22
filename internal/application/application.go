package application

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/internal/adapter/repository"
	services "github.com/joaofilippe/subclub/internal/adapter/service"
	"github.com/joaofilippe/subclub/internal/infra/database"

	entsql "entgo.io/ent/dialect/sql"
)

type Application struct {
	dbConnection        *database.Connection
	entClient           *ent.Client
	UserService         *services.UserService
	ClientService       *services.ClientService
	PlanService         *services.PlanService
	ProductService      *services.ProductService
	SubscriptionService *services.SubscriptionService
}

func New(db *database.Connection) *Application {
	return &Application{dbConnection: db}
}

func (a *Application) Init(ctx context.Context) error {
	drv := entsql.OpenDB("postgres", a.dbConnection.GetDB().DB)
	a.entClient = ent.NewClient(ent.Driver(drv))
	if err := a.entClient.Schema.Create(ctx); err != nil {
		log.Printf("Failed creating schema resources: %v", err)
	}

	database.SeedAll(ctx, a.entClient)

	clientRepo := repository.NewClientEntRepository(a.entClient)
	planRepo := repository.NewPlanEntRepository(a.entClient)
	productRepo := repository.NewProductEntRepository(a.entClient)
	subRepo := repository.NewSubscriptionEntRepository(a.entClient)
	userRepo := repository.NewUserEntRepository(a.entClient)

	a.ClientService = services.NewClientService(clientRepo)
	a.PlanService = services.NewPlanService(planRepo)
	a.ProductService = services.NewProductService(productRepo)
	a.SubscriptionService = services.NewSubscriptionService(subRepo)
	a.UserService = services.NewUserService(userRepo)

	return nil
}
