package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type ListClientsUseCase struct {
	repo client.Repository
}

func NewListClientsUseCase(repo client.Repository) *ListClientsUseCase {
	return &ListClientsUseCase{repo: repo}
}

func (uc *ListClientsUseCase) Execute(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return uc.repo.List(ctx, filter)
}
