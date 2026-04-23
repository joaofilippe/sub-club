package subscription

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/subscription/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateSubscriptionInput) (*model.Subscription, error)
	GetByID(ctx context.Context, id string) (*model.Subscription, error)
	Update(ctx context.Context, input model.UpdateSubscriptionInput) (*model.Subscription, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error)
}
