package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "cards"},
	}
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").StorageKey("name").NotEmpty().Unique(),
		// Type field with type string.
		field.String("Type").StorageKey("type"),
		// Energy field with type int.
		field.Int("Energy").StorageKey("energy").StructTag(`json:"Energy"`),
		// Target field with type string.
		field.String("Target").StorageKey("target"),
		// Block field with type int.
		field.Int("Block").StorageKey("block"),
		// Damage field with type int.
		field.Int("Damage").StorageKey("damage"),
		// Power field with type map[string]int, stored as JSON.
		field.JSON("Power", map[string]int{}).StorageKey("power"),
		// Action field with type []string, stored as JSON.
		field.JSON("Action", []string{}).StorageKey("action"),
		// Description field with type string.
		field.String("Description").StorageKey("description"),
		// Image field with type string.
		field.String("Image").StorageKey("image"),
		// Upgrade field with type int.
		field.Int("Upgrade").StorageKey("upgrade"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	// If Card has relationships with other entities, define them here.
	// Since there are no relationships provided in the example, return an empty slice.
	return nil
}
