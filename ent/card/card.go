// Code generated by ent, DO NOT EDIT.

package card

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldEnergy holds the string denoting the energy field in the database.
	FieldEnergy = "energy"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldBlock holds the string denoting the block field in the database.
	FieldBlock = "block"
	// FieldDamage holds the string denoting the damage field in the database.
	FieldDamage = "damage"
	// FieldPower holds the string denoting the power field in the database.
	FieldPower = "power"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldUpgrade holds the string denoting the upgrade field in the database.
	FieldUpgrade = "upgrade"
	// Table holds the table name of the card in the database.
	Table = "cards"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldEnergy,
	FieldTarget,
	FieldBlock,
	FieldDamage,
	FieldPower,
	FieldAction,
	FieldDescription,
	FieldImage,
	FieldUpgrade,
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
)

// OrderOption defines the ordering options for the Card queries.
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

// ByEnergy orders the results by the Energy field.
func ByEnergy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnergy, opts...).ToFunc()
}

// ByTarget orders the results by the Target field.
func ByTarget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTarget, opts...).ToFunc()
}

// ByBlock orders the results by the Block field.
func ByBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlock, opts...).ToFunc()
}

// ByDamage orders the results by the Damage field.
func ByDamage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDamage, opts...).ToFunc()
}

// ByDescription orders the results by the Description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByImage orders the results by the Image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByUpgrade orders the results by the Upgrade field.
func ByUpgrade(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpgrade, opts...).ToFunc()
}
