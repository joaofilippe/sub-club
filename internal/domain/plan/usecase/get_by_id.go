package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type GetPlanByIDUseCase struct {
	repo plan.Repository
}

func NewGetPlanByIDUseCase(repo plan.Repository) *GetPlanByIDUseCase {
	return &GetPlanByIDUseCase{repo: repo}
}

func (uc *GetPlanByIDUseCase) Execute(ctx context.Context, id string) (*model.Plan, error) {
	return uc.repo.GetByID(ctx, id)
}
