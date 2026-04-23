package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type GetUsersByRoleUseCase struct {
	repo user.Repository
}

func NewGetUsersByRoleUseCase(repo user.Repository) *GetUsersByRoleUseCase {
	return &GetUsersByRoleUseCase{repo: repo}
}

func (uc *GetUsersByRoleUseCase) Execute(ctx context.Context, role model.UserRole) ([]*model.User, error) {
	return uc.repo.GetByRole(ctx, role)
}
