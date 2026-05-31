package model

type CreateModuleInput struct {
	Name string
}

type UpdateModuleInput struct {
	ID     string
	Name   string
	Active bool
}
