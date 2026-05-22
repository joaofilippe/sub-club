package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type ListAccountsUseCase struct {
	repo account.Repository
}

func NewListAccountsUseCase(repo account.Repository) *ListAccountsUseCase {
	return &ListAccountsUseCase{repo: repo}
}

func (uc *ListAccountsUseCase) Execute(ctx context.Context) ([]*model.Account, error) {
	return uc.repo.List(ctx)
}
