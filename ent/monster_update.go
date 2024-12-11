// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/monster"
	"cardGame/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MonsterUpdate is the builder for updating Monster entities.
type MonsterUpdate struct {
	config
	hooks    []Hook
	mutation *MonsterMutation
}

// Where appends a list predicates to the MonsterUpdate builder.
func (mu *MonsterUpdate) Where(ps ...predicate.Monster) *MonsterUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "Name" field.
func (mu *MonsterUpdate) SetName(s string) *MonsterUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (mu *MonsterUpdate) SetNillableName(s *string) *MonsterUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetHP sets the "HP" field.
func (mu *MonsterUpdate) SetHP(i int) *MonsterUpdate {
	mu.mutation.ResetHP()
	mu.mutation.SetHP(i)
	return mu
}

// SetNillableHP sets the "HP" field if the given value is not nil.
func (mu *MonsterUpdate) SetNillableHP(i *int) *MonsterUpdate {
	if i != nil {
		mu.SetHP(*i)
	}
	return mu
}

// AddHP adds i to the "HP" field.
func (mu *MonsterUpdate) AddHP(i int) *MonsterUpdate {
	mu.mutation.AddHP(i)
	return mu
}

// SetBlock sets the "Block" field.
func (mu *MonsterUpdate) SetBlock(i int) *MonsterUpdate {
	mu.mutation.ResetBlock()
	mu.mutation.SetBlock(i)
	return mu
}

// SetNillableBlock sets the "Block" field if the given value is not nil.
func (mu *MonsterUpdate) SetNillableBlock(i *int) *MonsterUpdate {
	if i != nil {
		mu.SetBlock(*i)
	}
	return mu
}

// AddBlock adds i to the "Block" field.
func (mu *MonsterUpdate) AddBlock(i int) *MonsterUpdate {
	mu.mutation.AddBlock(i)
	return mu
}

// SetPower sets the "Power" field.
func (mu *MonsterUpdate) SetPower(m map[string]int) *MonsterUpdate {
	mu.mutation.SetPower(m)
	return mu
}

// SetActions sets the "Actions" field.
func (mu *MonsterUpdate) SetActions(m map[string]int) *MonsterUpdate {
	mu.mutation.SetActions(m)
	return mu
}

// SetImage sets the "Image" field.
func (mu *MonsterUpdate) SetImage(s string) *MonsterUpdate {
	mu.mutation.SetImage(s)
	return mu
}

// SetNillableImage sets the "Image" field if the given value is not nil.
func (mu *MonsterUpdate) SetNillableImage(s *string) *MonsterUpdate {
	if s != nil {
		mu.SetImage(*s)
	}
	return mu
}

