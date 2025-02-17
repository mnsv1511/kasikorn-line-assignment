// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcarddesign"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/debitcards"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

// DebitCardDesign is the model entity for the DebitCardDesign schema.
type DebitCardDesign struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CardID holds the value of the "card_id" field.
	CardID int `json:"card_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// BorderColor holds the value of the "border_color" field.
	BorderColor string `json:"border_color,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DebitCardDesignQuery when eager-loading is set.
	Edges        DebitCardDesignEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DebitCardDesignEdges holds the relations/edges for other nodes in the graph.
type DebitCardDesignEdges struct {
	// Users holds the value of the users edge.
	Users *Users `json:"users,omitempty"`
	// DebitCards holds the value of the debit_cards edge.
	DebitCards *DebitCards `json:"debit_cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DebitCardDesignEdges) UsersOrErr() (*Users, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// DebitCardsOrErr returns the DebitCards value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DebitCardDesignEdges) DebitCardsOrErr() (*DebitCards, error) {
	if e.DebitCards != nil {
		return e.DebitCards, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: debitcards.Label}
	}
	return nil, &NotLoadedError{edge: "debit_cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DebitCardDesign) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case debitcarddesign.FieldID, debitcarddesign.FieldCardID, debitcarddesign.FieldUserID:
			values[i] = new(sql.NullInt64)
		case debitcarddesign.FieldColor, debitcarddesign.FieldBorderColor:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DebitCardDesign fields.
func (dcd *DebitCardDesign) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case debitcarddesign.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dcd.ID = int(value.Int64)
		case debitcarddesign.FieldCardID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card_id", values[i])
			} else if value.Valid {
				dcd.CardID = int(value.Int64)
			}
		case debitcarddesign.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				dcd.UserID = int(value.Int64)
			}
		case debitcarddesign.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				dcd.Color = value.String
			}
		case debitcarddesign.FieldBorderColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field border_color", values[i])
			} else if value.Valid {
				dcd.BorderColor = value.String
			}
		default:
			dcd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DebitCardDesign.
// This includes values selected through modifiers, order, etc.
func (dcd *DebitCardDesign) Value(name string) (ent.Value, error) {
	return dcd.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the DebitCardDesign entity.
func (dcd *DebitCardDesign) QueryUsers() *UsersQuery {
	return NewDebitCardDesignClient(dcd.config).QueryUsers(dcd)
}

// QueryDebitCards queries the "debit_cards" edge of the DebitCardDesign entity.
func (dcd *DebitCardDesign) QueryDebitCards() *DebitCardsQuery {
	return NewDebitCardDesignClient(dcd.config).QueryDebitCards(dcd)
}

// Update returns a builder for updating this DebitCardDesign.
// Note that you need to call DebitCardDesign.Unwrap() before calling this method if this DebitCardDesign
// was returned from a transaction, and the transaction was committed or rolled back.
func (dcd *DebitCardDesign) Update() *DebitCardDesignUpdateOne {
	return NewDebitCardDesignClient(dcd.config).UpdateOne(dcd)
}

// Unwrap unwraps the DebitCardDesign entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dcd *DebitCardDesign) Unwrap() *DebitCardDesign {
	_tx, ok := dcd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DebitCardDesign is not a transactional entity")
	}
	dcd.config.driver = _tx.drv
	return dcd
}

// String implements the fmt.Stringer.
func (dcd *DebitCardDesign) String() string {
	var builder strings.Builder
	builder.WriteString("DebitCardDesign(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dcd.ID))
	builder.WriteString("card_id=")
	builder.WriteString(fmt.Sprintf("%v", dcd.CardID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", dcd.UserID))
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(dcd.Color)
	builder.WriteString(", ")
	builder.WriteString("border_color=")
	builder.WriteString(dcd.BorderColor)
	builder.WriteByte(')')
	return builder.String()
}

// DebitCardDesigns is a parsable slice of DebitCardDesign.
type DebitCardDesigns []*DebitCardDesign
