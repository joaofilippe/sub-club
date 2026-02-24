package user

import "context"

// Repository defines the contract for user data persistence.
type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByRole(ctx context.Context, role UserRole) ([]*User, error)
	GetByType(ctx context.Context, userType UserType) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*User, error)
}
