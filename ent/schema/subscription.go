package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("customer_id", uuid.UUID{}).Optional(),
		field.UUID("plan_id", uuid.UUID{}).Optional(),
		field.String("status").Default("active"),            // active, paused, cancelled, expired
		field.String("shipment_status").Default("pending"), // pending, preparing, shipped, delivered
		field.Time("start_date").Default(time.Now),
		field.Time("next_billing_date").Optional(),
		field.Time("next_shipment_date").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("subscriptions").
			Unique().
			Field("customer_id"),
		edge.From("plan", Plan.Type).
			Ref("subscriptions").
			Unique().
			Field("plan_id"),
	}
}
