package model

import "time"

type Role string

const (
	RoleAdmin      Role = "admin"
	RoleOperations Role = "operations"
)

type TenantUser struct {
	ID        string
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
