package model

type CreateProductInput struct {
	Code        string
	Name        string
	Description string
	CostPrice   float64
	Category    string
	ImageURL    string
	Active      bool
}

type UpdateProductInput struct {
	ID          string
	Code        string
	Name        string
	Description string
	CostPrice   float64
	Category    string
	ImageURL    string
	Active      bool
}
