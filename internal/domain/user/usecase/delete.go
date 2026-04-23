package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
)

type DeleteUserUseCase struct {
	repo user.Repository
}

func NewDeleteUserUseCase(repo user.Repository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
