package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

func (UserConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "userconfigs"},
	}
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("playerID").StorageKey("playerID").
			NotEmpty().Unique(),
		field.Strings("cards").StorageKey("cards"),
		field.String("ladder").StorageKey("ladder").
			NotEmpty(),
		field.Int("playerHP").StorageKey("playerHP"),
		field.Int("playerEnergy").StorageKey("playerEnergy"),
		field.String("image").StorageKey("image").
			Optional(),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return nil
}
