package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
)

type DeleteTenantUserUseCase struct {
	repo tenantuser.Repository
}

func NewDeleteTenantUserUseCase(repo tenantuser.Repository) *DeleteTenantUserUseCase {
	return &DeleteTenantUserUseCase{repo: repo}
}

func (uc *DeleteTenantUserUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
