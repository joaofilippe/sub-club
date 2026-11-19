package accountplan

type AccountPlanDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	MaxCustomers int     `json:"maxCustomers"`
	MaxPlans     int     `json:"maxPlans"`
	MaxProducts  int     `json:"maxProducts"`
	Active       bool    `json:"active"`
	CreatedAt    string  `json:"createdAt"`
}

type AccountPlanInputDTO struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	MaxCustomers int     `json:"maxCustomers"`
	MaxPlans     int     `json:"maxPlans"`
	MaxProducts  int     `json:"maxProducts"`
	Active       bool    `json:"active"`
}
