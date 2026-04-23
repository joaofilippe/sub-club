package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type UpdateSubscriptionUseCase struct {
	repo subscription.Repository
}

func NewUpdateSubscriptionUseCase(repo subscription.Repository) *UpdateSubscriptionUseCase {
	return &UpdateSubscriptionUseCase{repo: repo}
}

func (uc *UpdateSubscriptionUseCase) Execute(ctx context.Context, input model.UpdateSubscriptionInput) (*model.Subscription, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Status = input.Status
	existing.ShipmentStatus = input.ShipmentStatus
	existing.NextBillingDate = input.NextBillingDate
	existing.NextShipmentDate = input.NextShipmentDate

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
