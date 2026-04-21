package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("code").Unique().NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Float("product_value").Default(0),
		field.Float("discount_value").Default(0),
		field.Float("price").Default(0),
		field.Int("interval_days").Default(30),
		field.Bool("active").Default(true),
		field.String("image_url").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriptions", Subscription.Type),
	}
}
