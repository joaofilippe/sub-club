package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/user"
	domain "github.com/joaofilippe/subclub/internal/domain/user"
)

type userEntRepository struct {
	client *ent.Client
}

// NewUserEntRepository creates a new instance of user.Repository
func NewUserEntRepository(client *ent.Client) domain.Repository {
	return &userEntRepository{
		client: client,
	}
}

func (r *userEntRepository) Create(ctx context.Context, u *domain.User) error {
	id, err := uuid.Parse(u.ID)
	if err != nil {
		id = uuid.New()
	}

	builder := r.client.User.Create().
		SetID(id).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(user.Type(u.Type)).
		SetRole(user.Role(u.Role)).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(u.UpdatedAt)

	if u.DeletedAt != nil {
		builder.SetDeletedAt(*u.DeletedAt)
	}

	_, err = builder.Save(ctx)
	return err
}

func (r *userEntRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	u, err := r.client.User.Query().
		Where(user.IDEQ(parsedID), user.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntUserToDomain(u), nil
}

func (r *userEntRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.client.User.Query().
		Where(user.EmailEQ(email), user.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return mapEntUserToDomain(u), nil
}

func (r *userEntRepository) GetByRole(ctx context.Context, role domain.UserRole) ([]*domain.User, error) {
	users, err := r.client.User.Query().
		Where(user.RoleEQ(user.Role(role)), user.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.User
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	if results == nil {
		results = []*domain.User{}
	}
	return results, nil
}

func (r *userEntRepository) GetByType(ctx context.Context, userType domain.UserType) ([]*domain.User, error) {
	users, err := r.client.User.Query().
		Where(user.TypeEQ(user.Type(userType)), user.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.User
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	if results == nil {
		results = []*domain.User{}
	}
	return results, nil
}

func (r *userEntRepository) Update(ctx context.Context, u *domain.User) error {
	parsedID, err := uuid.Parse(u.ID)
	if err != nil {
		return err
	}

	builder := r.client.User.UpdateOneID(parsedID).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetType(user.Type(u.Type)).
		SetRole(user.Role(u.Role)).
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
	// Soft delete
	_, err = r.client.User.UpdateOneID(parsedID).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func (r *userEntRepository) List(ctx context.Context) ([]*domain.User, error) {
	users, err := r.client.User.Query().
		Where(user.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*domain.User
	for _, eu := range users {
		results = append(results, mapEntUserToDomain(eu))
	}
	if results == nil {
		results = []*domain.User{}
	}
	return results, nil
}

func mapEntUserToDomain(eu *ent.User) *domain.User {
	return &domain.User{
		ID:        eu.ID.String(),
		Name:      eu.Name,
		Email:     eu.Email,
		Password:  eu.Password,
		Type:      domain.UserType(eu.Type),
		Role:      domain.UserRole(eu.Role),
		CreatedAt: eu.CreatedAt,
		UpdatedAt: eu.UpdatedAt,
		DeletedAt: eu.DeletedAt,
	}
}
