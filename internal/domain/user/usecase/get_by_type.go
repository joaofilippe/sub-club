package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type GetUsersByTypeUseCase struct {
	repo user.Repository
}

func NewGetUsersByTypeUseCase(repo user.Repository) *GetUsersByTypeUseCase {
	return &GetUsersByTypeUseCase{repo: repo}
}

func (uc *GetUsersByTypeUseCase) Execute(ctx context.Context, userType model.UserType) ([]*model.User, error) {
	return uc.repo.GetByType(ctx, userType)
}
