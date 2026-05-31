package server

import (
	_ "github.com/joaofilippe/subclub/docs"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/middleware"
	"github.com/joaofilippe/subclub/internal/web"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type router struct {
	echo          *echo.Echo
	handlers      *web.Handlers
	tenantManager *database.TenantClientManager
	jwtSecret     []byte
}

func newRouter(e *echo.Echo, handlers *web.Handlers, tenantManager *database.TenantClientManager, jwtSecret []byte) *router {
	return &router{echo: e, handlers: handlers, tenantManager: tenantManager, jwtSecret: jwtSecret}
}

func (r *router) registerRoutes() {
	r.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	r.echo.POST("/auth/login", r.handlers.Auth.Login)

	adminMW := middleware.RequireAdminMiddleware(r.jwtSecret)

	accountPlanGroup := r.echo.Group("/account-plans", adminMW)
	accountPlanGroup.POST("", r.handlers.AccountPlan.Create)
	accountPlanGroup.GET("", r.handlers.AccountPlan.List)
	accountPlanGroup.GET("/:id", r.handlers.AccountPlan.Get)
	accountPlanGroup.PUT("/:id", r.handlers.AccountPlan.Update)
	accountPlanGroup.DELETE("/:id", r.handlers.AccountPlan.Delete)

	moduleGroup := r.echo.Group("/modules", adminMW)
	moduleGroup.POST("", r.handlers.Module.Create)
	moduleGroup.GET("", r.handlers.Module.List)
	moduleGroup.GET("/:id", r.handlers.Module.Get)
	moduleGroup.PUT("/:id", r.handlers.Module.Update)
	moduleGroup.DELETE("/:id", r.handlers.Module.Delete)

	accountGroup := r.echo.Group("/accounts", adminMW)
	accountGroup.POST("", r.handlers.Account.Create)
	accountGroup.GET("", r.handlers.Account.List)
	accountGroup.GET("/:id", r.handlers.Account.Get)
	accountGroup.PUT("/:id", r.handlers.Account.Update)
	accountGroup.DELETE("/:id", r.handlers.Account.Delete)

	userGroup2 := r.echo.Group("/users", adminMW)
	userGroup2.POST("", r.handlers.User.Create)

	authMW := middleware.AuthMiddleware(r.tenantManager, r.jwtSecret)

	customerGroup := r.echo.Group("/customers", authMW)
	customerGroup.POST("", r.handlers.Customer.Create)
	customerGroup.GET("", r.handlers.Customer.List)
	customerGroup.GET("/:id", r.handlers.Customer.Get)
	customerGroup.PUT("/:id", r.handlers.Customer.Update)
	customerGroup.DELETE("/:id", r.handlers.Customer.Delete)

	planGroup := r.echo.Group("/plans", authMW)
	planGroup.POST("", r.handlers.Plan.Create)
	planGroup.GET("", r.handlers.Plan.List)
	planGroup.GET("/:id", r.handlers.Plan.Get)
	planGroup.PUT("/:id", r.handlers.Plan.Update)
	planGroup.DELETE("/:id", r.handlers.Plan.Delete)

	subGroup := r.echo.Group("/subscriptions", authMW)
	subGroup.POST("", r.handlers.Subscription.Create)
	subGroup.GET("", r.handlers.Subscription.List)
	subGroup.GET("/:id", r.handlers.Subscription.Get)
	subGroup.PUT("/:id", r.handlers.Subscription.Update)
	subGroup.DELETE("/:id", r.handlers.Subscription.Delete)

	productGroup := r.echo.Group("/products", authMW)
	productGroup.POST("", r.handlers.Product.Create)
	productGroup.GET("", r.handlers.Product.List)
	productGroup.GET("/:id", r.handlers.Product.Get)
	productGroup.PUT("/:id", r.handlers.Product.Update)
	productGroup.DELETE("/:id", r.handlers.Product.Delete)

}
