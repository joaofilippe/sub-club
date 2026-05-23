package account

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/account/model"
)

type Repository interface {
	Create(ctx context.Context, account *model.Account) error
	GetByID(ctx context.Context, id string) (*model.Account, error)
	GetBySlug(ctx context.Context, slug string) (*model.Account, error)
	Update(ctx context.Context, account *model.Account) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Account, error)
}
