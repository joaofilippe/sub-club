package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SystemUser holds platform-level administrators that live in the public schema.
type SystemUser struct {
	ent.Schema
}

func (SystemUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.Enum("type").
			Values("individual", "corporate", "system").
			Default("individual"),
		field.Enum("role").
			Values("admin", "operations").
			Default("admin"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SystemUser) Edges() []ent.Edge {
	return nil
}
