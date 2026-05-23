package accountplan

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/accountplan/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateAccountPlanInput) (*model.AccountPlan, error)
	GetByID(ctx context.Context, id string) (*model.AccountPlan, error)
	Update(ctx context.Context, input model.UpdateAccountPlanInput) (*model.AccountPlan, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.AccountPlan, error)
}
