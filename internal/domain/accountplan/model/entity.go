package model

import "time"

type AccountPlan struct {
	ID           string
	Name         string
	Description  string
	Price        float64
	MaxCustomers int
	MaxPlans     int
	MaxProducts  int
	Active       bool
	CreatedAt    time.Time
}
