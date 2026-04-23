package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type UpdateClientUseCase struct {
	repo client.Repository
}

func NewUpdateClientUseCase(repo client.Repository) *UpdateClientUseCase {
	return &UpdateClientUseCase{repo: repo}
}

func (uc *UpdateClientUseCase) Execute(ctx context.Context, input model.UpdateClientInput) (*model.Client, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Name = input.Name
	existing.Email = input.Email
	existing.Phone = input.Phone
	existing.Document = input.Document
	existing.Active = input.Active
	existing.Address = input.Address

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
