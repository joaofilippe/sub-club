package account

type AccountDTO struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	Email                 string  `json:"email"`
	Document              string  `json:"document"`
	Slug                  string  `json:"slug"`
	AccountPlanID         string  `json:"accountPlanId"`
	SubscriptionStatus    string  `json:"subscriptionStatus"`
	SubscriptionExpiresAt *string `json:"subscriptionExpiresAt,omitempty"`
	Active                bool    `json:"active"`
	CreatedAt             string  `json:"createdAt"`
	OwnerPassword         *string `json:"ownerPassword,omitempty"`
}

type AccountInputDTO struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Document      string `json:"document"`
	Slug          string `json:"slug"`
	AccountPlanID string `json:"accountPlanId"`
}

type UpdateAccountInputDTO struct {
	Name                  string  `json:"name"`
	Email                 string  `json:"email"`
	Document              string  `json:"document"`
	AccountPlanID         string  `json:"accountPlanId"`
	SubscriptionStatus    string  `json:"subscriptionStatus"`
	SubscriptionExpiresAt *string `json:"subscriptionExpiresAt,omitempty"`
	Active                bool    `json:"active"`
}
