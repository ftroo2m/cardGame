package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Leaderboard holds the schema definition for the Leaderboard entity.
type Leaderboard struct {
	ent.Schema
}

func (Leaderboard) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "leaderboards"},
	}
}

// Fields of the Leaderboard.
func (Leaderboard) Fields() []ent.Field {
	return []ent.Field{
		field.String("playerID").StorageKey("playerID").NotEmpty().Unique(),
		field.Int("counts").StorageKey("counts"),
	}
}

// Edges of the Leaderboard.
func (Leaderboard) Edges() []ent.Edge {
	return nil
}
