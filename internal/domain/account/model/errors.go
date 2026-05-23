package model

import "errors"

var (
	ErrNotFound  = errors.New("account not found")
	ErrSlugTaken = errors.New("account slug already in use")
)
