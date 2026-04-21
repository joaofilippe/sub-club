package product

import (
	"context"
)

// Service defines the interface for product-related business logic
type Service interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
