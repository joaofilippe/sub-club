package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type GetAccountBySlugUseCase struct {
	repo account.Repository
}

func NewGetAccountBySlugUseCase(repo account.Repository) *GetAccountBySlugUseCase {
	return &GetAccountBySlugUseCase{repo: repo}
}

func (uc *GetAccountBySlugUseCase) Execute(ctx context.Context, slug string) (*model.Account, error) {
	return uc.repo.GetBySlug(ctx, slug)
}
