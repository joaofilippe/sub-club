package server_test

import (
	"testing"

	"github.com/joaofilippe/subclub/internal/application"
	clientsvc "github.com/joaofilippe/subclub/internal/application/service/client"
	plansvc "github.com/joaofilippe/subclub/internal/application/service/plan"
	productsvc "github.com/joaofilippe/subclub/internal/application/service/product"
	subsvc "github.com/joaofilippe/subclub/internal/application/service/subscription"
	usersvc "github.com/joaofilippe/subclub/internal/application/service/user"
	"github.com/joaofilippe/subclub/internal/infra/server"
	"github.com/joaofilippe/subclub/internal/web"
)

func stubApp() *application.Application {
	return &application.Application{
		UserService:         usersvc.NewUserService(nil),
		ClientService:       clientsvc.NewClientService(nil),
		PlanService:         plansvc.NewPlanService(nil),
		ProductService:      productsvc.NewProductService(nil),
		SubscriptionService: subsvc.NewSubscriptionService(nil),
	}
}

func TestServer_AllRoutesPresent(t *testing.T) {
	srv := server.NewServer(web.NewHandlers(stubApp()))

	registered := map[string]bool{}
	for _, r := range srv.GetEcho().Routes() {
		registered[r.Method+":"+r.Path] = true
	}

	want := []struct{ method, path string }{
		{"POST", "/clients"},
		{"GET", "/clients"},
		{"GET", "/clients/:id"},
		{"PUT", "/clients/:id"},
		{"DELETE", "/clients/:id"},
		{"POST", "/plans"},
		{"GET", "/plans"},
		{"GET", "/plans/:id"},
		{"PUT", "/plans/:id"},
		{"DELETE", "/plans/:id"},
		{"POST", "/subscriptions"},
		{"GET", "/subscriptions"},
		{"GET", "/subscriptions/:id"},
		{"PUT", "/subscriptions/:id"},
		{"DELETE", "/subscriptions/:id"},
		{"POST", "/products"},
		{"GET", "/products"},
		{"GET", "/products/:id"},
		{"PUT", "/products/:id"},
		{"DELETE", "/products/:id"},
		{"POST", "/users"},
	}

	for _, w := range want {
		if !registered[w.method+":"+w.path] {
			t.Errorf("route %s %s not registered", w.method, w.path)
		}
	}
}
