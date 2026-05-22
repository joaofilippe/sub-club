package model

import "errors"

var ErrInvalidCredentials = errors.New("invalid email or password")

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenOutput struct {
	Token string `json:"token"`
}
