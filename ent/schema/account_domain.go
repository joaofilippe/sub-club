package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountDomain maps an email domain (e.g. "empresa.com") to an Account,
// enabling fast tenant lookup without scanning every tenant schema.
type AccountDomain struct {
	ent.Schema
}

func (AccountDomain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("domain").Unique().NotEmpty(),
		field.UUID("account_id", uuid.UUID{}),
	}
}

func (AccountDomain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("domains").
			Unique().
			Required().
			Field("account_id"),
	}
}
