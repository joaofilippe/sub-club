package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
	"github.com/joaofilippe/subclub/internal/domain/customer/usecase"
)

type CustomerService struct {
	createUseCase  *usecase.CreateCustomerUseCase
	getByIDUseCase *usecase.GetCustomerByIDUseCase
	updateUseCase  *usecase.UpdateCustomerUseCase
	deleteUseCase  *usecase.DeleteCustomerUseCase
	listUseCase    *usecase.ListCustomersUseCase
}

func NewCustomerService(repo customer.Repository) *CustomerService {
	return &CustomerService{
		createUseCase:  usecase.NewCreateCustomerUseCase(repo),
		getByIDUseCase: usecase.NewGetCustomerByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateCustomerUseCase(repo),
		deleteUseCase:  usecase.NewDeleteCustomerUseCase(repo),
		listUseCase:    usecase.NewListCustomersUseCase(repo),
	}
}

func (s *CustomerService) Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *CustomerService) GetByID(ctx context.Context, id string) (*model.Customer, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *CustomerService) Update(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *CustomerService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *CustomerService) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return s.listUseCase.Execute(ctx, filter)
}
