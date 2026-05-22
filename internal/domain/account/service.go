package account

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateAccountInput) (*model.Account, error)
	GetByID(ctx context.Context, id string) (*model.Account, error)
	GetBySlug(ctx context.Context, slug string) (*model.Account, error)
	Update(ctx context.Context, input model.UpdateAccountInput) (*model.Account, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Account, error)
}
