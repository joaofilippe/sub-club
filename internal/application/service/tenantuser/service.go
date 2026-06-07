package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/usecase"
)

type TenantUserService struct {
	createUseCase  *usecase.CreateTenantUserUseCase
	getByIDUseCase *usecase.GetTenantUserByIDUseCase
	updateUseCase  *usecase.UpdateTenantUserUseCase
	deleteUseCase  *usecase.DeleteTenantUserUseCase
	listUseCase    *usecase.ListTenantUsersUseCase
}

func NewTenantUserService(repo tenantuser.Repository) *TenantUserService {
	return &TenantUserService{
		createUseCase:  usecase.NewCreateTenantUserUseCase(repo),
		getByIDUseCase: usecase.NewGetTenantUserByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateTenantUserUseCase(repo),
		deleteUseCase:  usecase.NewDeleteTenantUserUseCase(repo),
		listUseCase:    usecase.NewListTenantUsersUseCase(repo),
	}
}

func (s *TenantUserService) Create(ctx context.Context, input model.CreateTenantUserInput) (*model.TenantUser, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *TenantUserService) GetByID(ctx context.Context, id string) (*model.TenantUser, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *TenantUserService) Update(ctx context.Context, input model.UpdateTenantUserInput) (*model.TenantUser, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *TenantUserService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *TenantUserService) List(ctx context.Context) ([]*model.TenantUser, error) {
	return s.listUseCase.Execute(ctx)
}
