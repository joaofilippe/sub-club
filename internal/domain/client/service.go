package client

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/client/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateClientInput) (*model.Client, error)
	GetByID(ctx context.Context, id string) (*model.Client, error)
	Update(ctx context.Context, input model.UpdateClientInput) (*model.Client, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error)
}
