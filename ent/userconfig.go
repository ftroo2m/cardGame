// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cardGame/ent/userconfig"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserConfig is the model entity for the UserConfig schema.
type UserConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PlayerID holds the value of the "playerID" field.
	PlayerID string `json:"playerID,omitempty"`
	// Cards holds the value of the "cards" field.
	Cards []string `json:"cards,omitempty"`
	// Ladder holds the value of the "ladder" field.
	Ladder string `json:"ladder,omitempty"`
	// PlayerHP holds the value of the "playerHP" field.
	PlayerHP string `json:"playerHP,omitempty"`
	// PlayerEnergy holds the value of the "playerEnergy" field.
	PlayerEnergy int `json:"playerEnergy,omitempty"`
	// Image holds the value of the "image" field.
	Image        string `json:"image,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userconfig.FieldCards:
			values[i] = new([]byte)
		case userconfig.FieldID, userconfig.FieldPlayerEnergy:
			values[i] = new(sql.NullInt64)
		case userconfig.FieldPlayerID, userconfig.FieldLadder, userconfig.FieldPlayerHP, userconfig.FieldImage:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserConfig fields.
func (uc *UserConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int(value.Int64)
		case userconfig.FieldPlayerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field playerID", values[i])
			} else if value.Valid {
				uc.PlayerID = value.String
			}
		case userconfig.FieldCards:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cards", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &uc.Cards); err != nil {
					return fmt.Errorf("unmarshal field cards: %w", err)
				}
			}
		case userconfig.FieldLadder:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ladder", values[i])
			} else if value.Valid {
				uc.Ladder = value.String
			}
		case userconfig.FieldPlayerHP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field playerHP", values[i])
			} else if value.Valid {
				uc.PlayerHP = value.String
			}
		case userconfig.FieldPlayerEnergy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field playerEnergy", values[i])
			} else if value.Valid {
				uc.PlayerEnergy = int(value.Int64)
			}
		case userconfig.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				uc.Image = value.String
			}
		default:
			uc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserConfig.
// This includes values selected through modifiers, order, etc.
func (uc *UserConfig) Value(name string) (ent.Value, error) {
	return uc.selectValues.Get(name)
}

// Update returns a builder for updating this UserConfig.
// Note that you need to call UserConfig.Unwrap() before calling this method if this UserConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserConfig) Update() *UserConfigUpdateOne {
	return NewUserConfigClient(uc.config).UpdateOne(uc)
}

// Unwrap unwraps the UserConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserConfig) Unwrap() *UserConfig {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserConfig is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserConfig) String() string {
	var builder strings.Builder
	builder.WriteString("UserConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("playerID=")
	builder.WriteString(uc.PlayerID)
	builder.WriteString(", ")
	builder.WriteString("cards=")
	builder.WriteString(fmt.Sprintf("%v", uc.Cards))
	builder.WriteString(", ")
	builder.WriteString("ladder=")
	builder.WriteString(uc.Ladder)
	builder.WriteString(", ")
	builder.WriteString("playerHP=")
	builder.WriteString(uc.PlayerHP)
	builder.WriteString(", ")
	builder.WriteString("playerEnergy=")
	builder.WriteString(fmt.Sprintf("%v", uc.PlayerEnergy))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(uc.Image)
	builder.WriteByte(')')
	return builder.String()
}

// UserConfigs is a parsable slice of UserConfig.
type UserConfigs []*UserConfig
