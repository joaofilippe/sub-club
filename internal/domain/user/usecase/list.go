package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type ListUsersUseCase struct {
	repo user.Repository
}

func NewListUsersUseCase(repo user.Repository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (uc *ListUsersUseCase) Execute(ctx context.Context) ([]*model.User, error) {
	return uc.repo.List(ctx)
}
