package application

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/ent"
	customerrepo "github.com/joaofilippe/subclub/internal/application/repository/customer"
	planrepo "github.com/joaofilippe/subclub/internal/application/repository/plan"
	productrepo "github.com/joaofilippe/subclub/internal/application/repository/product"
	subrepo "github.com/joaofilippe/subclub/internal/application/repository/subscription"
	userrepo "github.com/joaofilippe/subclub/internal/application/repository/user"
	customersvc "github.com/joaofilippe/subclub/internal/application/service/customer"
	plansvc "github.com/joaofilippe/subclub/internal/application/service/plan"
	productsvc "github.com/joaofilippe/subclub/internal/application/service/product"
	subsvc "github.com/joaofilippe/subclub/internal/application/service/subscription"
	usersvc "github.com/joaofilippe/subclub/internal/application/service/user"
	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/infra/database"

	entsql "entgo.io/ent/dialect/sql"
)

type Application struct {
	dbConnection        *database.Connection
	entClient           *ent.Client
	UserService         *usersvc.UserService
	CustomerService     *customersvc.CustomerService
	PlanService         *plansvc.PlanService
	ProductService      *productsvc.ProductService
	SubscriptionService *subsvc.SubscriptionService
}

func New(db *database.Connection) *Application {
	return &Application{dbConnection: db}
}

func (a *Application) Init(ctx context.Context, cfg *config.Config) error {
	drv := entsql.OpenDB("postgres", a.dbConnection.GetDB().DB)
	client := ent.NewClient(ent.Driver(drv))
	return a.initServices(ctx, client, cfg)
}

func (a *Application) initServices(ctx context.Context, client *ent.Client, cfg *config.Config) error {
	a.entClient = client
	if err := a.entClient.Schema.Create(ctx); err != nil {
		log.Printf("Failed creating schema resources: %v", err)
	}

	database.SeedAll(ctx, a.entClient, cfg)

	cr := customerrepo.NewCustomerEntRepository(a.entClient)
	pr := planrepo.NewPlanEntRepository(a.entClient)
	prodr := productrepo.NewProductEntRepository(a.entClient)
	subr := subrepo.NewSubscriptionEntRepository(a.entClient)
	ur := userrepo.NewUserEntRepository(a.entClient)

	a.CustomerService = customersvc.NewCustomerService(cr)
	a.PlanService = plansvc.NewPlanService(pr)
	a.ProductService = productsvc.NewProductService(prodr)
	a.SubscriptionService = subsvc.NewSubscriptionService(subr)
	a.UserService = usersvc.NewUserService(ur)

	return nil
}
