package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description"),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
	}
}
