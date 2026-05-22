package customer

type CustomerDTO struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Document  string      `json:"document"`
	Active    bool        `json:"active"`
	CreatedAt string      `json:"createdAt"`
	Address   *AddressDTO `json:"address,omitempty"`
}

type AddressDTO struct {
	ZipCode      string `json:"zipCode"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type PaginatedCustomerResponse struct {
	Items      []CustomerDTO `json:"items"`
	TotalCount int           `json:"totalCount"`
}

type CustomerInputDTO struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Document string      `json:"document"`
	Active   bool        `json:"active"`
	Address  *AddressDTO `json:"address,omitempty"`
}
