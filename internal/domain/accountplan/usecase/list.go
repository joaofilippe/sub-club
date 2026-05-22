package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
)

type ListAccountPlansUseCase struct {
	repo accountplan.Repository
}

func NewListAccountPlansUseCase(repo accountplan.Repository) *ListAccountPlansUseCase {
	return &ListAccountPlansUseCase{repo: repo}
}

func (uc *ListAccountPlansUseCase) Execute(ctx context.Context) ([]*model.AccountPlan, error) {
	return uc.repo.List(ctx)
}
