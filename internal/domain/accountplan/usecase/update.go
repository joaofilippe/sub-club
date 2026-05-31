package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	modulemodel "github.com/joaofilippe/subclub/internal/domain/module/model"
)

type UpdateAccountPlanUseCase struct {
	repo accountplan.Repository
}

func NewUpdateAccountPlanUseCase(repo accountplan.Repository) *UpdateAccountPlanUseCase {
	return &UpdateAccountPlanUseCase{repo: repo}
}

func (uc *UpdateAccountPlanUseCase) Execute(ctx context.Context, input model.UpdateAccountPlanInput) (*model.AccountPlan, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	modules := make([]*modulemodel.Module, 0, len(input.ModuleIDs))
	for _, id := range input.ModuleIDs {
		modules = append(modules, &modulemodel.Module{ID: id})
	}

	existing.Name = input.Name
	existing.Description = input.Description
	existing.Price = input.Price
	existing.MaxCustomers = input.MaxCustomers
	existing.MaxPlans = input.MaxPlans
	existing.MaxProducts = input.MaxProducts
	existing.Active = input.Active
	existing.Modules = modules

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
