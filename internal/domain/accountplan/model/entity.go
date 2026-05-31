package model

import (
	"time"

	modulemodel "github.com/joaofilippe/subclub/internal/domain/module/model"
)

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
	Modules      []*modulemodel.Module
}
