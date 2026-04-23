package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type GetUserByIDUseCase struct {
	repo user.Repository
}

func NewGetUserByIDUseCase(repo user.Repository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: repo}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id string) (*model.User, error) {
	return uc.repo.GetByID(ctx, id)
}
