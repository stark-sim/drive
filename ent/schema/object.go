package schema

import (
	"entgo.io/ent"
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
		field.String("name").NotEmpty(),
		field.String("url"),
		field.Bool("is_public").Default(true).StructTag(`json:"is_public"`),
		field.Int64("user_id").StructTag(`json:"user_id"`),
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("objects").Unique().Field("user_id").Required(),
		edge.From("directory", Directory.Type).Ref("objects").Unique(),
	}
}

func (Object) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
