package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/subclub/internal/infra/authctx"
	"github.com/joaofilippe/subclub/internal/web/common"
)

// RequireTenantAdminMiddleware must be chained after AuthMiddleware.
// It rejects requests from tenant users whose role is not "admin".
func RequireTenantAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, ok := authctx.ClaimsFromContext(c.Request().Context())
			if !ok || claims.Role != "admin" {
				return c.JSON(http.StatusForbidden, common.Response{Message: "admin access required"})
			}
			return next(c)
		}
	}
}
