package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
)

type ListTenantUsersUseCase struct {
	repo tenantuser.Repository
}

func NewListTenantUsersUseCase(repo tenantuser.Repository) *ListTenantUsersUseCase {
	return &ListTenantUsersUseCase{repo: repo}
}

func (uc *ListTenantUsersUseCase) Execute(ctx context.Context) ([]*model.TenantUser, error) {
	return uc.repo.List(ctx)
}
