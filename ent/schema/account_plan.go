package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type AccountPlan struct {
	ent.Schema
}

func (AccountPlan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Float("price"),
		field.Int("max_customers").Default(0),
		field.Int("max_plans").Default(0),
		field.Int("max_products").Default(0),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (AccountPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
	}
}
