package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("playerID").
			NotEmpty().Unique(),
		field.Strings("cards"),
		field.String("ladder").
			NotEmpty(),
		field.Int("playerHP"),
		field.Int("playerEnergy"),
		field.String("image").
			Optional(),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return nil
}
