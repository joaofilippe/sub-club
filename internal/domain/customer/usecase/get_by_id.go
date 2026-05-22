package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
)

type GetCustomerByIDUseCase struct {
	repo customer.Repository
}

func NewGetCustomerByIDUseCase(repo customer.Repository) *GetCustomerByIDUseCase {
	return &GetCustomerByIDUseCase{repo: repo}
}

func (uc *GetCustomerByIDUseCase) Execute(ctx context.Context, id string) (*model.Customer, error) {
	return uc.repo.GetByID(ctx, id)
}
