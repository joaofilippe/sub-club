package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/customer"
)

type DeleteCustomerUseCase struct {
	repo customer.Repository
}

func NewDeleteCustomerUseCase(repo customer.Repository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{repo: repo}
}

func (uc *DeleteCustomerUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
