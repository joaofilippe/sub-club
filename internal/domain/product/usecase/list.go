package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type ListProductsUseCase struct {
	repo product.Repository
}

func NewListProductsUseCase(repo product.Repository) *ListProductsUseCase {
	return &ListProductsUseCase{repo: repo}
}

func (uc *ListProductsUseCase) Execute(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return uc.repo.List(ctx, filter)
}
