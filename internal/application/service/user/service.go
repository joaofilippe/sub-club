package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
	"github.com/joaofilippe/subclub/internal/domain/user/usecase"
)

type UserService struct {
	createUseCase     *usecase.CreateUserUseCase
	getByIDUseCase    *usecase.GetUserByIDUseCase
	getByEmailUseCase *usecase.GetUserByEmailUseCase
	getByRoleUseCase  *usecase.GetUsersByRoleUseCase
	getByTypeUseCase  *usecase.GetUsersByTypeUseCase
	updateUseCase     *usecase.UpdateUserUseCase
	deleteUseCase     *usecase.DeleteUserUseCase
	listUseCase       *usecase.ListUsersUseCase
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		createUseCase:     usecase.NewCreateUserUseCase(repo),
		getByIDUseCase:    usecase.NewGetUserByIDUseCase(repo),
		getByEmailUseCase: usecase.NewGetUserByEmailUseCase(repo),
		getByRoleUseCase:  usecase.NewGetUsersByRoleUseCase(repo),
		getByTypeUseCase:  usecase.NewGetUsersByTypeUseCase(repo),
		updateUseCase:     usecase.NewUpdateUserUseCase(repo),
		deleteUseCase:     usecase.NewDeleteUserUseCase(repo),
		listUseCase:       usecase.NewListUsersUseCase(repo),
	}
}

func (s *UserService) Create(ctx context.Context, input model.CreateUserInput) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password: %w", err)
	}
	input.Password = string(hash)
	return s.createUseCase.Execute(ctx, input)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.getByEmailUseCase.Execute(ctx, email)
}

func (s *UserService) GetByRole(ctx context.Context, role model.UserRole) ([]*model.User, error) {
	return s.getByRoleUseCase.Execute(ctx, role)
}

func (s *UserService) GetByType(ctx context.Context, userType model.UserType) ([]*model.User, error) {
	return s.getByTypeUseCase.Execute(ctx, userType)
}

func (s *UserService) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]*model.User, error) {
	return s.listUseCase.Execute(ctx)
}
