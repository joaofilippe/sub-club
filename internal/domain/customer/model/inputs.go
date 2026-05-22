package model

type CreateCustomerInput struct {
	Name     string
	Email    string
	Phone    string
	Document string
	Active   bool
	Address  *Address
}

type UpdateCustomerInput struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Document string
	Active   bool
	Address  *Address
}
