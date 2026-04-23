package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type GetSubscriptionByIDUseCase struct {
	repo subscription.Repository
}

func NewGetSubscriptionByIDUseCase(repo subscription.Repository) *GetSubscriptionByIDUseCase {
	return &GetSubscriptionByIDUseCase{repo: repo}
}

func (uc *GetSubscriptionByIDUseCase) Execute(ctx context.Context, id string) (*model.Subscription, error) {
	return uc.repo.GetByID(ctx, id)
}
