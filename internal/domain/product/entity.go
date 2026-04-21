package product

import (
	"time"
)

type Product struct {
	ID          string
	Code        string
	Name        string
	Description string
	CostPrice   float64
	Category    string
	ImageURL    string
	Active      bool
	CreatedAt   time.Time
}
