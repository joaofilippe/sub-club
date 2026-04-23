package services

import (
	"context"

	domain "github.com/joaofilippe/subclub/internal/domain/subscription"
)

type SubscriptionService struct {
	repo domain.Repository
}

func NewSubscriptionService(repo domain.Repository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(ctx context.Context, sub *domain.Subscription) error {
	return s.repo.Create(ctx, sub)
}

func (s *SubscriptionService) GetByID(ctx context.Context, id string) (*domain.Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SubscriptionService) Update(ctx context.Context, sub *domain.Subscription) error {
	return s.repo.Update(ctx, sub)
}

func (s *SubscriptionService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *SubscriptionService) List(ctx context.Context, filter domain.Filter) (*domain.PaginatedList, error) {
	return s.repo.List(ctx, filter)
}
