package user

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
	UserRoleCustomer   UserRole = "customer"
	UserRoleOperations UserRole = "operations"
)

type User struct {
	ID        string
	Email     string
	Type      UserType
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
