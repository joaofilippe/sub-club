package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Email string        `json:"email"`
	Type  UserType `json:"type"`
	Role  UserRole `json:"role"`
}

type CreateUserUseCase struct {
	repo Repository
}

func NewCreateUserUseCase(repo Repository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (string, error) {
	now := time.Now()
	newID := uuid.New().String()

	u := &User{
		ID:        newID,
		Email:     input.Email,
		Type:      input.Type,
		Role:      input.Role,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := uc.repo.Create(ctx, u)
	if err != nil {
		return "", err
	}

	return newID, nil
}
