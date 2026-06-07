package tenantuser

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateTenantUserInput) (*model.TenantUser, error)
	GetByID(ctx context.Context, id string) (*model.TenantUser, error)
	Update(ctx context.Context, input model.UpdateTenantUserInput) (*model.TenantUser, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.TenantUser, error)
}
