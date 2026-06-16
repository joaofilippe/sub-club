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
	srv := server.NewServer(web.NewHandlers(stubApp()), nil, []byte("test-secret"), false)

	registered := map[string]bool{}
	for _, r := range srv.GetEcho().Routes() {
		registered[r.Method+":"+r.Path] = true
	}

	want := []struct{ method, path string }{
		// public
		{"POST", "/v1/auth/login"},
		{"POST", "/v1/auth/login/username"},
		{"POST", "/v1/auth/lookup"},
		// admin-only
		{"POST", "/v1/accounts"},
		{"GET", "/v1/accounts"},
		{"GET", "/v1/accounts/:id"},
		{"PUT", "/v1/accounts/:id"},
		{"DELETE", "/v1/accounts/:id"},
		{"POST", "/v1/account-plans"},
		{"GET", "/v1/account-plans"},
		{"GET", "/v1/account-plans/:id"},
		{"PUT", "/v1/account-plans/:id"},
		{"DELETE", "/v1/account-plans/:id"},
		{"POST", "/v1/users"},
		// tenant-scoped
		{"POST", "/v1/customers"},
		{"GET", "/v1/customers"},
		{"GET", "/v1/customers/:id"},
		{"PUT", "/v1/customers/:id"},
		{"DELETE", "/v1/customers/:id"},
		{"POST", "/v1/plans"},
		{"GET", "/v1/plans"},
		{"GET", "/v1/plans/:id"},
		{"PUT", "/v1/plans/:id"},
		{"DELETE", "/v1/plans/:id"},
		{"POST", "/v1/subscriptions"},
		{"GET", "/v1/subscriptions"},
		{"GET", "/v1/subscriptions/:id"},
		{"PUT", "/v1/subscriptions/:id"},
		{"DELETE", "/v1/subscriptions/:id"},
		{"POST", "/v1/products"},
		{"GET", "/v1/products"},
		{"GET", "/v1/products/:id"},
		{"PUT", "/v1/products/:id"},
		{"DELETE", "/v1/products/:id"},
	}

	for _, w := range want {
		if !registered[w.method+":"+w.path] {
			t.Errorf("route %s %s not registered", w.method, w.path)
		}
	}
}
