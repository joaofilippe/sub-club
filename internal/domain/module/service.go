package module

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/module/model"
)

type Service interface {
	Create(ctx context.Context, input model.CreateModuleInput) (*model.Module, error)
	GetByID(ctx context.Context, id string) (*model.Module, error)
	Update(ctx context.Context, input model.UpdateModuleInput) (*model.Module, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Module, error)
}
