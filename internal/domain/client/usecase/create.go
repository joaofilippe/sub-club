package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/client"
	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type CreateClientUseCase struct {
	repo client.Repository
}

func NewCreateClientUseCase(repo client.Repository) *CreateClientUseCase {
	return &CreateClientUseCase{repo: repo}
}

func (uc *CreateClientUseCase) Execute(ctx context.Context, input model.CreateClientInput) (*model.Client, error) {
	c := &model.Client{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Phone:     input.Phone,
		Document:  input.Document,
		Active:    input.Active,
		Address:   input.Address,
		CreatedAt: time.Now(),
	}
	if err := uc.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
