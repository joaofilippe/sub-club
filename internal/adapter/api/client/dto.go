package client

type ClientDTO struct {
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

type PaginatedClientResponse struct {
	Items      []ClientDTO `json:"items"`
	TotalCount int         `json:"totalCount"`
}

// Input for both Create and Update
type ClientInputDTO struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Document string      `json:"document"`
	Active   bool        `json:"active"`
	Address  *AddressDTO `json:"address,omitempty"`
}
