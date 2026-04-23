package usecase

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type UpdateProductUseCase struct {
	repo product.Repository
}

func NewUpdateProductUseCase(repo product.Repository) *UpdateProductUseCase {
	return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) Execute(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	existing, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, model.ErrNotFound
	}

	existing.Code = input.Code
	existing.Name = input.Name
	existing.Description = input.Description
	existing.CostPrice = input.CostPrice
	existing.Category = input.Category
	existing.ImageURL = input.ImageURL
	existing.Active = input.Active

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
