package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/module"
	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type CreateModuleUseCase struct {
	repo module.Repository
}

func NewCreateModuleUseCase(repo module.Repository) *CreateModuleUseCase {
	return &CreateModuleUseCase{repo: repo}
}

func (uc *CreateModuleUseCase) Execute(ctx context.Context, input model.CreateModuleInput) (*model.Module, error) {
	m := &model.Module{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Active:    true,
		CreatedAt: time.Now(),
	}
	if err := uc.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
