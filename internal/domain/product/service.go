package product

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error)
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error)
}
