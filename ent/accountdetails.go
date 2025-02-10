// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountdetails"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/accounts"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// AccountDetails is the model entity for the AccountDetails schema.
type AccountDetails struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID int `json:"account_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// IsMainAccount holds the value of the "is_main_account" field.
	IsMainAccount bool `json:"is_main_account,omitempty"`
	// Progress holds the value of the "progress" field.
	Progress int `json:"progress,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountDetailsQuery when eager-loading is set.
	Edges        AccountDetailsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountDetailsEdges holds the relations/edges for other nodes in the graph.
type AccountDetailsEdges struct {
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
func (e AccountDetailsEdges) UsersOrErr() (*Users, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountDetailsEdges) AccountsOrErr() (*Accounts, error) {
	if e.Accounts != nil {
		return e.Accounts, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: accounts.Label}
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountDetails) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountdetails.FieldIsMainAccount:
			values[i] = new(sql.NullBool)
		case accountdetails.FieldID, accountdetails.FieldAccountID, accountdetails.FieldUserID, accountdetails.FieldProgress:
			values[i] = new(sql.NullInt64)
		case accountdetails.FieldColor, accountdetails.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountDetails fields.
func (ad *AccountDetails) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountdetails.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ad.ID = int(value.Int64)
		case accountdetails.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				ad.AccountID = int(value.Int64)
			}
		case accountdetails.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ad.UserID = int(value.Int64)
			}
		case accountdetails.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				ad.Color = value.String
			}
		case accountdetails.FieldIsMainAccount:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_main_account", values[i])
			} else if value.Valid {
				ad.IsMainAccount = value.Bool
			}
		case accountdetails.FieldProgress:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field progress", values[i])
			} else if value.Valid {
				ad.Progress = int(value.Int64)
			}
		case accountdetails.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ad.Name = value.String
			}
		default:
			ad.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccountDetails.
// This includes values selected through modifiers, order, etc.
func (ad *AccountDetails) Value(name string) (ent.Value, error) {
	return ad.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the AccountDetails entity.
func (ad *AccountDetails) QueryUsers() *UsersQuery {
	return NewAccountDetailsClient(ad.config).QueryUsers(ad)
}

// QueryAccounts queries the "accounts" edge of the AccountDetails entity.
func (ad *AccountDetails) QueryAccounts() *AccountsQuery {
	return NewAccountDetailsClient(ad.config).QueryAccounts(ad)
}

// Update returns a builder for updating this AccountDetails.
// Note that you need to call AccountDetails.Unwrap() before calling this method if this AccountDetails
// was returned from a transaction, and the transaction was committed or rolled back.
func (ad *AccountDetails) Update() *AccountDetailsUpdateOne {
	return NewAccountDetailsClient(ad.config).UpdateOne(ad)
}

// Unwrap unwraps the AccountDetails entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ad *AccountDetails) Unwrap() *AccountDetails {
	_tx, ok := ad.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountDetails is not a transactional entity")
	}
	ad.config.driver = _tx.drv
	return ad
}

// String implements the fmt.Stringer.
func (ad *AccountDetails) String() string {
	var builder strings.Builder
	builder.WriteString("AccountDetails(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ad.ID))
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.AccountID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.UserID))
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(ad.Color)
	builder.WriteString(", ")
	builder.WriteString("is_main_account=")
	builder.WriteString(fmt.Sprintf("%v", ad.IsMainAccount))
	builder.WriteString(", ")
	builder.WriteString("progress=")
	builder.WriteString(fmt.Sprintf("%v", ad.Progress))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ad.Name)
	builder.WriteByte(')')
	return builder.String()
}

// AccountDetailsSlice is a parsable slice of AccountDetails.
type AccountDetailsSlice []*AccountDetails
