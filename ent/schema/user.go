package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 定义 username 字段，类型为字符串，不允许为空。
		field.String("username").NotEmpty(),
		// 定义 password 字段，类型为字符串，不允许为空。
		field.String("password").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
