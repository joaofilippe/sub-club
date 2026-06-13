package application

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/ent"
	accountrepo "github.com/joaofilippe/subclub/internal/application/repository/account"
	accountdomainrepo "github.com/joaofilippe/subclub/internal/application/repository/accountdomain"
	accountplanrepo "github.com/joaofilippe/subclub/internal/application/repository/accountplan"
	customerrepo "github.com/joaofilippe/subclub/internal/application/repository/customer"
	modulerepo "github.com/joaofilippe/subclub/internal/application/repository/module"
	planrepo "github.com/joaofilippe/subclub/internal/application/repository/plan"
	productrepo "github.com/joaofilippe/subclub/internal/application/repository/product"
	subrepo "github.com/joaofilippe/subclub/internal/application/repository/subscription"
	tenantuserrepo "github.com/joaofilippe/subclub/internal/application/repository/tenantuser"
	userrepo "github.com/joaofilippe/subclub/internal/application/repository/user"
	accountsvc "github.com/joaofilippe/subclub/internal/application/service/account"
	accountplansvc "github.com/joaofilippe/subclub/internal/application/service/accountplan"
	authsvc "github.com/joaofilippe/subclub/internal/application/service/auth"
	customersvc "github.com/joaofilippe/subclub/internal/application/service/customer"
	modulesvc "github.com/joaofilippe/subclub/internal/application/service/module"
	plansvc "github.com/joaofilippe/subclub/internal/application/service/plan"
	productsvc "github.com/joaofilippe/subclub/internal/application/service/product"
	subsvc "github.com/joaofilippe/subclub/internal/application/service/subscription"
	tenantusersvc "github.com/joaofilippe/subclub/internal/application/service/tenantuser"
	usersvc "github.com/joaofilippe/subclub/internal/application/service/user"
	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/infra/database"

	entsql "entgo.io/ent/dialect/sql"
)

type Application struct {
	dbConnection        *database.Connection
	entClient           *ent.Client
	TenantManager       *database.TenantClientManager
	AuthService         *authsvc.AuthService
	AccountService      *accountsvc.AccountService
	AccountPlanService  *accountplansvc.AccountPlanService
	ModuleService       *modulesvc.ModuleService
	UserService         *usersvc.UserService
	TenantUserService   *tenantusersvc.TenantUserService
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
	if err := database.MigratePublic(ctx, a.entClient); err != nil {
		log.Printf("Failed migrating public schema: %v", err)
	}

	a.TenantManager = database.NewTenantClientManager(cfg.DatabaseURL, a.dbConnection.GetDB().DB, a.entClient, cfg.IsDevelopment())

	database.SeedAll(ctx, a.entClient, cfg, a.TenantManager)

	userRepo := userrepo.NewUserEntRepository(a.entClient)
	accountRepo := accountrepo.NewAccountEntRepository(a.entClient)
	accountDomainRepo := accountdomainrepo.NewAccountDomainEntRepository(a.entClient)

	a.AuthService = authsvc.NewAuthService(userRepo, accountRepo, accountDomainRepo, a.TenantManager, []byte(cfg.JWTSecret))
	a.AccountService = accountsvc.NewAccountService(accountRepo, a.TenantManager)
	a.AccountPlanService = accountplansvc.NewAccountPlanService(accountplanrepo.NewAccountPlanEntRepository(a.entClient))
	a.ModuleService = modulesvc.NewModuleService(modulerepo.NewModuleEntRepository(a.entClient))
	a.CustomerService = customersvc.NewCustomerService(customerrepo.NewCustomerEntRepository(a.entClient))
	a.PlanService = plansvc.NewPlanService(planrepo.NewPlanEntRepository(a.entClient))
	a.ProductService = productsvc.NewProductService(productrepo.NewProductEntRepository(a.entClient))
	a.SubscriptionService = subsvc.NewSubscriptionService(subrepo.NewSubscriptionEntRepository(a.entClient))
	a.UserService = usersvc.NewUserService(userRepo)
	a.TenantUserService = tenantusersvc.NewTenantUserService(tenantuserrepo.NewTenantUserEntRepository(a.entClient))

	return nil
}
