package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Module struct {
	ent.Schema
}

func (Module) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Unique(),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Module) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account_plans", AccountPlan.Type).
			Ref("modules"),
	}
}
