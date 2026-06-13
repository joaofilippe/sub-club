package model

import "errors"

var ErrInvalidCredentials = errors.New("invalid email or password")

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TenantLoginInput struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccountSlug string `json:"account_slug"`
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
