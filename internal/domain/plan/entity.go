package plan

import (
	"context"
	"time"
)

type Plan struct {
	ID            string
	Code          string
	Name          string
	Description   string
	ProductValue  float64
	DiscountValue float64
	Price         float64
	IntervalDays  int
	Active        bool
	ImageURL      string
	CreatedAt     time.Time
}

type Filter struct {
	Search   *string
	IsActive *bool
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Plan
	TotalCount int
}

// Repository defines the interface for interacting with plan storage
type Repository interface {
	Create(ctx context.Context, plan *Plan) error
	GetByID(ctx context.Context, id string) (*Plan, error)
	Update(ctx context.Context, plan *Plan) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
