// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/monster"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Monster is the model entity for the Monster schema.
type Monster struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// HP holds the value of the "HP" field.
	HP int `json:"HP,omitempty"`
	// Block holds the value of the "Block" field.
	Block int `json:"Block,omitempty"`
	// Power holds the value of the "Power" field.
	Power map[string]int `json:"Power,omitempty"`
	// Actions holds the value of the "Actions" field.
	Actions map[string]int `json:"Actions,omitempty"`
	// Image holds the value of the "Image" field.
	Image        string `json:"Image,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Monster) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case monster.FieldPower, monster.FieldActions:
			values[i] = new([]byte)
		case monster.FieldID, monster.FieldHP, monster.FieldBlock:
			values[i] = new(sql.NullInt64)
		case monster.FieldName, monster.FieldImage:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Monster fields.
func (m *Monster) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case monster.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case monster.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case monster.FieldHP:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field HP", values[i])
			} else if value.Valid {
				m.HP = int(value.Int64)
			}
		case monster.FieldBlock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Block", values[i])
			} else if value.Valid {
				m.Block = int(value.Int64)
			}
		case monster.FieldPower:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field Power", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Power); err != nil {
					return fmt.Errorf("unmarshal field Power: %w", err)
				}
			}
		case monster.FieldActions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field Actions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Actions); err != nil {
					return fmt.Errorf("unmarshal field Actions: %w", err)
				}
			}
		case monster.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Image", values[i])
			} else if value.Valid {
				m.Image = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Monster.
// This includes values selected through modifiers, order, etc.
func (m *Monster) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Monster.
// Note that you need to call Monster.Unwrap() before calling this method if this Monster
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Monster) Update() *MonsterUpdateOne {
	return NewMonsterClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Monster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Monster) Unwrap() *Monster {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Monster is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Monster) String() string {
	var builder strings.Builder
	builder.WriteString("Monster(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("Name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("HP=")
	builder.WriteString(fmt.Sprintf("%v", m.HP))
	builder.WriteString(", ")
	builder.WriteString("Block=")
	builder.WriteString(fmt.Sprintf("%v", m.Block))
	builder.WriteString(", ")
	builder.WriteString("Power=")
	builder.WriteString(fmt.Sprintf("%v", m.Power))
	builder.WriteString(", ")
	builder.WriteString("Actions=")
	builder.WriteString(fmt.Sprintf("%v", m.Actions))
	builder.WriteString(", ")
	builder.WriteString("Image=")
	builder.WriteString(m.Image)
	builder.WriteByte(')')
	return builder.String()
}

// Monsters is a parsable slice of Monster.
type Monsters []*Monster
