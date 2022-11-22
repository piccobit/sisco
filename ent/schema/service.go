package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Optional().Default(""),
		field.String("protocol"),
		field.String("host"),
		field.String("port"),
		field.Bool("available").Default(false),
		field.Time("heartbeat").Default(time.Now()),
		field.String("owner"),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.From("area", Area.Type).Ref("services").Unique(),
	}
}
