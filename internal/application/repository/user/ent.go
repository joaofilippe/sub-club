package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	entuser "github.com/joaofilippe/subclub/ent/user"
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

	builder := r.client.User.Create().
		SetID(id).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(entuser.Type(u.Type)).
		SetRole(entuser.Role(u.Role)).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(u.UpdatedAt)

	if u.DeletedAt != nil {
		builder.SetDeletedAt(*u.DeletedAt)
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *userEntRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	u, err := r.client.User.Query().
		Where(entuser.IDEQ(parsedID), entuser.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntUserToDomain(u), nil
}

func (r *userEntRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := r.client.User.Query().
		Where(entuser.EmailEQ(email), entuser.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntUserToDomain(u), nil
}

func (r *userEntRepository) GetByRole(ctx context.Context, role model.UserRole) ([]*model.User, error) {
	users, err := r.client.User.Query().
		Where(entuser.RoleEQ(entuser.Role(role)), entuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	return results, nil
}

func (r *userEntRepository) GetByType(ctx context.Context, userType model.UserType) ([]*model.User, error) {
	users, err := r.client.User.Query().
		Where(entuser.TypeEQ(entuser.Type(userType)), entuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	return results, nil
}

func (r *userEntRepository) Update(ctx context.Context, u *model.User) error {
	parsedID, err := uuid.Parse(u.ID)
	if err != nil {
		return err
	}

	builder := r.client.User.UpdateOneID(parsedID).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(entuser.Type(u.Type)).
		SetRole(entuser.Role(u.Role)).
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
	_, err = r.client.User.UpdateOneID(parsedID).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func (r *userEntRepository) List(ctx context.Context) ([]*model.User, error) {
	users, err := r.client.User.Query().
		Where(entuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*model.User, 0, len(users))
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	return results, nil
}

func mapEntUserToDomain(eu *ent.User) *model.User {
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
