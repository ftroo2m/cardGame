// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/card"
	"cardGame/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "Name" field.
func (cu *CardUpdate) SetName(s string) *CardUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(s *string) *CardUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetType sets the "Type" field.
func (cu *CardUpdate) SetType(s string) *CardUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetNillableType sets the "Type" field if the given value is not nil.
func (cu *CardUpdate) SetNillableType(s *string) *CardUpdate {
	if s != nil {
		cu.SetType(*s)
	}
	return cu
}

// SetEnergy sets the "Energy" field.
func (cu *CardUpdate) SetEnergy(i int) *CardUpdate {
	cu.mutation.ResetEnergy()
	cu.mutation.SetEnergy(i)
	return cu
}

// SetNillableEnergy sets the "Energy" field if the given value is not nil.
func (cu *CardUpdate) SetNillableEnergy(i *int) *CardUpdate {
	if i != nil {
		cu.SetEnergy(*i)
	}
	return cu
}

// AddEnergy adds i to the "Energy" field.
func (cu *CardUpdate) AddEnergy(i int) *CardUpdate {
	cu.mutation.AddEnergy(i)
	return cu
}

// SetTarget sets the "Target" field.
func (cu *CardUpdate) SetTarget(s string) *CardUpdate {
	cu.mutation.SetTarget(s)
	return cu
}

// SetNillableTarget sets the "Target" field if the given value is not nil.
func (cu *CardUpdate) SetNillableTarget(s *string) *CardUpdate {
	if s != nil {
		cu.SetTarget(*s)
	}
	return cu
}

// SetBlock sets the "Block" field.
func (cu *CardUpdate) SetBlock(i int) *CardUpdate {
	cu.mutation.ResetBlock()
	cu.mutation.SetBlock(i)
	return cu
}

// SetNillableBlock sets the "Block" field if the given value is not nil.
func (cu *CardUpdate) SetNillableBlock(i *int) *CardUpdate {
	if i != nil {
		cu.SetBlock(*i)
	}
	return cu
}

// AddBlock adds i to the "Block" field.
func (cu *CardUpdate) AddBlock(i int) *CardUpdate {
	cu.mutation.AddBlock(i)
	return cu
}

// SetDamage sets the "Damage" field.
func (cu *CardUpdate) SetDamage(i int) *CardUpdate {
	cu.mutation.ResetDamage()
	cu.mutation.SetDamage(i)
	return cu
}

// SetNillableDamage sets the "Damage" field if the given value is not nil.
func (cu *CardUpdate) SetNillableDamage(i *int) *CardUpdate {
	if i != nil {
		cu.SetDamage(*i)
	}
	return cu
}

// AddDamage adds i to the "Damage" field.
func (cu *CardUpdate) AddDamage(i int) *CardUpdate {
	cu.mutation.AddDamage(i)
	return cu
}

// SetPower sets the "Power" field.
func (cu *CardUpdate) SetPower(m map[string]int) *CardUpdate {
	cu.mutation.SetPower(m)
	return cu
}

// SetAction sets the "Action" field.
func (cu *CardUpdate) SetAction(s []string) *CardUpdate {
	cu.mutation.SetAction(s)
	return cu
}

// AppendAction appends s to the "Action" field.
func (cu *CardUpdate) AppendAction(s []string) *CardUpdate {
	cu.mutation.AppendAction(s)
	return cu
}

// SetDescription sets the "Description" field.
func (cu *CardUpdate) SetDescription(s string) *CardUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (cu *CardUpdate) SetNillableDescription(s *string) *CardUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// SetImage sets the "Image" field.
func (cu *CardUpdate) SetImage(s string) *CardUpdate {
	cu.mutation.SetImage(s)
	return cu
}

// SetNillableImage sets the "Image" field if the given value is not nil.
func (cu *CardUpdate) SetNillableImage(s *string) *CardUpdate {
	if s != nil {
		cu.SetImage(*s)
	}
	return cu
}

// SetUpgrade sets the "Upgrade" field.
func (cu *CardUpdate) SetUpgrade(i int) *CardUpdate {
	cu.mutation.ResetUpgrade()
	cu.mutation.SetUpgrade(i)
	return cu
}

// SetNillableUpgrade sets the "Upgrade" field if the given value is not nil.
func (cu *CardUpdate) SetNillableUpgrade(i *int) *CardUpdate {
	if i != nil {
		cu.SetUpgrade(*i)
	}
	return cu
}

