package schema

import "entgo.io/ent"

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return nil
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return nil
}
