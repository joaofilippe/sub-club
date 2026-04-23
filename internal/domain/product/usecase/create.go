package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type CreateProductUseCase struct {
	repo product.Repository
}

func NewCreateProductUseCase(repo product.Repository) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	p := &model.Product{
		ID:          uuid.New().String(),
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		CostPrice:   input.CostPrice,
		Category:    input.Category,
		ImageURL:    input.ImageURL,
		Active:      input.Active,
		CreatedAt:   time.Now(),
	}
	if err := uc.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
