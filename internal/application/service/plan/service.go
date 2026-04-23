package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan"
	"github.com/joaofilippe/subclub/internal/domain/plan/model"
	"github.com/joaofilippe/subclub/internal/domain/plan/usecase"
)

type PlanService struct {
	createUseCase  *usecase.CreatePlanUseCase
	getByIDUseCase *usecase.GetPlanByIDUseCase
	updateUseCase  *usecase.UpdatePlanUseCase
	deleteUseCase  *usecase.DeletePlanUseCase
	listUseCase    *usecase.ListPlansUseCase
}

func NewPlanService(repo plan.Repository) *PlanService {
	return &PlanService{
		createUseCase:  usecase.NewCreatePlanUseCase(repo),
		getByIDUseCase: usecase.NewGetPlanByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdatePlanUseCase(repo),
		deleteUseCase:  usecase.NewDeletePlanUseCase(repo),
		listUseCase:    usecase.NewListPlansUseCase(repo),
	}
}

func (s *PlanService) Create(ctx context.Context, input model.CreatePlanInput) (*model.Plan, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *PlanService) GetByID(ctx context.Context, id string) (*model.Plan, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *PlanService) Update(ctx context.Context, input model.UpdatePlanInput) (*model.Plan, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *PlanService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *PlanService) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return s.listUseCase.Execute(ctx, filter)
}
