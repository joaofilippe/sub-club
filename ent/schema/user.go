package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)


// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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
		field.UUID("account_id", uuid.UUID{}).Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
