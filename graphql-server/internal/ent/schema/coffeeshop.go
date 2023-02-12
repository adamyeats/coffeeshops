package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Coffeeshop holds the schema definition for the Coffeeshop entity.
type Coffeeshop struct {
	ent.Schema
}

func (Coffeeshop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Coffeeshop"},
	}
}

// Fields of the Coffeeshop.
func (Coffeeshop) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("address"),
		field.Float("latitude"),
		field.Float("longitude"),
	}
}

// Edges of the Coffeeshop.
func (Coffeeshop) Edges() []ent.Edge {
	return nil
}