// Mutation returns the MonsterMutation object of the builder.
func (mu *MonsterUpdate) Mutation() *MonsterMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MonsterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MonsterUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MonsterUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MonsterUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MonsterUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := monster.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "Monster.Name": %w`, err)}
		}
	}
	return nil
}

func (mu *MonsterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(monster.Table, monster.Columns, sqlgraph.NewFieldSpec(monster.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(monster.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.HP(); ok {
		_spec.SetField(monster.FieldHP, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedHP(); ok {
		_spec.AddField(monster.FieldHP, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Block(); ok {
		_spec.SetField(monster.FieldBlock, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedBlock(); ok {
		_spec.AddField(monster.FieldBlock, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Power(); ok {
		_spec.SetField(monster.FieldPower, field.TypeJSON, value)
	}
	if value, ok := mu.mutation.Actions(); ok {
		_spec.SetField(monster.FieldActions, field.TypeJSON, value)
	}
	if value, ok := mu.mutation.Image(); ok {
		_spec.SetField(monster.FieldImage, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{monster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MonsterUpdateOne is the builder for updating a single Monster entity.
type MonsterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MonsterMutation
}

// SetName sets the "Name" field.
func (muo *MonsterUpdateOne) SetName(s string) *MonsterUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (muo *MonsterUpdateOne) SetNillableName(s *string) *MonsterUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetHP sets the "HP" field.
func (muo *MonsterUpdateOne) SetHP(i int) *MonsterUpdateOne {
	muo.mutation.ResetHP()
	muo.mutation.SetHP(i)
	return muo
}

// SetNillableHP sets the "HP" field if the given value is not nil.
func (muo *MonsterUpdateOne) SetNillableHP(i *int) *MonsterUpdateOne {
	if i != nil {
		muo.SetHP(*i)
	}
	return muo
}

// AddHP adds i to the "HP" field.
func (muo *MonsterUpdateOne) AddHP(i int) *MonsterUpdateOne {
	muo.mutation.AddHP(i)
	return muo
}

// SetBlock sets the "Block" field.
func (muo *MonsterUpdateOne) SetBlock(i int) *MonsterUpdateOne {
	muo.mutation.ResetBlock()
	muo.mutation.SetBlock(i)
	return muo
}

// SetNillableBlock sets the "Block" field if the given value is not nil.
func (muo *MonsterUpdateOne) SetNillableBlock(i *int) *MonsterUpdateOne {
	if i != nil {
		muo.SetBlock(*i)
	}
	return muo
}

// AddBlock adds i to the "Block" field.
func (muo *MonsterUpdateOne) AddBlock(i int) *MonsterUpdateOne {
	muo.mutation.AddBlock(i)
	return muo
}

// SetPower sets the "Power" field.
func (muo *MonsterUpdateOne) SetPower(m map[string]int) *MonsterUpdateOne {
	muo.mutation.SetPower(m)
	return muo
}

// SetActions sets the "Actions" field.
func (muo *MonsterUpdateOne) SetActions(m map[string]int) *MonsterUpdateOne {
	muo.mutation.SetActions(m)
	return muo
}

// SetImage sets the "Image" field.
func (muo *MonsterUpdateOne) SetImage(s string) *MonsterUpdateOne {
	muo.mutation.SetImage(s)
	return muo
}

// SetNillableImage sets the "Image" field if the given value is not nil.
func (muo *MonsterUpdateOne) SetNillableImage(s *string) *MonsterUpdateOne {
	if s != nil {
		muo.SetImage(*s)
	}
	return muo
}

// Mutation returns the MonsterMutation object of the builder.
func (muo *MonsterUpdateOne) Mutation() *MonsterMutation {
	return muo.mutation
}

// Where appends a list predicates to the MonsterUpdate builder.
func (muo *MonsterUpdateOne) Where(ps ...predicate.Monster) *MonsterUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MonsterUpdateOne) Select(field string, fields ...string) *MonsterUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Monster entity.
func (muo *MonsterUpdateOne) Save(ctx context.Context) (*Monster, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MonsterUpdateOne) SaveX(ctx context.Context) *Monster {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MonsterUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MonsterUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MonsterUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := monster.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "Monster.Name": %w`, err)}
		}
	}
	return nil
}

func (muo *MonsterUpdateOne) sqlSave(ctx context.Context) (_node *Monster, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(monster.Table, monster.Columns, sqlgraph.NewFieldSpec(monster.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Monster.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, monster.FieldID)
		for _, f := range fields {
			if !monster.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != monster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(monster.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.HP(); ok {
		_spec.SetField(monster.FieldHP, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedHP(); ok {
		_spec.AddField(monster.FieldHP, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Block(); ok {
		_spec.SetField(monster.FieldBlock, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedBlock(); ok {
		_spec.AddField(monster.FieldBlock, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Power(); ok {
		_spec.SetField(monster.FieldPower, field.TypeJSON, value)
	}
	if value, ok := muo.mutation.Actions(); ok {
		_spec.SetField(monster.FieldActions, field.TypeJSON, value)
	}
	if value, ok := muo.mutation.Image(); ok {
		_spec.SetField(monster.FieldImage, field.TypeString, value)
	}
	_node = &Monster{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{monster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}