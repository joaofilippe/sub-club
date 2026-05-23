package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	authsvc "github.com/joaofilippe/subclub/internal/application/service/auth"
	"github.com/joaofilippe/subclub/internal/infra/authctx"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/tenantctx"
	"github.com/joaofilippe/subclub/internal/web/common"
)

// AuthMiddleware validates the JWT from the Authorization header, resolves the
// tenant client from the account_slug claim, and injects it into the request context.
func AuthMiddleware(manager *database.TenantClientManager, jwtSecret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, common.Response{Message: "missing or malformed Authorization header"})
			}
			raw := strings.TrimPrefix(header, "Bearer ")

			claims := &authsvc.Claims{}
			token, err := jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, common.Response{Message: "invalid or expired token"})
			}

			if claims.AccountSlug == "" {
				return c.JSON(http.StatusForbidden, common.Response{Message: "token does not belong to any account"})
			}

			if err := manager.IsAccountAccessible(c.Request().Context(), claims.AccountSlug); err != nil {
				if errors.Is(err, database.ErrAccountSuspended) {
					return c.JSON(http.StatusForbidden, common.Response{Message: "account is suspended or cancelled"})
				}
				return c.JSON(http.StatusForbidden, common.Response{Message: "account not found"})
			}

			client, err := manager.GetOrCreate(claims.AccountSlug)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.Response{Message: "failed to resolve tenant"})
			}

			ctx := tenantctx.WithTenantClient(c.Request().Context(), client)
			ctx = authctx.WithClaims(ctx, authctx.Claims{
				UserID:      claims.UserID,
				AccountSlug: claims.AccountSlug,
				Role:        claims.Role,
			})
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
