package web

import (
	"github.com/joaofilippe/subclub/internal/application"
	authhandler "github.com/joaofilippe/subclub/internal/web/auth"
	"github.com/joaofilippe/subclub/internal/web/account"
	"github.com/joaofilippe/subclub/internal/web/accountplan"
	"github.com/joaofilippe/subclub/internal/web/customer"
	"github.com/joaofilippe/subclub/internal/web/module"
	"github.com/joaofilippe/subclub/internal/web/plan"
	"github.com/joaofilippe/subclub/internal/web/product"
	"github.com/joaofilippe/subclub/internal/web/subscription"
	userhandler "github.com/joaofilippe/subclub/internal/web/user"
)

type Handlers struct {
	Auth         *authhandler.AuthHandler
	Account      *account.AccountHandler
	AccountPlan  *accountplan.AccountPlanHandler
	Module       *module.ModuleHandler
	Customer     *customer.CustomerHandler
	Plan         *plan.PlanHandler
	Product      *product.ProductHandler
	Subscription *subscription.SubscriptionHandler
	User         *userhandler.UserHandler
}

func NewHandlers(app *application.Application) *Handlers {
	return &Handlers{
		Auth:         authhandler.NewAuthHandler(app.AuthService),
		Account:      account.NewAccountHandler(app.AccountService),
		AccountPlan:  accountplan.NewAccountPlanHandler(app.AccountPlanService),
		Module:       module.NewModuleHandler(app.ModuleService),
		Customer:     customer.NewCustomerHandler(app.CustomerService),
		Plan:         plan.NewPlanHandler(app.PlanService),
		Product:      product.NewProductHandler(app.ProductService),
		Subscription: subscription.NewSubscriptionHandler(app.SubscriptionService),
		User:         userhandler.NewUserHandler(app.UserService),
	}
}
