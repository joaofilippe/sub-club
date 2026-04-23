package model

type CreateClientInput struct {
	Name     string
	Email    string
	Phone    string
	Document string
	Active   bool
	Address  *Address
}

type UpdateClientInput struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Document string
	Active   bool
	Address  *Address
}
