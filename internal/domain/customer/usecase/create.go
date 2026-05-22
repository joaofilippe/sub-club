package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
)

type CreateCustomerUseCase struct {
	repo customer.Repository
}

func NewCreateCustomerUseCase(repo customer.Repository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{repo: repo}
}

func (uc *CreateCustomerUseCase) Execute(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	c := &model.Customer{
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
