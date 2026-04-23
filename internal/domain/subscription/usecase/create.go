package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/subscription"
	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type CreateSubscriptionUseCase struct {
	repo subscription.Repository
}

func NewCreateSubscriptionUseCase(repo subscription.Repository) *CreateSubscriptionUseCase {
	return &CreateSubscriptionUseCase{repo: repo}
}

func (uc *CreateSubscriptionUseCase) Execute(ctx context.Context, input model.CreateSubscriptionInput) (*model.Subscription, error) {
	sub := &model.Subscription{
		ID:               uuid.New().String(),
		ClientID:         input.ClientID,
		PlanID:           input.PlanID,
		Status:           input.Status,
		ShipmentStatus:   input.ShipmentStatus,
		StartDate:        input.StartDate,
		NextBillingDate:  input.NextBillingDate,
		NextShipmentDate: input.NextShipmentDate,
		CreatedAt:        time.Now(),
	}
	if err := uc.repo.Create(ctx, sub); err != nil {
		return nil, err
	}
	return sub, nil
}
