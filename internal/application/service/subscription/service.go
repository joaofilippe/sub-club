package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
	"github.com/joaofilippe/subclub/internal/domain/subscription/usecase"
)

type SubscriptionService struct {
	createUseCase  *usecase.CreateSubscriptionUseCase
	getByIDUseCase *usecase.GetSubscriptionByIDUseCase
	updateUseCase  *usecase.UpdateSubscriptionUseCase
	deleteUseCase  *usecase.DeleteSubscriptionUseCase
	listUseCase    *usecase.ListSubscriptionsUseCase
}

func NewSubscriptionService(repo subscription.Repository) *SubscriptionService {
	return &SubscriptionService{
		createUseCase:  usecase.NewCreateSubscriptionUseCase(repo),
		getByIDUseCase: usecase.NewGetSubscriptionByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateSubscriptionUseCase(repo),
		deleteUseCase:  usecase.NewDeleteSubscriptionUseCase(repo),
		listUseCase:    usecase.NewListSubscriptionsUseCase(repo),
	}
}

func (s *SubscriptionService) Create(ctx context.Context, input model.CreateSubscriptionInput) (*model.Subscription, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *SubscriptionService) GetByID(ctx context.Context, id string) (*model.Subscription, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *SubscriptionService) Update(ctx context.Context, input model.UpdateSubscriptionInput) (*model.Subscription, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *SubscriptionService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *SubscriptionService) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return s.listUseCase.Execute(ctx, filter)
}
