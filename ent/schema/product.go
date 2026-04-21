package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("code").Unique().NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Float("cost_price").Default(0),
		field.String("category").Optional(), // 'alimento' | 'bebida' | ...
		field.String("image_url").Optional(),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
