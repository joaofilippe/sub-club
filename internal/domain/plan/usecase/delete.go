package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan"
)

type DeletePlanUseCase struct {
	repo plan.Repository
}

func NewDeletePlanUseCase(repo plan.Repository) *DeletePlanUseCase {
	return &DeletePlanUseCase{repo: repo}
}

func (uc *DeletePlanUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
