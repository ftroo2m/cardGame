package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Leaderboard holds the schema definition for the Leaderboard entity.
type Leaderboard struct {
	ent.Schema
}

// Fields of the Leaderboard.
func (Leaderboard) Fields() []ent.Field {
	return []ent.Field{
		field.String("playerID").NotEmpty().Unique(),
		field.Int("counts"),
	}
}

// Edges of the Leaderboard.
func (Leaderboard) Edges() []ent.Edge {
	return nil
}
