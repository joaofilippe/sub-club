package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type GetModuleByIDUseCase struct {
	repo module.Repository
}

func NewGetModuleByIDUseCase(repo module.Repository) *GetModuleByIDUseCase {
	return &GetModuleByIDUseCase{repo: repo}
}

func (uc *GetModuleByIDUseCase) Execute(ctx context.Context, id string) (*model.Module, error) {
	return uc.repo.GetByID(ctx, id)
}
