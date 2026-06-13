package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	authusecase "github.com/joaofilippe/subclub/internal/domain/auth/usecase"
	"github.com/joaofilippe/subclub/internal/infra/authctx"
	"github.com/joaofilippe/subclub/internal/web/common"
)

// RequireAdminMiddleware validates the JWT and rejects requests whose role
// claim is not "admin". It does not resolve a tenant client — admin routes
// operate on the global schema.
func RequireAdminMiddleware(jwtSecret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, common.Response{Message: "missing or malformed Authorization header"})
			}
			raw := strings.TrimPrefix(header, "Bearer ")

			claims := &authusecase.Claims{}
			token, err := jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, common.Response{Message: "invalid or expired token"})
			}

			if claims.Role != "admin" {
				return c.JSON(http.StatusForbidden, common.Response{Message: "admin access required"})
			}

			ctx := authctx.WithClaims(c.Request().Context(), authctx.Claims{
				UserID:      claims.UserID,
				AccountSlug: claims.AccountSlug,
				Role:        claims.Role,
			})
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
