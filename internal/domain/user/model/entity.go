package model

import "time"

type UserType string

const (
	UserTypeIndividual UserType = "individual"
	UserTypeCorporate  UserType = "corporate"
	UserTypeSystem     UserType = "system"
)

type UserRole string

const (
	UserRoleAdmin      UserRole = "admin"
	UserRoleOperations UserRole = "operations"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Type      UserType
	Role      UserRole
	AccountID *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
