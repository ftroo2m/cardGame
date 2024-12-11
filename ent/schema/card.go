package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty().Unique(),
		// Type field with type string.
		field.String("Type"),
		// Energy field with type int.
		field.Int("Energy"),
		// Target field with type string.
		field.String("Target"),
		// Block field with type int.
		field.Int("Block"),
		// Damage field with type int.
		field.Int("Damage"),
		// Power field with type map[string]int, stored as JSON.
		field.JSON("Power", map[string]int{}),
		// Action field with type []string, stored as JSON.
		field.JSON("Action", []string{}),
		// Description field with type string.
		field.String("Description"),
		// Image field with type string.
		field.String("Image"),
		// Upgrade field with type int.
		field.Int("Upgrade"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	// If Card has relationships with other entities, define them here.
	// Since there are no relationships provided in the example, return an empty slice.
	return nil
}
