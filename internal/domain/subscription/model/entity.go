package model

import "time"

type Status string

const (
	StatusActive    Status = "active"
	StatusPaused    Status = "paused"
	StatusCancelled Status = "cancelled"
	StatusExpired   Status = "expired"
)

type ShipmentStatus string

const (
	ShipmentPending   ShipmentStatus = "pending"
	ShipmentPreparing ShipmentStatus = "preparing"
	ShipmentShipped   ShipmentStatus = "shipped"
	ShipmentDelivered ShipmentStatus = "delivered"
)

type Subscription struct {
	ID               string
	CustomerID       string
	CustomerName     string
	PlanID           string
	PlanName         string
	Status           Status
	ShipmentStatus   ShipmentStatus
	StartDate        time.Time
	NextBillingDate  *time.Time
	NextShipmentDate *time.Time
	DaysUntilRenewal int
	CreatedAt        time.Time
}

type Filter struct {
	Search   *string
	Status   *Status
	CustomerID *string
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Subscription
	TotalCount int
}
