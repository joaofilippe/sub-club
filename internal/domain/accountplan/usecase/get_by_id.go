package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
)

type GetAccountPlanByIDUseCase struct {
	repo accountplan.Repository
}

func NewGetAccountPlanByIDUseCase(repo accountplan.Repository) *GetAccountPlanByIDUseCase {
	return &GetAccountPlanByIDUseCase{repo: repo}
}

func (uc *GetAccountPlanByIDUseCase) Execute(ctx context.Context, id string) (*model.AccountPlan, error) {
	return uc.repo.GetByID(ctx, id)
}
