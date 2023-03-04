package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Wechat struct {
	ent.Schema
}

func (Wechat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
	}
}

func (Wechat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Wechat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("socials", Social.Type),
	}
}
