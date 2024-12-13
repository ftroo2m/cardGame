package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Monster holds the schema definition for the Monster entity.
// Monster holds the schema definition for the Monster entity.
type Monster struct {
	ent.Schema
}

func (Monster) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "monsters"},
	}
}

// Fields of the Monster.
func (Monster) Fields() []ent.Field {
	return []ent.Field{
		// Name field with type string and not empty constraint.
		field.String("Name").StorageKey("name").NotEmpty().Unique(),
		field.String("Type").StorageKey("Type").NotEmpty(),
		// HP field with type int.
		field.Int("HP").StorageKey("HP"),
		// Block field with type int.
		field.Int("Block").StorageKey("block"),
		// Power field with type map[string]int.
		// Note: Ent does not support map fields directly, so you might need to use JSON or a custom type.
		field.JSON("Power", map[string]int{}).StorageKey("power"),
		// Actions field with type map[string]int.
		// Note: Ent does not support map fields directly, so you might need to use JSON or a custom type.
		field.JSON("ActionName", []string{}).StorageKey("actionName"),
		field.JSON("ActionValue", []int{}).StorageKey("actionValue"),
		// Image field with type string.
		field.String("Image").StorageKey("image"),
	}
}

// Edges of the Monster.
func (Monster) Edges() []ent.Edge {
	// If Monster has relationships with other entities, define them here.
	// Since there are no relationships provided in the example, return an empty slice.
	return nil
}
