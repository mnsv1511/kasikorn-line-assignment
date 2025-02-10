// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountflags"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accounts"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// AccountFlags is the model entity for the AccountFlags schema.
type AccountFlags struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID int `json:"account_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// FlagType holds the value of the "flag_type" field.
	FlagType string `json:"flag_type,omitempty"`
	// FlagValue holds the value of the "flag_value" field.
	FlagValue string `json:"flag_value,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountFlagsQuery when eager-loading is set.
	Edges        AccountFlagsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountFlagsEdges holds the relations/edges for other nodes in the graph.
type AccountFlagsEdges struct {
	// Users holds the value of the users edge.
	Users *Users `json:"users,omitempty"`
	// Accounts holds the value of the accounts edge.
	Accounts *Accounts `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountFlagsEdges) UsersOrErr() (*Users, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountFlagsEdges) AccountsOrErr() (*Accounts, error) {
	if e.Accounts != nil {
		return e.Accounts, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: accounts.Label}
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountFlags) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountflags.FieldID, accountflags.FieldAccountID, accountflags.FieldUserID:
			values[i] = new(sql.NullInt64)
		case accountflags.FieldFlagType, accountflags.FieldFlagValue:
			values[i] = new(sql.NullString)
		case accountflags.FieldCreatedAt, accountflags.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountFlags fields.
func (af *AccountFlags) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountflags.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			af.ID = int(value.Int64)
		case accountflags.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				af.AccountID = int(value.Int64)
			}
		case accountflags.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				af.UserID = int(value.Int64)
			}
		case accountflags.FieldFlagType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flag_type", values[i])
			} else if value.Valid {
				af.FlagType = value.String
			}
		case accountflags.FieldFlagValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flag_value", values[i])
			} else if value.Valid {
				af.FlagValue = value.String
			}
		case accountflags.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				af.CreatedAt = value.Time
			}
		case accountflags.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				af.UpdatedAt = value.Time
			}
		default:
			af.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccountFlags.
// This includes values selected through modifiers, order, etc.
func (af *AccountFlags) Value(name string) (ent.Value, error) {
	return af.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the AccountFlags entity.
func (af *AccountFlags) QueryUsers() *UsersQuery {
	return NewAccountFlagsClient(af.config).QueryUsers(af)
}

// QueryAccounts queries the "accounts" edge of the AccountFlags entity.
func (af *AccountFlags) QueryAccounts() *AccountsQuery {
	return NewAccountFlagsClient(af.config).QueryAccounts(af)
}

// Update returns a builder for updating this AccountFlags.
// Note that you need to call AccountFlags.Unwrap() before calling this method if this AccountFlags
// was returned from a transaction, and the transaction was committed or rolled back.
func (af *AccountFlags) Update() *AccountFlagsUpdateOne {
	return NewAccountFlagsClient(af.config).UpdateOne(af)
}

// Unwrap unwraps the AccountFlags entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (af *AccountFlags) Unwrap() *AccountFlags {
	_tx, ok := af.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountFlags is not a transactional entity")
	}
	af.config.driver = _tx.drv
	return af
}

// String implements the fmt.Stringer.
func (af *AccountFlags) String() string {
	var builder strings.Builder
	builder.WriteString("AccountFlags(")
	builder.WriteString(fmt.Sprintf("id=%v, ", af.ID))
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", af.AccountID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", af.UserID))
	builder.WriteString(", ")
	builder.WriteString("flag_type=")
	builder.WriteString(af.FlagType)
	builder.WriteString(", ")
	builder.WriteString("flag_value=")
	builder.WriteString(af.FlagValue)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(af.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(af.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AccountFlagsSlice is a parsable slice of AccountFlags.
type AccountFlagsSlice []*AccountFlags
