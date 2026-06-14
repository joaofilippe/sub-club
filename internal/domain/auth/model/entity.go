package model

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrMultipleAccountsFound = errors.New("multiple accounts found for this email — use username and account slug to login")
)

type LoginInput struct {
	Email    string
	Password string
}

type UsernameLoginInput struct {
	Username    string
	AccountSlug string
	Password    string
}

type LookupInput struct {
	Email string `json:"email"`
}

type AccountInfo struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type TokenOutput struct {
	Token string `json:"token"`
}

type TenantAuthUser struct {
	ID           string
	PasswordHash string
	Role         string
}
