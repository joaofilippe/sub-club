package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entsystemuser "github.com/joaofilippe/subclub/ent/systemuser"
	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
)

type userEntRepository struct {
	client *ent.Client
}

func NewUserEntRepository(client *ent.Client) user.Repository {
	return &userEntRepository{client: client}
}

func (r *userEntRepository) Create(ctx context.Context, u *model.User) error {
	id, err := uuid.Parse(u.ID)
	if err != nil {
		id = uuid.New()
	}

	_, err = r.client.SystemUser.Create().
		SetID(id).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(entsystemuser.Type(u.Type)).
		SetRole(entsystemuser.Role(u.Role)).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(u.UpdatedAt).
		Save(ctx)
	return err
}

func (r *userEntRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	u, err := r.client.SystemUser.Query().
		Where(entsystemuser.IDEQ(parsedID), entsystemuser.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntSystemUserToDomain(u), nil
}

func (r *userEntRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := r.client.SystemUser.Query().
		Where(entsystemuser.EmailEQ(email), entsystemuser.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntSystemUserToDomain(u), nil
}

func (r *userEntRepository) GetByRole(ctx context.Context, role model.UserRole) ([]*model.User, error) {
	users, err := r.client.SystemUser.Query().
		Where(entsystemuser.RoleEQ(entsystemuser.Role(role)), entsystemuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntSystemUserToDomain(eu))
	}
	return results, nil
}

func (r *userEntRepository) GetByType(ctx context.Context, userType model.UserType) ([]*model.User, error) {
	users, err := r.client.SystemUser.Query().
		Where(entsystemuser.TypeEQ(entsystemuser.Type(userType)), entsystemuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntSystemUserToDomain(eu))
	}
	return results, nil
}

func (r *userEntRepository) Update(ctx context.Context, u *model.User) error {
	parsedID, err := uuid.Parse(u.ID)
	if err != nil {
		return err
	}

	builder := r.client.SystemUser.UpdateOneID(parsedID).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(entsystemuser.Type(u.Type)).
		SetRole(entsystemuser.Role(u.Role)).
		SetUpdatedAt(u.UpdatedAt)

	if u.DeletedAt != nil {
		builder.SetDeletedAt(*u.DeletedAt)
	} else {
		builder.ClearDeletedAt()
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *userEntRepository) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.client.SystemUser.UpdateOneID(parsedID).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func (r *userEntRepository) List(ctx context.Context) ([]*model.User, error) {
	users, err := r.client.SystemUser.Query().
		Where(entsystemuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntSystemUserToDomain(eu))
	}
	return results, nil
}

func mapEntSystemUserToDomain(eu *ent.SystemUser) *model.User {
	return &model.User{
		ID:        eu.ID.String(),
		Name:      eu.Name,
		Email:     eu.Email,
		Password:  eu.Password,
		Type:      model.UserType(eu.Type),
		Role:      model.UserRole(eu.Role),
		CreatedAt: eu.CreatedAt,
		UpdatedAt: eu.UpdatedAt,
		DeletedAt: eu.DeletedAt,
	}
}
