package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type GetUserByEmailUseCase struct {
	repo user.Repository
}

func NewGetUserByEmailUseCase(repo user.Repository) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{repo: repo}
}

func (uc *GetUserByEmailUseCase) Execute(ctx context.Context, email string) (*model.User, error) {
	return uc.repo.GetByEmail(ctx, email)
}
