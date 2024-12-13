// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/leaderboard"
	"cardGame/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeaderboardDelete is the builder for deleting a Leaderboard entity.
type LeaderboardDelete struct {
	config
	hooks    []Hook
	mutation *LeaderboardMutation
}

// Where appends a list predicates to the LeaderboardDelete builder.
func (ld *LeaderboardDelete) Where(ps ...predicate.Leaderboard) *LeaderboardDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LeaderboardDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LeaderboardDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LeaderboardDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(leaderboard.Table, sqlgraph.NewFieldSpec(leaderboard.FieldID, field.TypeInt))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LeaderboardDeleteOne is the builder for deleting a single Leaderboard entity.
type LeaderboardDeleteOne struct {
	ld *LeaderboardDelete
}

// Where appends a list predicates to the LeaderboardDelete builder.
func (ldo *LeaderboardDeleteOne) Where(ps ...predicate.Leaderboard) *LeaderboardDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LeaderboardDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{leaderboard.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LeaderboardDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}