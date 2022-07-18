package schema

import (
	"drive/tools"
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
		}),
		//field.Int64("id").Immutable().DefaultFunc(tools.GenSnowflakeID()),
		field.Int64("created_by").Default(0).StructTag(`json:"created_by"`),
		field.Int64("updated_by").Default(0).StructTag(`json:"updated_by"`),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Default(time.Time{}),
	}
}
