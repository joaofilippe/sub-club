package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Address defines the nested address structure
type Address struct {
	ZipCode      string `json:"zipCode"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("phone").Optional(),
		field.String("document").Optional().Unique(),
		field.Bool("active").Default(true),
		field.JSON("address", &Address{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriptions", Subscription.Type),
	}
}
