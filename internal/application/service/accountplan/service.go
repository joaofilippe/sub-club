package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	"github.com/joaofilippe/subclub/internal/domain/accountplan/usecase"
)

type AccountPlanService struct {
	createUseCase  *usecase.CreateAccountPlanUseCase
	getByIDUseCase *usecase.GetAccountPlanByIDUseCase
	updateUseCase  *usecase.UpdateAccountPlanUseCase
	deleteUseCase  *usecase.DeleteAccountPlanUseCase
	listUseCase    *usecase.ListAccountPlansUseCase
}

func NewAccountPlanService(repo accountplan.Repository) *AccountPlanService {
	return &AccountPlanService{
		createUseCase:  usecase.NewCreateAccountPlanUseCase(repo),
		getByIDUseCase: usecase.NewGetAccountPlanByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateAccountPlanUseCase(repo),
		deleteUseCase:  usecase.NewDeleteAccountPlanUseCase(repo),
		listUseCase:    usecase.NewListAccountPlansUseCase(repo),
	}
}

func (s *AccountPlanService) Create(ctx context.Context, input model.CreateAccountPlanInput) (*model.AccountPlan, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *AccountPlanService) GetByID(ctx context.Context, id string) (*model.AccountPlan, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *AccountPlanService) Update(ctx context.Context, input model.UpdateAccountPlanInput) (*model.AccountPlan, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *AccountPlanService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *AccountPlanService) List(ctx context.Context) ([]*model.AccountPlan, error) {
	return s.listUseCase.Execute(ctx)
}
