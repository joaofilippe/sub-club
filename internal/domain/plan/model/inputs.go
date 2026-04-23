package model

type CreatePlanInput struct {
	Code          string
	Name          string
	Description   string
	ProductValue  float64
	DiscountValue float64
	Price         float64
	IntervalDays  int
	Active        bool
	ImageURL      string
}

type UpdatePlanInput struct {
	ID            string
	Code          string
	Name          string
	Description   string
	ProductValue  float64
	DiscountValue float64
	Price         float64
	IntervalDays  int
	Active        bool
	ImageURL      string
}
