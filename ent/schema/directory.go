package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("name").Annotations(entproto.Field(11)),
		field.Bool("is_public").Default(true).StructTag(`json:"is_public"`).Annotations(entproto.Field(12)),
		field.Int64("parent_id").Optional().StructTag(`json:"parent_id"`).Annotations(entproto.Field(13)),
	}
}

// Edges of the Directory.
func (Directory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("objects", Object.Type).Annotations(entproto.Field(21)),
		edge.To("children", Directory.Type).Annotations(entproto.Field(22)).From("parent").Annotations(entproto.Field(23)).Field("parent_id").Unique(),
	}
}

func (Directory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Directory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
