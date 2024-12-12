// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/userconfig"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserConfigCreate is the builder for creating a UserConfig entity.
type UserConfigCreate struct {
	config
	mutation *UserConfigMutation
	hooks    []Hook
}

// SetPlayerID sets the "playerID" field.
func (ucc *UserConfigCreate) SetPlayerID(s string) *UserConfigCreate {
	ucc.mutation.SetPlayerID(s)
	return ucc
}

// SetCards sets the "cards" field.
func (ucc *UserConfigCreate) SetCards(s []string) *UserConfigCreate {
	ucc.mutation.SetCards(s)
	return ucc
}

// SetLadder sets the "ladder" field.
func (ucc *UserConfigCreate) SetLadder(s string) *UserConfigCreate {
	ucc.mutation.SetLadder(s)
	return ucc
}

// SetPlayerHP sets the "playerHP" field.
func (ucc *UserConfigCreate) SetPlayerHP(i int) *UserConfigCreate {
	ucc.mutation.SetPlayerHP(i)
	return ucc
}

// SetPlayerEnergy sets the "playerEnergy" field.
func (ucc *UserConfigCreate) SetPlayerEnergy(i int) *UserConfigCreate {
	ucc.mutation.SetPlayerEnergy(i)
	return ucc
}

// SetImage sets the "image" field.
func (ucc *UserConfigCreate) SetImage(s string) *UserConfigCreate {
	ucc.mutation.SetImage(s)
	return ucc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (ucc *UserConfigCreate) SetNillableImage(s *string) *UserConfigCreate {
	if s != nil {
		ucc.SetImage(*s)
	}
	return ucc
}

// Mutation returns the UserConfigMutation object of the builder.
func (ucc *UserConfigCreate) Mutation() *UserConfigMutation {
	return ucc.mutation
}

// Save creates the UserConfig in the database.
func (ucc *UserConfigCreate) Save(ctx context.Context) (*UserConfig, error) {
	return withHooks(ctx, ucc.sqlSave, ucc.mutation, ucc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UserConfigCreate) SaveX(ctx context.Context) *UserConfig {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UserConfigCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UserConfigCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UserConfigCreate) check() error {
	if _, ok := ucc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "playerID", err: errors.New(`ent: missing required field "UserConfig.playerID"`)}
	}
	if v, ok := ucc.mutation.PlayerID(); ok {
		if err := userconfig.PlayerIDValidator(v); err != nil {
			return &ValidationError{Name: "playerID", err: fmt.Errorf(`ent: validator failed for field "UserConfig.playerID": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.Cards(); !ok {
		return &ValidationError{Name: "cards", err: errors.New(`ent: missing required field "UserConfig.cards"`)}
	}
	if _, ok := ucc.mutation.Ladder(); !ok {
		return &ValidationError{Name: "ladder", err: errors.New(`ent: missing required field "UserConfig.ladder"`)}
	}
	if v, ok := ucc.mutation.Ladder(); ok {
		if err := userconfig.LadderValidator(v); err != nil {
			return &ValidationError{Name: "ladder", err: fmt.Errorf(`ent: validator failed for field "UserConfig.ladder": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.PlayerHP(); !ok {
		return &ValidationError{Name: "playerHP", err: errors.New(`ent: missing required field "UserConfig.playerHP"`)}
	}
	if _, ok := ucc.mutation.PlayerEnergy(); !ok {
		return &ValidationError{Name: "playerEnergy", err: errors.New(`ent: missing required field "UserConfig.playerEnergy"`)}
	}
	return nil
}

func (ucc *UserConfigCreate) sqlSave(ctx context.Context) (*UserConfig, error) {
	if err := ucc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ucc.mutation.id = &_node.ID
	ucc.mutation.done = true
	return _node, nil
}

func (ucc *UserConfigCreate) createSpec() (*UserConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &UserConfig{config: ucc.config}
		_spec = sqlgraph.NewCreateSpec(userconfig.Table, sqlgraph.NewFieldSpec(userconfig.FieldID, field.TypeInt))
	)
	if value, ok := ucc.mutation.PlayerID(); ok {
		_spec.SetField(userconfig.FieldPlayerID, field.TypeString, value)
		_node.PlayerID = value
	}
	if value, ok := ucc.mutation.Cards(); ok {
		_spec.SetField(userconfig.FieldCards, field.TypeJSON, value)
		_node.Cards = value
	}
	if value, ok := ucc.mutation.Ladder(); ok {
		_spec.SetField(userconfig.FieldLadder, field.TypeString, value)
		_node.Ladder = value
	}
	if value, ok := ucc.mutation.PlayerHP(); ok {
		_spec.SetField(userconfig.FieldPlayerHP, field.TypeInt, value)
		_node.PlayerHP = value
	}
	if value, ok := ucc.mutation.PlayerEnergy(); ok {
		_spec.SetField(userconfig.FieldPlayerEnergy, field.TypeInt, value)
		_node.PlayerEnergy = value
	}
	if value, ok := ucc.mutation.Image(); ok {
		_spec.SetField(userconfig.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	return _node, _spec
}

// UserConfigCreateBulk is the builder for creating many UserConfig entities in bulk.
type UserConfigCreateBulk struct {
	config
	err      error
	builders []*UserConfigCreate
}

// Save creates the UserConfig entities in the database.
func (uccb *UserConfigCreateBulk) Save(ctx context.Context) ([]*UserConfig, error) {
	if uccb.err != nil {
		return nil, uccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UserConfig, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UserConfigCreateBulk) SaveX(ctx context.Context) []*UserConfig {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UserConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UserConfigCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}
