package model

import "time"

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

type Filter struct {
	Search   *string
	Category *string
	IsActive *bool
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Product
	TotalCount int
}
