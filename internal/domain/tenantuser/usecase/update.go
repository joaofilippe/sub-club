package usecase

import (
	"context"
	"time"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
)

type UpdateTenantUserUseCase struct {
	repo tenantuser.Repository
}

func NewUpdateTenantUserUseCase(repo tenantuser.Repository) *UpdateTenantUserUseCase {
	return &UpdateTenantUserUseCase{repo: repo}
}

func (uc *UpdateTenantUserUseCase) Execute(ctx context.Context, input model.UpdateTenantUserInput) (*model.TenantUser, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Name = input.Name
	existing.Role = input.Role
	existing.UpdatedAt = time.Now()

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
