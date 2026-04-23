package user

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type Repository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByRole(ctx context.Context, role model.UserRole) ([]*model.User, error)
	GetByType(ctx context.Context, userType model.UserType) ([]*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.User, error)
}
