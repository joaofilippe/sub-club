package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account"
)

type DeleteAccountUseCase struct {
	repo account.Repository
}

func NewDeleteAccountUseCase(repo account.Repository) *DeleteAccountUseCase {
	return &DeleteAccountUseCase{repo: repo}
}

func (uc *DeleteAccountUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
