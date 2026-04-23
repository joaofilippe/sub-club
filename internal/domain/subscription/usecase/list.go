package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type ListSubscriptionsUseCase struct {
	repo subscription.Repository
}

func NewListSubscriptionsUseCase(repo subscription.Repository) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{repo: repo}
}

func (uc *ListSubscriptionsUseCase) Execute(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return uc.repo.List(ctx, filter)
}
