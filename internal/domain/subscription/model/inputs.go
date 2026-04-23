package model

import "time"

type CreateSubscriptionInput struct {
	ClientID         string
	PlanID           string
	Status           Status
	ShipmentStatus   ShipmentStatus
	StartDate        time.Time
	NextBillingDate  *time.Time
	NextShipmentDate *time.Time
}

type UpdateSubscriptionInput struct {
	ID               string
	Status           Status
	ShipmentStatus   ShipmentStatus
	NextBillingDate  *time.Time
	NextShipmentDate *time.Time
}
