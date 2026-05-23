package service

import (
	"context"
	"fmt"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
	"github.com/joaofilippe/subclub/internal/domain/account/usecase"
)

type AccountService struct {
	createUseCase    *usecase.CreateAccountUseCase
	getByIDUseCase   *usecase.GetAccountByIDUseCase
	getBySlugUseCase *usecase.GetAccountBySlugUseCase
	updateUseCase    *usecase.UpdateAccountUseCase
	deleteUseCase    *usecase.DeleteAccountUseCase
	listUseCase      *usecase.ListAccountsUseCase
	schemaCreator    account.SchemaCreator
}

func NewAccountService(repo account.Repository, schemaCreator account.SchemaCreator) *AccountService {
	return &AccountService{
		createUseCase:    usecase.NewCreateAccountUseCase(repo),
		getByIDUseCase:   usecase.NewGetAccountByIDUseCase(repo),
		getBySlugUseCase: usecase.NewGetAccountBySlugUseCase(repo),
		updateUseCase:    usecase.NewUpdateAccountUseCase(repo),
		deleteUseCase:    usecase.NewDeleteAccountUseCase(repo),
		listUseCase:      usecase.NewListAccountsUseCase(repo),
		schemaCreator:    schemaCreator,
	}
}

func (s *AccountService) Create(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	acc, err := s.createUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	if err := s.schemaCreator.CreateTenantSchema(ctx, acc.Slug); err != nil {
		return nil, fmt.Errorf("provisioning tenant schema: %w", err)
	}

	return acc, nil
}

func (s *AccountService) GetByID(ctx context.Context, id string) (*model.Account, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *AccountService) GetBySlug(ctx context.Context, slug string) (*model.Account, error) {
	return s.getBySlugUseCase.Execute(ctx, slug)
}

func (s *AccountService) Update(ctx context.Context, input model.UpdateAccountInput) (*model.Account, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *AccountService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *AccountService) List(ctx context.Context) ([]*model.Account, error) {
	return s.listUseCase.Execute(ctx)
}
