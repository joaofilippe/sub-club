package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan"
)

type DeleteAccountPlanUseCase struct {
	repo accountplan.Repository
}

func NewDeleteAccountPlanUseCase(repo accountplan.Repository) *DeleteAccountPlanUseCase {
	return &DeleteAccountPlanUseCase{repo: repo}
}

func (uc *DeleteAccountPlanUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
