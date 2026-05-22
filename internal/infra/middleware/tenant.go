package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/tenantctx"
	"github.com/joaofilippe/subclub/internal/web/common"
)

// TenantMiddleware resolves the account slug from the X-Account-Slug request
// header, retrieves (or lazily creates) the tenant-scoped Ent client, and
// stores it in the request context.
//
// Phase 3 will replace the header with a JWT claim.
// Routes that don't require a tenant context (e.g. /accounts, /account-plans)
// should be registered outside the tenant-protected group.
func TenantMiddleware(manager *database.TenantClientManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			slug := c.Request().Header.Get("X-Account-Slug")
			if slug == "" {
				return c.JSON(http.StatusUnauthorized, common.Response{Message: "missing X-Account-Slug header"})
			}

			client, err := manager.GetOrCreate(slug)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.Response{Message: "failed to resolve tenant"})
			}

			ctx := tenantctx.WithTenantClient(c.Request().Context(), client)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
