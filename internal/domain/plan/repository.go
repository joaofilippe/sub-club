package plan

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/plan/model"
)

type Repository interface {
	Create(ctx context.Context, plan *model.Plan) error
	GetByID(ctx context.Context, id string) (*model.Plan, error)
	Update(ctx context.Context, plan *model.Plan) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error)
}
