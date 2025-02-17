// Code generated by ent, DO NOT EDIT.

package debitcarddesign

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLTE(FieldID, id))
}

// CardID applies equality check predicate on the "card_id" field. It's identical to CardIDEQ.
func CardID(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldCardID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldUserID, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldColor, v))
}

// BorderColor applies equality check predicate on the "border_color" field. It's identical to BorderColorEQ.
func BorderColor(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldBorderColor, v))
}

// CardIDEQ applies the EQ predicate on the "card_id" field.
func CardIDEQ(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldCardID, v))
}

// CardIDNEQ applies the NEQ predicate on the "card_id" field.
func CardIDNEQ(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNEQ(FieldCardID, v))
}

// CardIDIn applies the In predicate on the "card_id" field.
func CardIDIn(vs ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIn(FieldCardID, vs...))
}

// CardIDNotIn applies the NotIn predicate on the "card_id" field.
func CardIDNotIn(vs ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotIn(FieldCardID, vs...))
}

// CardIDIsNil applies the IsNil predicate on the "card_id" field.
func CardIDIsNil() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIsNull(FieldCardID))
}

// CardIDNotNil applies the NotNil predicate on the "card_id" field.
func CardIDNotNil() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotNull(FieldCardID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotNull(FieldUserID))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldHasSuffix(FieldColor, v))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldContainsFold(FieldColor, v))
}

// BorderColorEQ applies the EQ predicate on the "border_color" field.
func BorderColorEQ(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEQ(FieldBorderColor, v))
}

// BorderColorNEQ applies the NEQ predicate on the "border_color" field.
func BorderColorNEQ(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNEQ(FieldBorderColor, v))
}

// BorderColorIn applies the In predicate on the "border_color" field.
func BorderColorIn(vs ...string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldIn(FieldBorderColor, vs...))
}

// BorderColorNotIn applies the NotIn predicate on the "border_color" field.
func BorderColorNotIn(vs ...string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldNotIn(FieldBorderColor, vs...))
}

// BorderColorGT applies the GT predicate on the "border_color" field.
func BorderColorGT(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGT(FieldBorderColor, v))
}

// BorderColorGTE applies the GTE predicate on the "border_color" field.
func BorderColorGTE(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldGTE(FieldBorderColor, v))
}

// BorderColorLT applies the LT predicate on the "border_color" field.
func BorderColorLT(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLT(FieldBorderColor, v))
}

// BorderColorLTE applies the LTE predicate on the "border_color" field.
func BorderColorLTE(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldLTE(FieldBorderColor, v))
}

// BorderColorContains applies the Contains predicate on the "border_color" field.
func BorderColorContains(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldContains(FieldBorderColor, v))
}

// BorderColorHasPrefix applies the HasPrefix predicate on the "border_color" field.
func BorderColorHasPrefix(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldHasPrefix(FieldBorderColor, v))
}

// BorderColorHasSuffix applies the HasSuffix predicate on the "border_color" field.
func BorderColorHasSuffix(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldHasSuffix(FieldBorderColor, v))
}

// BorderColorEqualFold applies the EqualFold predicate on the "border_color" field.
func BorderColorEqualFold(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldEqualFold(FieldBorderColor, v))
}

// BorderColorContainsFold applies the ContainsFold predicate on the "border_color" field.
func BorderColorContainsFold(v string) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.FieldContainsFold(FieldBorderColor, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.Users) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDebitCards applies the HasEdge predicate on the "debit_cards" edge.
func HasDebitCards() predicate.DebitCardDesign {
	return predicate.DebitCardDesign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DebitCardsTable, DebitCardsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDebitCardsWith applies the HasEdge predicate on the "debit_cards" edge with a given conditions (other predicates).
func HasDebitCardsWith(preds ...predicate.DebitCards) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(func(s *sql.Selector) {
		step := newDebitCardsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DebitCardDesign) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DebitCardDesign) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DebitCardDesign) predicate.DebitCardDesign {
	return predicate.DebitCardDesign(sql.NotPredicates(p))
}
