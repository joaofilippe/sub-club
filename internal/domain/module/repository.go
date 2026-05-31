package module

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type Repository interface {
	Create(ctx context.Context, m *model.Module) error
	GetByID(ctx context.Context, id string) (*model.Module, error)
	Update(ctx context.Context, m *model.Module) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Module, error)
}
