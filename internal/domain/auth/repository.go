package auth

import (
	"context"

	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

// TenantRepository abstracts per-tenant user lookups needed for authentication.
// Implemented by the infra layer using TenantClientManager.
type TenantRepository interface {
	FindUserByEmailAndSlug(ctx context.Context, slug, email string) (*authmodel.TenantAuthUser, error)
	FindUserByUsernameAndSlug(ctx context.Context, slug, username string) (*authmodel.TenantAuthUser, error)
	UserExistsByEmailAndSlug(ctx context.Context, slug, email string) (bool, error)
}
