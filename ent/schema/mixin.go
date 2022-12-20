package schema

import (
	"drive/tools"
	"entgo.io/contrib/entproto"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().DefaultFunc(func() int64 {
			return tools.GenSnowflakeID()
		}).Annotations(entproto.Field(1)),
		//field.Int64("id").Immutable().DefaultFunc(tools.GenSnowflakeID()),
		field.Int64("created_by").Default(0).StructTag(`json:"created_by"`).Annotations(entproto.Field(2)),
		field.Int64("updated_by").Default(0).StructTag(`json:"updated_by"`).Annotations(entproto.Field(3)),
		field.Time("created_at").Immutable().Default(time.Now).Annotations(entproto.Field(4)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entproto.Field(5)),
		field.Time("deleted_at").Default(time.Time{}).Annotations(entproto.Field(6)),
	}
}
