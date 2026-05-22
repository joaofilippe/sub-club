package accountplan

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
)

type Repository interface {
	Create(ctx context.Context, plan *model.AccountPlan) error
	GetByID(ctx context.Context, id string) (*model.AccountPlan, error)
	Update(ctx context.Context, plan *model.AccountPlan) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.AccountPlan, error)
}
