package product

import (
	"context"
	"time"
)

type Product struct {
	ID          string
	Code        string
	Name        string
	Description string
	CostPrice   float64
	Category    string
	ImageURL    string
	Active      bool
	CreatedAt   time.Time
}

type Filter struct {
	Search   *string
	Category *string
	IsActive *bool
}

// Repository defines the interface for interacting with product storage
type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) ([]*Product, error)
}
