package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
)

type GetTenantUserByIDUseCase struct {
	repo tenantuser.Repository
}

func NewGetTenantUserByIDUseCase(repo tenantuser.Repository) *GetTenantUserByIDUseCase {
	return &GetTenantUserByIDUseCase{repo: repo}
}

func (uc *GetTenantUserByIDUseCase) Execute(ctx context.Context, id string) (*model.TenantUser, error) {
	return uc.repo.GetByID(ctx, id)
}
