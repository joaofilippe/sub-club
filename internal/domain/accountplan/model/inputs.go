package model

type CreateAccountPlanInput struct {
	Name         string
	Description  string
	Price        float64
	MaxCustomers int
	MaxPlans     int
	MaxProducts  int
	ModuleIDs    []string
}

type UpdateAccountPlanInput struct {
	ID           string
	Name         string
	Description  string
	Price        float64
	MaxCustomers int
	MaxPlans     int
	MaxProducts  int
	Active       bool
	ModuleIDs    []string
}
