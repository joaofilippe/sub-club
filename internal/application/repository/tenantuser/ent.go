package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entuser "github.com/joaofilippe/subclub/ent/user"
	tenantuserdomain "github.com/joaofilippe/subclub/internal/domain/tenantuser"
	"github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
	"github.com/joaofilippe/subclub/internal/infra/tenantctx"
)

type TenantUserEntRepository struct {
	entClient *ent.Client
}

func NewTenantUserEntRepository(entClient *ent.Client) tenantuserdomain.Repository {
	return &TenantUserEntRepository{entClient: entClient}
}

func (r *TenantUserEntRepository) tenantClient(ctx context.Context) *ent.Client {
	if c := tenantctx.TenantClientFromContext(ctx); c != nil {
		return c
	}
	return r.entClient
}

func (r *TenantUserEntRepository) Create(ctx context.Context, u *model.TenantUser, passwordHash string) error {
	id, err := uuid.Parse(u.ID)
	if err != nil {
		id = uuid.New()
	}
	_, err = r.tenantClient(ctx).User.Create().
		SetID(id).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(passwordHash).
		SetRole(entuser.Role(u.Role)).
		Save(ctx)
	return err
}

func (r *TenantUserEntRepository) GetByID(ctx context.Context, id string) (*model.TenantUser, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	u, err := r.tenantClient(ctx).User.Query().
		Where(entuser.IDEQ(parsedID), entuser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntUserToDomain(u), nil
}

func (r *TenantUserEntRepository) Update(ctx context.Context, u *model.TenantUser) error {
	parsedID, err := uuid.Parse(u.ID)
	if err != nil {
		return err
	}
	_, err = r.tenantClient(ctx).User.UpdateOneID(parsedID).
		SetName(u.Name).
		SetRole(entuser.Role(u.Role)).
		SetUpdatedAt(u.UpdatedAt).
		Save(ctx)
	return err
}

func (r *TenantUserEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = r.tenantClient(ctx).User.UpdateOneID(parsedID).
		SetDeletedAt(now).
		Save(ctx)
	return err
}

func (r *TenantUserEntRepository) List(ctx context.Context) ([]*model.TenantUser, error) {
	users, err := r.tenantClient(ctx).User.Query().
		Where(entuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TenantUser, len(users))
	for i, u := range users {
		result[i] = mapEntUserToDomain(u)
	}
	return result, nil
}

func mapEntUserToDomain(u *ent.User) *model.TenantUser {
	tu := &model.TenantUser{
		ID:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Role:      model.Role(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if u.DeletedAt != nil {
		tu.DeletedAt = u.DeletedAt
	}
	return tu
}
