package tenantctx

import (
	"context"

	"github.com/joaofilippe/subclub/ent"
)

type contextKey string

const tenantClientKey contextKey = "tenant_client"

// WithTenantClient stores the tenant-scoped Ent client in the context.
func WithTenantClient(ctx context.Context, client *ent.Client) context.Context {
	return context.WithValue(ctx, tenantClientKey, client)
}

// TenantClientFromContext retrieves the tenant-scoped Ent client from the context.
// Returns nil if no tenant client is set (e.g. admin-only routes).
func TenantClientFromContext(ctx context.Context) *ent.Client {
	c, _ := ctx.Value(tenantClientKey).(*ent.Client)
	return c
}
