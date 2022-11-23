package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("user").Unique(),
		field.String("token"),
		field.Time("created").Default(time.Now()),
		field.Uint64("permissions"),
		field.String("group"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return nil
}
