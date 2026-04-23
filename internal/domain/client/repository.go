package client

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type Repository interface {
	Create(ctx context.Context, client *model.Client) error
	GetByID(ctx context.Context, id string) (*model.Client, error)
	Update(ctx context.Context, client *model.Client) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error)
}
