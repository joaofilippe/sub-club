package web

import (
	"github.com/joaofilippe/subclub/internal/application"
	"github.com/joaofilippe/subclub/internal/web/client"
	"github.com/joaofilippe/subclub/internal/web/plan"
	"github.com/joaofilippe/subclub/internal/web/product"
	"github.com/joaofilippe/subclub/internal/web/subscription"
	userhandler "github.com/joaofilippe/subclub/internal/web/user"
)

type Handlers struct {
	Client       *client.ClientHandler
	Plan         *plan.PlanHandler
	Product      *product.ProductHandler
	Subscription *subscription.SubscriptionHandler
	User         *userhandler.UserHandler
}

func NewHandlers(app *application.Application) *Handlers {
	return &Handlers{
		Client:       client.NewClientHandler(app.ClientService),
		Plan:         plan.NewPlanHandler(app.PlanService),
		Product:      product.NewProductHandler(app.ProductService),
		Subscription: subscription.NewSubscriptionHandler(app.SubscriptionService),
		User:         userhandler.NewUserHandler(app.UserService),
	}
}
