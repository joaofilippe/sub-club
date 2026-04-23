package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
	"github.com/joaofilippe/subclub/internal/domain/client/usecase"
)

type ClientService struct {
	createUseCase  *usecase.CreateClientUseCase
	getByIDUseCase *usecase.GetClientByIDUseCase
	updateUseCase  *usecase.UpdateClientUseCase
	deleteUseCase  *usecase.DeleteClientUseCase
	listUseCase    *usecase.ListClientsUseCase
}

func NewClientService(repo client.Repository) *ClientService {
	return &ClientService{
		createUseCase:  usecase.NewCreateClientUseCase(repo),
		getByIDUseCase: usecase.NewGetClientByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateClientUseCase(repo),
		deleteUseCase:  usecase.NewDeleteClientUseCase(repo),
		listUseCase:    usecase.NewListClientsUseCase(repo),
	}
}

func (s *ClientService) Create(ctx context.Context, input model.CreateClientInput) (*model.Client, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *ClientService) GetByID(ctx context.Context, id string) (*model.Client, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *ClientService) Update(ctx context.Context, input model.UpdateClientInput) (*model.Client, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *ClientService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *ClientService) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return s.listUseCase.Execute(ctx, filter)
}
