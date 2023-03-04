package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Social holds the schema definition for the Social entity.
type Social struct {
	ent.Schema
}

// Fields of the User.
func (Social) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown").Annotations(entgql.OrderField("NAME")),
		field.Int64("relation_id"),
		field.Int32("type"),
	}
}

// Edges of the User.
func (Social) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("email", Email.Type).Ref("socials").Unique().Field("relation_id").Required(),
		edge.From("wechat", Wechat.Type).Ref("socials").Unique().Field("relation_id").Required(),
	}
}

func (Social) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Social) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Social) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
