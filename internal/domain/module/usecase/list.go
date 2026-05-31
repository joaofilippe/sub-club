package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type ListModulesUseCase struct {
	repo module.Repository
}

func NewListModulesUseCase(repo module.Repository) *ListModulesUseCase {
	return &ListModulesUseCase{repo: repo}
}

func (uc *ListModulesUseCase) Execute(ctx context.Context) ([]*model.Module, error) {
	return uc.repo.List(ctx)
}
