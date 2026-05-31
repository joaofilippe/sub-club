package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module"
)

type DeleteModuleUseCase struct {
	repo module.Repository
}

func NewDeleteModuleUseCase(repo module.Repository) *DeleteModuleUseCase {
	return &DeleteModuleUseCase{repo: repo}
}

func (uc *DeleteModuleUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
