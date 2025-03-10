// Code generated by ent, DO NOT EDIT.

package monster

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the monster type in the database.
	Label = "monster"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "Type"
	// FieldHP holds the string denoting the hp field in the database.
	FieldHP = "HP"
	// FieldBlock holds the string denoting the block field in the database.
	FieldBlock = "block"
	// FieldPower holds the string denoting the power field in the database.
	FieldPower = "power"
	// FieldActionName holds the string denoting the actionname field in the database.
	FieldActionName = "actionName"
	// FieldActionValue holds the string denoting the actionvalue field in the database.
	FieldActionValue = "actionValue"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// Table holds the table name of the monster in the database.
	Table = "monsters"
)

// Columns holds all SQL columns for monster fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldHP,
	FieldBlock,
	FieldPower,
	FieldActionName,
	FieldActionValue,
	FieldImage,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TypeValidator is a validator for the "Type" field. It is called by the builders before save.
	TypeValidator func(string) error
)

// OrderOption defines the ordering options for the Monster queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the Type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByHP orders the results by the HP field.
func ByHP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHP, opts...).ToFunc()
}

// ByBlock orders the results by the Block field.
func ByBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlock, opts...).ToFunc()
}

// ByImage orders the results by the Image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}
