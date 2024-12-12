// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/monster"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MonsterCreate is the builder for creating a Monster entity.
type MonsterCreate struct {
	config
	mutation *MonsterMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (mc *MonsterCreate) SetName(s string) *MonsterCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetType sets the "Type" field.
func (mc *MonsterCreate) SetType(s string) *MonsterCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetHP sets the "HP" field.
func (mc *MonsterCreate) SetHP(i int) *MonsterCreate {
	mc.mutation.SetHP(i)
	return mc
}

// SetBlock sets the "Block" field.
func (mc *MonsterCreate) SetBlock(i int) *MonsterCreate {
	mc.mutation.SetBlock(i)
	return mc
}

// SetPower sets the "Power" field.
func (mc *MonsterCreate) SetPower(m map[string]int) *MonsterCreate {
	mc.mutation.SetPower(m)
	return mc
}

// SetActionName sets the "ActionName" field.
func (mc *MonsterCreate) SetActionName(s []string) *MonsterCreate {
	mc.mutation.SetActionName(s)
	return mc
}

// SetActionValue sets the "ActionValue" field.
func (mc *MonsterCreate) SetActionValue(i []int) *MonsterCreate {
	mc.mutation.SetActionValue(i)
	return mc
}

// SetImage sets the "Image" field.
func (mc *MonsterCreate) SetImage(s string) *MonsterCreate {
	mc.mutation.SetImage(s)
	return mc
}

// Mutation returns the MonsterMutation object of the builder.
func (mc *MonsterCreate) Mutation() *MonsterMutation {
	return mc.mutation
}

// Save creates the Monster in the database.
func (mc *MonsterCreate) Save(ctx context.Context) (*Monster, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MonsterCreate) SaveX(ctx context.Context) *Monster {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MonsterCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MonsterCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MonsterCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "Monster.Name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := monster.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "Monster.Name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "Monster.Type"`)}
	}
	if v, ok := mc.mutation.GetType(); ok {
		if err := monster.TypeValidator(v); err != nil {
			return &ValidationError{Name: "Type", err: fmt.Errorf(`ent: validator failed for field "Monster.Type": %w`, err)}
		}
	}
	if _, ok := mc.mutation.HP(); !ok {
		return &ValidationError{Name: "HP", err: errors.New(`ent: missing required field "Monster.HP"`)}
	}
	if _, ok := mc.mutation.Block(); !ok {
		return &ValidationError{Name: "Block", err: errors.New(`ent: missing required field "Monster.Block"`)}
	}
	if _, ok := mc.mutation.Power(); !ok {
		return &ValidationError{Name: "Power", err: errors.New(`ent: missing required field "Monster.Power"`)}
	}
	if _, ok := mc.mutation.ActionName(); !ok {
		return &ValidationError{Name: "ActionName", err: errors.New(`ent: missing required field "Monster.ActionName"`)}
	}
	if _, ok := mc.mutation.ActionValue(); !ok {
		return &ValidationError{Name: "ActionValue", err: errors.New(`ent: missing required field "Monster.ActionValue"`)}
	}
	if _, ok := mc.mutation.Image(); !ok {
		return &ValidationError{Name: "Image", err: errors.New(`ent: missing required field "Monster.Image"`)}
	}
	return nil
}

func (mc *MonsterCreate) sqlSave(ctx context.Context) (*Monster, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MonsterCreate) createSpec() (*Monster, *sqlgraph.CreateSpec) {
	var (
		_node = &Monster{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(monster.Table, sqlgraph.NewFieldSpec(monster.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(monster.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(monster.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.HP(); ok {
		_spec.SetField(monster.FieldHP, field.TypeInt, value)
		_node.HP = value
	}
	if value, ok := mc.mutation.Block(); ok {
		_spec.SetField(monster.FieldBlock, field.TypeInt, value)
		_node.Block = value
	}
	if value, ok := mc.mutation.Power(); ok {
		_spec.SetField(monster.FieldPower, field.TypeJSON, value)
		_node.Power = value
	}
	if value, ok := mc.mutation.ActionName(); ok {
		_spec.SetField(monster.FieldActionName, field.TypeJSON, value)
		_node.ActionName = value
	}
	if value, ok := mc.mutation.ActionValue(); ok {
		_spec.SetField(monster.FieldActionValue, field.TypeJSON, value)
		_node.ActionValue = value
	}
	if value, ok := mc.mutation.Image(); ok {
		_spec.SetField(monster.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	return _node, _spec
}

// MonsterCreateBulk is the builder for creating many Monster entities in bulk.
type MonsterCreateBulk struct {
	config
	err      error
	builders []*MonsterCreate
}

// Save creates the Monster entities in the database.
func (mcb *MonsterCreateBulk) Save(ctx context.Context) ([]*Monster, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Monster, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MonsterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MonsterCreateBulk) SaveX(ctx context.Context) []*Monster {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MonsterCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MonsterCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
