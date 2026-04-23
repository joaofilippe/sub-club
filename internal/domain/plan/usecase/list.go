package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type ListPlansUseCase struct {
	repo plan.Repository
}

func NewListPlansUseCase(repo plan.Repository) *ListPlansUseCase {
	return &ListPlansUseCase{repo: repo}
}

func (uc *ListPlansUseCase) Execute(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return uc.repo.List(ctx, filter)
}
