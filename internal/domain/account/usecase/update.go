package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type UpdateAccountUseCase struct {
	repo account.Repository
}

func NewUpdateAccountUseCase(repo account.Repository) *UpdateAccountUseCase {
	return &UpdateAccountUseCase{repo: repo}
}

func (uc *UpdateAccountUseCase) Execute(ctx context.Context, input model.UpdateAccountInput) (*model.Account, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Name = input.Name
	existing.Email = input.Email
	existing.Document = input.Document
	existing.AccountPlanID = input.AccountPlanID
	existing.SubscriptionStatus = input.SubscriptionStatus
	existing.SubscriptionExpiresAt = input.SubscriptionExpiresAt
	existing.Active = input.Active

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
