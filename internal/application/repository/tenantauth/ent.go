package repository

import (
	"context"

	entuser "github.com/joaofilippe/subclub/ent/user"
	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/infra/database"
)

type TenantAuthEntRepository struct {
	manager *database.TenantClientManager
}

func NewTenantAuthEntRepository(manager *database.TenantClientManager) authdomain.TenantRepository {
	return &TenantAuthEntRepository{manager: manager}
}

func (r *TenantAuthEntRepository) FindUserByEmailAndSlug(ctx context.Context, slug, email string) (*authmodel.TenantAuthUser, error) {
	client, err := r.manager.GetOrCreate(slug)
	if err != nil {
		return nil, err
	}
	u, err := client.User.Query().
		Where(entuser.EmailEQ(email), entuser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &authmodel.TenantAuthUser{
		ID:           u.ID.String(),
		PasswordHash: u.Password,
		Role:         string(u.Role),
	}, nil
}

func (r *TenantAuthEntRepository) FindUserByUsernameAndSlug(ctx context.Context, slug, username string) (*authmodel.TenantAuthUser, error) {
	client, err := r.manager.GetOrCreate(slug)
	if err != nil {
		return nil, err
	}
	u, err := client.User.Query().
		Where(entuser.UsernameEQ(username), entuser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &authmodel.TenantAuthUser{
		ID:           u.ID.String(),
		PasswordHash: u.Password,
		Role:         string(u.Role),
	}, nil
}

func (r *TenantAuthEntRepository) UserExistsByEmailAndSlug(ctx context.Context, slug, email string) (bool, error) {
	client, err := r.manager.GetOrCreate(slug)
	if err != nil {
		return false, err
	}
	return client.User.Query().
		Where(entuser.EmailEQ(email), entuser.DeletedAtIsNil()).
		Exist(ctx)
}
