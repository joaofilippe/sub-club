package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type GetProductByIDUseCase struct {
	repo product.Repository
}

func NewGetProductByIDUseCase(repo product.Repository) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{repo: repo}
}

func (uc *GetProductByIDUseCase) Execute(ctx context.Context, id string) (*model.Product, error) {
	return uc.repo.GetByID(ctx, id)
}
