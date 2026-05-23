package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("document").Optional(),
		field.String("slug").Unique().NotEmpty(),
		field.UUID("account_plan_id", uuid.UUID{}).Optional(),
		field.Enum("subscription_status").
			Values("trial", "active", "suspended", "cancelled").
			Default("trial"),
		field.Time("subscription_expires_at").Optional().Nillable(),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account_plan", AccountPlan.Type).
			Ref("accounts").
			Unique().
			Field("account_plan_id"),
	}
}