// AddUpgrade adds i to the "Upgrade" field.
func (cu *CardUpdate) AddUpgrade(i int) *CardUpdate {
	cu.mutation.AddUpgrade(i)
	return cu
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "Card.Name": %w`, err)}
		}
	}
	return nil
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(card.FieldType, field.TypeString, value)
	}
	if value, ok := cu.mutation.Energy(); ok {
		_spec.SetField(card.FieldEnergy, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedEnergy(); ok {
		_spec.AddField(card.FieldEnergy, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Target(); ok {
		_spec.SetField(card.FieldTarget, field.TypeString, value)
	}
	if value, ok := cu.mutation.Block(); ok {
		_spec.SetField(card.FieldBlock, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedBlock(); ok {
		_spec.AddField(card.FieldBlock, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Damage(); ok {
		_spec.SetField(card.FieldDamage, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedDamage(); ok {
		_spec.AddField(card.FieldDamage, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Power(); ok {
		_spec.SetField(card.FieldPower, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.Action(); ok {
		_spec.SetField(card.FieldAction, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedAction(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, card.FieldAction, value)
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(card.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.Image(); ok {
		_spec.SetField(card.FieldImage, field.TypeString, value)
	}
	if value, ok := cu.mutation.Upgrade(); ok {
		_spec.SetField(card.FieldUpgrade, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedUpgrade(); ok {
		_spec.AddField(card.FieldUpgrade, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetName sets the "Name" field.
func (cuo *CardUpdateOne) SetName(s string) *CardUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetType sets the "Type" field.
func (cuo *CardUpdateOne) SetType(s string) *CardUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetNillableType sets the "Type" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableType(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetType(*s)
	}
	return cuo
}

// SetEnergy sets the "Energy" field.
func (cuo *CardUpdateOne) SetEnergy(i int) *CardUpdateOne {
	cuo.mutation.ResetEnergy()
	cuo.mutation.SetEnergy(i)
	return cuo
}

// SetNillableEnergy sets the "Energy" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableEnergy(i *int) *CardUpdateOne {
	if i != nil {
		cuo.SetEnergy(*i)
	}
	return cuo
}

// AddEnergy adds i to the "Energy" field.
func (cuo *CardUpdateOne) AddEnergy(i int) *CardUpdateOne {
	cuo.mutation.AddEnergy(i)
	return cuo
}

// SetTarget sets the "Target" field.
func (cuo *CardUpdateOne) SetTarget(s string) *CardUpdateOne {
	cuo.mutation.SetTarget(s)
	return cuo
}

// SetNillableTarget sets the "Target" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableTarget(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetTarget(*s)
	}
	return cuo
}

// SetBlock sets the "Block" field.
func (cuo *CardUpdateOne) SetBlock(i int) *CardUpdateOne {
	cuo.mutation.ResetBlock()
	cuo.mutation.SetBlock(i)
	return cuo
}

// SetNillableBlock sets the "Block" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableBlock(i *int) *CardUpdateOne {
	if i != nil {
		cuo.SetBlock(*i)
	}
	return cuo
}

// AddBlock adds i to the "Block" field.
func (cuo *CardUpdateOne) AddBlock(i int) *CardUpdateOne {
	cuo.mutation.AddBlock(i)
	return cuo
}

// SetDamage sets the "Damage" field.
func (cuo *CardUpdateOne) SetDamage(i int) *CardUpdateOne {
	cuo.mutation.ResetDamage()
	cuo.mutation.SetDamage(i)
	return cuo
}

// SetNillableDamage sets the "Damage" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableDamage(i *int) *CardUpdateOne {
	if i != nil {
		cuo.SetDamage(*i)
	}
	return cuo
}

// AddDamage adds i to the "Damage" field.
func (cuo *CardUpdateOne) AddDamage(i int) *CardUpdateOne {
	cuo.mutation.AddDamage(i)
	return cuo
}

// SetPower sets the "Power" field.
func (cuo *CardUpdateOne) SetPower(m map[string]int) *CardUpdateOne {
	cuo.mutation.SetPower(m)
	return cuo
}

// SetAction sets the "Action" field.
func (cuo *CardUpdateOne) SetAction(s []string) *CardUpdateOne {
	cuo.mutation.SetAction(s)
	return cuo
}

// AppendAction appends s to the "Action" field.
func (cuo *CardUpdateOne) AppendAction(s []string) *CardUpdateOne {
	cuo.mutation.AppendAction(s)
	return cuo
}

// SetDescription sets the "Description" field.
func (cuo *CardUpdateOne) SetDescription(s string) *CardUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "Description" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableDescription(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// SetImage sets the "Image" field.
func (cuo *CardUpdateOne) SetImage(s string) *CardUpdateOne {
	cuo.mutation.SetImage(s)
	return cuo
}

// SetNillableImage sets the "Image" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableImage(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetImage(*s)
	}
	return cuo
}

// SetUpgrade sets the "Upgrade" field.
func (cuo *CardUpdateOne) SetUpgrade(i int) *CardUpdateOne {
	cuo.mutation.ResetUpgrade()
	cuo.mutation.SetUpgrade(i)
	return cuo
}

// SetNillableUpgrade sets the "Upgrade" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableUpgrade(i *int) *CardUpdateOne {
	if i != nil {
		cuo.SetUpgrade(*i)
	}
	return cuo
}

// AddUpgrade adds i to the "Upgrade" field.
func (cuo *CardUpdateOne) AddUpgrade(i int) *CardUpdateOne {
	cuo.mutation.AddUpgrade(i)
	return cuo
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cuo *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "Card.Name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(card.FieldType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Energy(); ok {
		_spec.SetField(card.FieldEnergy, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedEnergy(); ok {
		_spec.AddField(card.FieldEnergy, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Target(); ok {
		_spec.SetField(card.FieldTarget, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Block(); ok {
		_spec.SetField(card.FieldBlock, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedBlock(); ok {
		_spec.AddField(card.FieldBlock, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Damage(); ok {
		_spec.SetField(card.FieldDamage, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedDamage(); ok {
		_spec.AddField(card.FieldDamage, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Power(); ok {
		_spec.SetField(card.FieldPower, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.Action(); ok {
		_spec.SetField(card.FieldAction, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedAction(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, card.FieldAction, value)
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(card.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Image(); ok {
		_spec.SetField(card.FieldImage, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Upgrade(); ok {
		_spec.SetField(card.FieldUpgrade, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedUpgrade(); ok {
		_spec.AddField(card.FieldUpgrade, field.TypeInt, value)
	}
	_node = &Card{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
