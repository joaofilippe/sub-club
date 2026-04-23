package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type CreateUserUseCase struct {
	repo user.Repository
}

func NewCreateUserUseCase(repo user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input model.CreateUserInput) (string, error) {
	now := time.Now()
	newID := uuid.New().String()

	u := &model.User{
		ID:        newID,
		Email:     input.Email,
		Type:      input.Type,
		Role:      input.Role,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := uc.repo.Create(ctx, u); err != nil {
		return "", err
	}
	return newID, nil
}
