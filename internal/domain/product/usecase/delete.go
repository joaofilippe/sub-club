package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product"
)

type DeleteProductUseCase struct {
	repo product.Repository
}

func NewDeleteProductUseCase(repo product.Repository) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
