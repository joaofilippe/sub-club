package server_test

import (
	"testing"

	"github.com/joaofilippe/subclub/internal/application"
	accountsvc "github.com/joaofilippe/subclub/internal/application/service/account"
	accountplansvc "github.com/joaofilippe/subclub/internal/application/service/accountplan"
	authsvc "github.com/joaofilippe/subclub/internal/application/service/auth"
	customersvc "github.com/joaofilippe/subclub/internal/application/service/customer"
	plansvc "github.com/joaofilippe/subclub/internal/application/service/plan"
	productsvc "github.com/joaofilippe/subclub/internal/application/service/product"
	subsvc "github.com/joaofilippe/subclub/internal/application/service/subscription"
	usersvc "github.com/joaofilippe/subclub/internal/application/service/user"
	"github.com/joaofilippe/subclub/internal/infra/server"
	"github.com/joaofilippe/subclub/internal/web"
)

func stubApp() *application.Application {
	return &application.Application{
		AuthService:         authsvc.NewAuthService(nil, nil, nil),
		AccountService:      accountsvc.NewAccountService(nil, nil),
		AccountPlanService:  accountplansvc.NewAccountPlanService(nil),
		UserService:         usersvc.NewUserService(nil),
		CustomerService:     customersvc.NewCustomerService(nil),
		PlanService:         plansvc.NewPlanService(nil),
		ProductService:      productsvc.NewProductService(nil),
		SubscriptionService: subsvc.NewSubscriptionService(nil),
	}
}

func TestServer_AllRoutesPresent(t *testing.T) {
	srv := server.NewServer(web.NewHandlers(stubApp()), nil, []byte("test-secret"))

	registered := map[string]bool{}
	for _, r := range srv.GetEcho().Routes() {
		registered[r.Method+":"+r.Path] = true
	}

	want := []struct{ method, path string }{
		// public
		{"POST", "/auth/login"},
		// admin-only
		{"POST", "/accounts"},
		{"GET", "/accounts"},
		{"GET", "/accounts/:id"},
		{"PUT", "/accounts/:id"},
		{"DELETE", "/accounts/:id"},
		{"POST", "/account-plans"},
		{"GET", "/account-plans"},
		{"GET", "/account-plans/:id"},
		{"PUT", "/account-plans/:id"},
		{"DELETE", "/account-plans/:id"},
		{"POST", "/users"},
		// tenant-scoped
		{"POST", "/customers"},
		{"GET", "/customers"},
		{"GET", "/customers/:id"},
		{"PUT", "/customers/:id"},
		{"DELETE", "/customers/:id"},
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
	}

	for _, w := range want {
		if !registered[w.method+":"+w.path] {
			t.Errorf("route %s %s not registered", w.method, w.path)
		}
	}
}
