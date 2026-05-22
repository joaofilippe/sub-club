package model

import "time"

type CreateAccountInput struct {
	Name          string
	Email         string
	Document      string
	Slug          string
	AccountPlanID string
}

type UpdateAccountInput struct {
	ID                    string
	Name                  string
	Email                 string
	Document              string
	AccountPlanID         string
	SubscriptionStatus    SubscriptionStatus
	SubscriptionExpiresAt *time.Time
	Active                bool
}
