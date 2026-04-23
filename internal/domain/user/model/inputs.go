package model

type CreateUserInput struct {
	Email    string   `json:"email"`
	Type     UserType `json:"type"`
	Role     UserRole `json:"role"`
}

type UpdateUserInput struct {
	ID    string
	Name  string
	Email string
	Type  UserType
	Role  UserRole
}
