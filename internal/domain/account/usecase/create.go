package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type CreateAccountUseCase struct {
	repo account.Repository
}

func NewCreateAccountUseCase(repo account.Repository) *CreateAccountUseCase {
	return &CreateAccountUseCase{repo: repo}
}

func (uc *CreateAccountUseCase) Execute(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	a := &model.Account{
		ID:                 uuid.New().String(),
		Name:               input.Name,
		Email:              input.Email,
		Document:           input.Document,
		Slug:               input.Slug,
		AccountPlanID:      input.AccountPlanID,
		SubscriptionStatus: model.SubscriptionStatusTrial,
		Active:             true,
		CreatedAt:          time.Now(),
	}
	if err := uc.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}
