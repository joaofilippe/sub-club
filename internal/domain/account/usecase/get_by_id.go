package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type GetAccountByIDUseCase struct {
	repo account.Repository
}

func NewGetAccountByIDUseCase(repo account.Repository) *GetAccountByIDUseCase {
	return &GetAccountByIDUseCase{repo: repo}
}

func (uc *GetAccountByIDUseCase) Execute(ctx context.Context, id string) (*model.Account, error) {
	return uc.repo.GetByID(ctx, id)
}
