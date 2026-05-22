package web

import (
	"github.com/joaofilippe/subclub/internal/application"
	"github.com/joaofilippe/subclub/internal/web/account"
	"github.com/joaofilippe/subclub/internal/web/accountplan"
	"github.com/joaofilippe/subclub/internal/web/customer"
	"github.com/joaofilippe/subclub/internal/web/plan"
	"github.com/joaofilippe/subclub/internal/web/product"
	"github.com/joaofilippe/subclub/internal/web/subscription"
	userhandler "github.com/joaofilippe/subclub/internal/web/user"
)

type Handlers struct {
	Account      *account.AccountHandler
	AccountPlan  *accountplan.AccountPlanHandler
	Customer     *customer.CustomerHandler
	Plan         *plan.PlanHandler
	Product      *product.ProductHandler
	Subscription *subscription.SubscriptionHandler
	User         *userhandler.UserHandler
}

func NewHandlers(app *application.Application) *Handlers {
	return &Handlers{
		Account:      account.NewAccountHandler(app.AccountService),
		AccountPlan:  accountplan.NewAccountPlanHandler(app.AccountPlanService),
		Customer:     customer.NewCustomerHandler(app.CustomerService),
		Plan:         plan.NewPlanHandler(app.PlanService),
		Product:      product.NewProductHandler(app.ProductService),
		Subscription: subscription.NewSubscriptionHandler(app.SubscriptionService),
		User:         userhandler.NewUserHandler(app.UserService),
	}
}
