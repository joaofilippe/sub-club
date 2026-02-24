package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/subclub/internal/domain/user"
)

// userPostgresRepository implements user.Repository interface
type userPostgresRepository struct {
	db *sqlx.DB
}

// NewUserPostgresRepository creates a new instance of user.Repository
func NewUserPostgresRepository(db *sqlx.DB) user.Repository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) Create(ctx context.Context, u *user.User) error {
	query := `
		INSERT INTO users (id, email, type, role) 
		VALUES ($1, $2, $3, $4)`
	// ExecContext is used so we can respect the context cancelation/timeout
	_, err := r.db.ExecContext(ctx, query, u.ID, u.Email, u.Type, u.Role)
	return err
}

func (r *userPostgresRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	var u user.User
	query := `SELECT id, email, type, role FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &u, query, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userPostgresRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	query := `SELECT id, email, type, role FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &u, query, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userPostgresRepository) GetByRole(ctx context.Context, role user.UserRole) ([]*user.User, error) {
	var users []*user.User
	query := `SELECT id, email, type, role FROM users WHERE role = $1`
	err := r.db.SelectContext(ctx, &users, query, role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userPostgresRepository) GetByType(ctx context.Context, userType user.UserType) ([]*user.User, error) {
	var users []*user.User
	query := `SELECT id, email, type, role FROM users WHERE type = $1`
	err := r.db.SelectContext(ctx, &users, query, userType)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userPostgresRepository) Update(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users 
		SET email = $1, type = $2, role = $3
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, u.Email, u.Type, u.Role, u.ID)
	return err
}

func (r *userPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userPostgresRepository) List(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	query := `SELECT id, email, type, role FROM users`
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
