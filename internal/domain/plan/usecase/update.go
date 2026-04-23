package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type UpdatePlanUseCase struct {
	repo plan.Repository
}

func NewUpdatePlanUseCase(repo plan.Repository) *UpdatePlanUseCase {
	return &UpdatePlanUseCase{repo: repo}
}

func (uc *UpdatePlanUseCase) Execute(ctx context.Context, input model.UpdatePlanInput) (*model.Plan, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Code = input.Code
	existing.Name = input.Name
	existing.Description = input.Description
	existing.ProductValue = input.ProductValue
	existing.DiscountValue = input.DiscountValue
	existing.Price = input.Price
	existing.IntervalDays = input.IntervalDays
	existing.Active = input.Active
	existing.ImageURL = input.ImageURL

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
