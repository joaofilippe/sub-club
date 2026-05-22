package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/customer"
	"github.com/joaofilippe/subclub/internal/domain/customer/model"
)

type ListCustomersUseCase struct {
	repo customer.Repository
}

func NewListCustomersUseCase(repo customer.Repository) *ListCustomersUseCase {
	return &ListCustomersUseCase{repo: repo}
}

func (uc *ListCustomersUseCase) Execute(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return uc.repo.List(ctx, filter)
}
