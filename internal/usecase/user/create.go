package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/user"
)

type CreateUserInput struct {
	Email string        `json:"email"`
	Type  user.UserType `json:"type"`
	Role  user.UserRole `json:"role"`
}

type CreateUserUseCase struct {
	repo user.Repository
}

func NewCreateUserUseCase(repo user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (string, error) {
	now := time.Now()
	newID := uuid.New().String()

	u := &user.User{
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
