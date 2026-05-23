package model

type CreateAccountPlanInput struct {
	Name         string
	Description  string
	Price        float64
	MaxCustomers int
	MaxPlans     int
	MaxProducts  int
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
}
