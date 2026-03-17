package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user"
	usecases "github.com/joaofilippe/subclub/internal/usecase/user"
)

type UserService struct {
	repo user.Repository
	createUseCase *usecases.CreateUserUseCase
}

func NewUserService(repo user.Repository, createUseCase *usecases.CreateUserUseCase) *UserService {
	return &UserService{
		repo: repo,
		createUseCase: createUseCase,
	}
}

func (s *UserService) Create(ctx context.Context, input usecases.CreateUserInput) (string,error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*user.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *UserService) GetByRole(ctx context.Context, role user.UserRole) ([]*user.User, error) {
	return s.repo.GetByRole(ctx, role)
}

func (s *UserService) GetByType(ctx context.Context, userType user.UserType) ([]*user.User, error) {
	return s.repo.GetByType(ctx, userType)
}

func (s *UserService) Update(ctx context.Context, user *user.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]*user.User, error) {
	return s.repo.List(ctx)
}