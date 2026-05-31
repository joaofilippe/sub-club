package model

import "errors"

var (
	ErrModuleNotFound   = errors.New("module not found")
	ErrModuleNameExists = errors.New("module name already exists")
)
