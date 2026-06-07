package tenantuser

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
)

type Repository interface {
	Create(ctx context.Context, user *model.TenantUser, passwordHash string) error
	GetByID(ctx context.Context, id string) (*model.TenantUser, error)
	Update(ctx context.Context, user *model.TenantUser) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.TenantUser, error)
}
