package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client"
)

type DeleteClientUseCase struct {
	repo client.Repository
}

func NewDeleteClientUseCase(repo client.Repository) *DeleteClientUseCase {
	return &DeleteClientUseCase{repo: repo}
}

func (uc *DeleteClientUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
