package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
)

type UpdateCustomerUseCase struct {
	repo customer.Repository
}

func NewUpdateCustomerUseCase(repo customer.Repository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{repo: repo}
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, input model.UpdateCustomerInput) (*model.Customer, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Name = input.Name
	existing.Email = input.Email
	existing.Phone = input.Phone
	existing.Document = input.Document
	existing.Active = input.Active
	existing.Address = input.Address

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
