package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
	"github.com/joaofilippe/subclub/internal/domain/module/usecase"
)

type ModuleService struct {
	createUseCase  *usecase.CreateModuleUseCase
	getByIDUseCase *usecase.GetModuleByIDUseCase
	updateUseCase  *usecase.UpdateModuleUseCase
	deleteUseCase  *usecase.DeleteModuleUseCase
	listUseCase    *usecase.ListModulesUseCase
}

func NewModuleService(repo module.Repository) *ModuleService {
	return &ModuleService{
		createUseCase:  usecase.NewCreateModuleUseCase(repo),
		getByIDUseCase: usecase.NewGetModuleByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateModuleUseCase(repo),
		deleteUseCase:  usecase.NewDeleteModuleUseCase(repo),
		listUseCase:    usecase.NewListModulesUseCase(repo),
	}
}

func (s *ModuleService) Create(ctx context.Context, input model.CreateModuleInput) (*model.Module, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *ModuleService) GetByID(ctx context.Context, id string) (*model.Module, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *ModuleService) Update(ctx context.Context, input model.UpdateModuleInput) (*model.Module, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *ModuleService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *ModuleService) List(ctx context.Context) ([]*model.Module, error) {
	return s.listUseCase.Execute(ctx)
}
