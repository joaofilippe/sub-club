package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type CreatePlanUseCase struct {
	repo plan.Repository
}

func NewCreatePlanUseCase(repo plan.Repository) *CreatePlanUseCase {
	return &CreatePlanUseCase{repo: repo}
}

func (uc *CreatePlanUseCase) Execute(ctx context.Context, input model.CreatePlanInput) (*model.Plan, error) {
	p := &model.Plan{
		ID:            uuid.New().String(),
		Code:          input.Code,
		Name:          input.Name,
		Description:   input.Description,
		ProductValue:  input.ProductValue,
		DiscountValue: input.DiscountValue,
		Price:         input.Price,
		IntervalDays:  input.IntervalDays,
		Active:        input.Active,
		ImageURL:      input.ImageURL,
		CreatedAt:     time.Now(),
	}
	if err := uc.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
