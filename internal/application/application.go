package application

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/ent"
	accountrepo "github.com/joaofilippe/subclub/internal/application/repository/account"
	accountplanrepo "github.com/joaofilippe/subclub/internal/application/repository/accountplan"
	customerrepo "github.com/joaofilippe/subclub/internal/application/repository/customer"
	planrepo "github.com/joaofilippe/subclub/internal/application/repository/plan"
	productrepo "github.com/joaofilippe/subclub/internal/application/repository/product"
	subrepo "github.com/joaofilippe/subclub/internal/application/repository/subscription"
	userrepo "github.com/joaofilippe/subclub/internal/application/repository/user"
	accountsvc "github.com/joaofilippe/subclub/internal/application/service/account"
	accountplansvc "github.com/joaofilippe/subclub/internal/application/service/accountplan"
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
	TenantManager       *database.TenantClientManager
	AccountService      *accountsvc.AccountService
	AccountPlanService  *accountplansvc.AccountPlanService
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

	a.TenantManager = database.NewTenantClientManager(cfg.DatabaseURL, a.dbConnection.GetDB().DB)

	database.SeedAll(ctx, a.entClient, cfg)

	a.AccountService = accountsvc.NewAccountService(accountrepo.NewAccountEntRepository(a.entClient), a.TenantManager)
	a.AccountPlanService = accountplansvc.NewAccountPlanService(accountplanrepo.NewAccountPlanEntRepository(a.entClient))
	a.CustomerService = customersvc.NewCustomerService(customerrepo.NewCustomerEntRepository(a.entClient))
	a.PlanService = plansvc.NewPlanService(planrepo.NewPlanEntRepository(a.entClient))
	a.ProductService = productsvc.NewProductService(productrepo.NewProductEntRepository(a.entClient))
	a.SubscriptionService = subsvc.NewSubscriptionService(subrepo.NewSubscriptionEntRepository(a.entClient))
	a.UserService = usersvc.NewUserService(userrepo.NewUserEntRepository(a.entClient))

	return nil
}
