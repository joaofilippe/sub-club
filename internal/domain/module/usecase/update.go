package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type UpdateModuleUseCase struct {
	repo module.Repository
}

func NewUpdateModuleUseCase(repo module.Repository) *UpdateModuleUseCase {
	return &UpdateModuleUseCase{repo: repo}
}

func (uc *UpdateModuleUseCase) Execute(ctx context.Context, input model.UpdateModuleInput) (*model.Module, error) {
	m, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	m.Name = input.Name
	m.Active = input.Active
	if err := uc.repo.Update(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
