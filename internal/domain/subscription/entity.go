package subscription

import (
	"context"
	"time"
)

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
	ClientID         string
	ClientName       string // From Join
	PlanID           string
	PlanName         string // From Join
	Status           Status
	ShipmentStatus   ShipmentStatus
	StartDate        time.Time
	NextBillingDate  *time.Time
	NextShipmentDate *time.Time
	DaysUntilRenewal int // Computed
	CreatedAt        time.Time
}

type Filter struct {
	Search   *string
	Status   *Status
	ClientID *string
	Page     int
	PageSize int
}

type PaginatedList struct {
	Items      []*Subscription
	TotalCount int
}

// Repository defines the interface for interacting with subscription storage
type Repository interface {
	Create(ctx context.Context, sub *Subscription) error
	GetByID(ctx context.Context, id string) (*Subscription, error)
	Update(ctx context.Context, sub *Subscription) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter Filter) (*PaginatedList, error)
}
