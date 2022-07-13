package schema

import "entgo.io/ent"

// Directory holds the schema definition for the Directory entity.
type Directory struct {
	ent.Schema
}

// Fields of the Directory.
func (Directory) Fields() []ent.Field {
	return nil
}

// Edges of the Directory.
func (Directory) Edges() []ent.Edge {
	return nil
}
