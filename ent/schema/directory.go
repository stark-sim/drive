package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Directory holds the schema definition for the Directory entity.
type Directory struct {
	ent.Schema
}

// Fields of the Directory.
func (Directory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("is_public").Default(true),
		field.Int64("parent_id").Default(0),
	}
}

// Edges of the Directory.
func (Directory) Edges() []ent.Edge {
	return nil
}

func (Directory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
