package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
	"golang.org/x/crypto/bcrypt"
)

type CreateTenantUserUseCase struct {
	repo tenantuser.Repository
}

func NewCreateTenantUserUseCase(repo tenantuser.Repository) *CreateTenantUserUseCase {
	return &CreateTenantUserUseCase{repo: repo}
}

func (uc *CreateTenantUserUseCase) Execute(ctx context.Context, input model.CreateTenantUserInput) (*model.TenantUser, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &model.TenantUser{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Username:  input.Username,
		Email:     input.Email,
		Role:      input.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := uc.repo.Create(ctx, u, string(hash)); err != nil {
		return nil, err
	}
	return u, nil
}
