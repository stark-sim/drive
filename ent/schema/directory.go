package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Bool("is_public").Default(true).StructTag(`json:"is_public"`),
		field.Int64("parent_id").Optional().StructTag(`json:"parent_id"`),
	}
}

// Edges of the Directory.
func (Directory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("objects", Object.Type),
		edge.To("children", Directory.Type).From("parent").Field("parent_id").Unique(),
	}
}

func (Directory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
