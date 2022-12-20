package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Annotations(entproto.Field(11)),
		field.String("url").Annotations(entproto.Field(12)),
		field.Bool("is_public").Default(true).StructTag(`json:"is_public"`).Annotations(entproto.Field(13)),
		field.Int64("user_id").StructTag(`json:"user_id"`).Annotations(entproto.Field(14)),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("objects").Unique().Field("user_id").Required().Annotations(entproto.Field(21)),
		edge.From("directory", Directory.Type).Ref("objects").Unique().Annotations(entproto.Field(22)),
	}
}

func (Object) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Object) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
