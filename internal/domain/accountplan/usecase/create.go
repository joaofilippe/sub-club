package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	modulemodel "github.com/joaofilippe/subclub/internal/domain/module/model"
)

type CreateAccountPlanUseCase struct {
	repo accountplan.Repository
}

func NewCreateAccountPlanUseCase(repo accountplan.Repository) *CreateAccountPlanUseCase {
	return &CreateAccountPlanUseCase{repo: repo}
}

func (uc *CreateAccountPlanUseCase) Execute(ctx context.Context, input model.CreateAccountPlanInput) (*model.AccountPlan, error) {
	modules := make([]*modulemodel.Module, 0, len(input.ModuleIDs))
	for _, id := range input.ModuleIDs {
		modules = append(modules, &modulemodel.Module{ID: id})
	}

	p := &model.AccountPlan{
		ID:           uuid.New().String(),
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		MaxCustomers: input.MaxCustomers,
		MaxPlans:     input.MaxPlans,
		MaxProducts:  input.MaxProducts,
		Active:       true,
		CreatedAt:    time.Now(),
		Modules:      modules,
	}
	if err := uc.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
