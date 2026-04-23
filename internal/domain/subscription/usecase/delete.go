package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription"
)

type DeleteSubscriptionUseCase struct {
	repo subscription.Repository
}

func NewDeleteSubscriptionUseCase(repo subscription.Repository) *DeleteSubscriptionUseCase {
	return &DeleteSubscriptionUseCase{repo: repo}
}

func (uc *DeleteSubscriptionUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
