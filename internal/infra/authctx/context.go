package authctx

import "context"

type contextKey string

const claimsKey contextKey = "auth_claims"

type Claims struct {
	UserID      string
	AccountSlug string
	Role        string
}

func WithClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	c, ok := ctx.Value(claimsKey).(Claims)
	return c, ok
}
