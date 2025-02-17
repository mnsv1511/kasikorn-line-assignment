// Code generated by ent, DO NOT EDIT.

package transactions

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transactions {
	return predicate.Transactions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transactions {
	return predicate.Transactions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transactions {
	return predicate.Transactions(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldUserID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldName, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldImage, v))
}

// IsBank applies equality check predicate on the "isBank" field. It's identical to IsBankEQ.
func IsBank(v bool) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldIsBank, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Transactions {
	return predicate.Transactions(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Transactions {
	return predicate.Transactions(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Transactions {
	return predicate.Transactions(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Transactions {
	return predicate.Transactions(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Transactions {
	return predicate.Transactions(sql.FieldNotNull(FieldUserID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Transactions {
	return predicate.Transactions(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Transactions {
	return predicate.Transactions(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldContainsFold(FieldName, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Transactions {
	return predicate.Transactions(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Transactions {
	return predicate.Transactions(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Transactions {
	return predicate.Transactions(sql.FieldContainsFold(FieldImage, v))
}

// IsBankEQ applies the EQ predicate on the "isBank" field.
func IsBankEQ(v bool) predicate.Transactions {
	return predicate.Transactions(sql.FieldEQ(FieldIsBank, v))
}

// IsBankNEQ applies the NEQ predicate on the "isBank" field.
func IsBankNEQ(v bool) predicate.Transactions {
	return predicate.Transactions(sql.FieldNEQ(FieldIsBank, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.Users) predicate.Transactions {
	return predicate.Transactions(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transactions) predicate.Transactions {
	return predicate.Transactions(sql.NotPredicates(p))
}
