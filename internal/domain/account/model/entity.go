package model

import "time"

type SubscriptionStatus string

const (
	SubscriptionStatusTrial     SubscriptionStatus = "trial"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"
)

type Account struct {
	ID                    string
	Name                  string
	Email                 string
	Document              string
	Slug                  string
	AccountPlanID         string
	SubscriptionStatus    SubscriptionStatus
	SubscriptionExpiresAt *time.Time
	Active                bool
	CreatedAt             time.Time
}
