package schema

import "entgo.io/ent"

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return nil
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return nil
}
