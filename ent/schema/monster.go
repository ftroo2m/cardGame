package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Monster holds the schema definition for the Monster entity.
// Monster holds the schema definition for the Monster entity.
type Monster struct {
	ent.Schema
}

// Fields of the Monster.
func (Monster) Fields() []ent.Field {
	return []ent.Field{
		// Name field with type string and not empty constraint.
		field.String("Name").NotEmpty(),
		// HP field with type int.
		field.Int("HP"),
		// Block field with type int.
		field.Int("Block"),
		// Power field with type map[string]int.
		// Note: Ent does not support map fields directly, so you might need to use JSON or a custom type.
		field.JSON("Power", map[string]int{}),
		// Actions field with type map[string]int.
		// Note: Ent does not support map fields directly, so you might need to use JSON or a custom type.
		field.JSON("Actions", map[string]int{}),
		// Image field with type string.
		field.String("Image"),
	}
}

// Edges of the Monster.
func (Monster) Edges() []ent.Edge {
	// If Monster has relationships with other entities, define them here.
	// Since there are no relationships provided in the example, return an empty slice.
	return nil
}

// Indexes of the Monster.
func (Monster) Indexes() []ent.Index {
	// Define indexes for the Monster entity if needed.
	// For example, you might want to index the Name field to improve query performance.
	return []ent.Index{
		// Unique index on the Name field.
		index.Fields("Name").Unique(),
	}
}
