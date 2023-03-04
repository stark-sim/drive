package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Email struct {
	ent.Schema
}

func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
	}
}

func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("socials", Social.Type),
	}
}

func (Email) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
