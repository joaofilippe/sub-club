package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type GetClientByIDUseCase struct {
	repo client.Repository
}

func NewGetClientByIDUseCase(repo client.Repository) *GetClientByIDUseCase {
	return &GetClientByIDUseCase{repo: repo}
}

func (uc *GetClientByIDUseCase) Execute(ctx context.Context, id string) (*model.Client, error) {
	return uc.repo.GetByID(ctx, id)
}
