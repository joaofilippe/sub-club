package model

import "errors"

var ErrNotFound = errors.New("tenant user not found")

type CreateTenantUserInput struct {
	Name     string
	Email    string
	Password string
	Role     Role
}

type UpdateTenantUserInput struct {
	ID   string
	Name string
	Role Role
}
