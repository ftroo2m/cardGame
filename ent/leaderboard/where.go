// Code generated by ent, DO NOT EDIT.

package leaderboard

import (
	"cardGame/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLTE(FieldID, id))
}

// PlayerID applies equality check predicate on the "playerID" field. It's identical to PlayerIDEQ.
func PlayerID(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldPlayerID, v))
}

// Counts applies equality check predicate on the "counts" field. It's identical to CountsEQ.
func Counts(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldCounts, v))
}

// PlayerIDEQ applies the EQ predicate on the "playerID" field.
func PlayerIDEQ(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldPlayerID, v))
}

// PlayerIDNEQ applies the NEQ predicate on the "playerID" field.
func PlayerIDNEQ(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNEQ(FieldPlayerID, v))
}

// PlayerIDIn applies the In predicate on the "playerID" field.
func PlayerIDIn(vs ...string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldIn(FieldPlayerID, vs...))
}

// PlayerIDNotIn applies the NotIn predicate on the "playerID" field.
func PlayerIDNotIn(vs ...string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNotIn(FieldPlayerID, vs...))
}

// PlayerIDGT applies the GT predicate on the "playerID" field.
func PlayerIDGT(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGT(FieldPlayerID, v))
}

// PlayerIDGTE applies the GTE predicate on the "playerID" field.
func PlayerIDGTE(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGTE(FieldPlayerID, v))
}

// PlayerIDLT applies the LT predicate on the "playerID" field.
func PlayerIDLT(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLT(FieldPlayerID, v))
}

// PlayerIDLTE applies the LTE predicate on the "playerID" field.
func PlayerIDLTE(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLTE(FieldPlayerID, v))
}

// PlayerIDContains applies the Contains predicate on the "playerID" field.
func PlayerIDContains(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldContains(FieldPlayerID, v))
}

// PlayerIDHasPrefix applies the HasPrefix predicate on the "playerID" field.
func PlayerIDHasPrefix(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldHasPrefix(FieldPlayerID, v))
}

// PlayerIDHasSuffix applies the HasSuffix predicate on the "playerID" field.
func PlayerIDHasSuffix(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldHasSuffix(FieldPlayerID, v))
}

// PlayerIDEqualFold applies the EqualFold predicate on the "playerID" field.
func PlayerIDEqualFold(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEqualFold(FieldPlayerID, v))
}

// PlayerIDContainsFold applies the ContainsFold predicate on the "playerID" field.
func PlayerIDContainsFold(v string) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldContainsFold(FieldPlayerID, v))
}

// CountsEQ applies the EQ predicate on the "counts" field.
func CountsEQ(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldEQ(FieldCounts, v))
}

// CountsNEQ applies the NEQ predicate on the "counts" field.
func CountsNEQ(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNEQ(FieldCounts, v))
}

// CountsIn applies the In predicate on the "counts" field.
func CountsIn(vs ...int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldIn(FieldCounts, vs...))
}

// CountsNotIn applies the NotIn predicate on the "counts" field.
func CountsNotIn(vs ...int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldNotIn(FieldCounts, vs...))
}

// CountsGT applies the GT predicate on the "counts" field.
func CountsGT(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGT(FieldCounts, v))
}

// CountsGTE applies the GTE predicate on the "counts" field.
func CountsGTE(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldGTE(FieldCounts, v))
}

// CountsLT applies the LT predicate on the "counts" field.
func CountsLT(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLT(FieldCounts, v))
}

// CountsLTE applies the LTE predicate on the "counts" field.
func CountsLTE(v int) predicate.Leaderboard {
	return predicate.Leaderboard(sql.FieldLTE(FieldCounts, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Leaderboard) predicate.Leaderboard {
	return predicate.Leaderboard(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Leaderboard) predicate.Leaderboard {
	return predicate.Leaderboard(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Leaderboard) predicate.Leaderboard {
	return predicate.Leaderboard(sql.NotPredicates(p))
}