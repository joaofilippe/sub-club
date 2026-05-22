package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
	"github.com/joaofilippe/subclub/internal/domain/account/usecase"
)

type AccountService struct {
	createUseCase     *usecase.CreateAccountUseCase
	getByIDUseCase    *usecase.GetAccountByIDUseCase
	getBySlugUseCase  *usecase.GetAccountBySlugUseCase
	updateUseCase     *usecase.UpdateAccountUseCase
	deleteUseCase     *usecase.DeleteAccountUseCase
	listUseCase       *usecase.ListAccountsUseCase
}

func NewAccountService(repo account.Repository) *AccountService {
	return &AccountService{
		createUseCase:    usecase.NewCreateAccountUseCase(repo),
		getByIDUseCase:   usecase.NewGetAccountByIDUseCase(repo),
		getBySlugUseCase: usecase.NewGetAccountBySlugUseCase(repo),
		updateUseCase:    usecase.NewUpdateAccountUseCase(repo),
		deleteUseCase:    usecase.NewDeleteAccountUseCase(repo),
		listUseCase:      usecase.NewListAccountsUseCase(repo),
	}
}

func (s *AccountService) Create(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	return s.createUseCase.Execute(ctx, input)
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
