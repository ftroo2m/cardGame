// Code generated by ent, DO NOT EDIT.

package leaderboard

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the leaderboard type in the database.
	Label = "leaderboard"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPlayerID holds the string denoting the playerid field in the database.
	FieldPlayerID = "playerID"
	// FieldCounts holds the string denoting the counts field in the database.
	FieldCounts = "counts"
	// Table holds the table name of the leaderboard in the database.
	Table = "leaderboards"
)

// Columns holds all SQL columns for leaderboard fields.
var Columns = []string{
	FieldID,
	FieldPlayerID,
	FieldCounts,
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
	// PlayerIDValidator is a validator for the "playerID" field. It is called by the builders before save.
	PlayerIDValidator func(string) error
)

// OrderOption defines the ordering options for the Leaderboard queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPlayerID orders the results by the playerID field.
func ByPlayerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayerID, opts...).ToFunc()
}

// ByCounts orders the results by the counts field.
func ByCounts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCounts, opts...).ToFunc()
}
