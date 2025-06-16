package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 添加 ID 字段
		field.Int("id").
			StorageKey("id"),
		// 定义 username 字段
		field.String("username").
			StorageKey("username").
			NotEmpty().
			Unique(),
		// 定义 password 字段
		field.String("password").
			StorageKey("password").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
