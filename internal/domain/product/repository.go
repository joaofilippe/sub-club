package product

import (
	"context"
)

type Filter struct {
	Search   *string
	Category *string
	IsActive *bool
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Product
	TotalCount int
}

// Repository defines the interface for interacting with product storage
type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
