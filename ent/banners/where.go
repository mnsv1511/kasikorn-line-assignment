// Code generated by ent, DO NOT EDIT.

package banners

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Banners {
	return predicate.Banners(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Banners {
	return predicate.Banners(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Banners {
	return predicate.Banners(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Banners {
	return predicate.Banners(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Banners {
	return predicate.Banners(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Banners {
	return predicate.Banners(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Banners {
	return predicate.Banners(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldUserID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldDescription, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldImage, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Banners {
	return predicate.Banners(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Banners {
	return predicate.Banners(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Banners {
	return predicate.Banners(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Banners {
	return predicate.Banners(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Banners {
	return predicate.Banners(sql.FieldNotNull(FieldUserID))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContainsFold(FieldDescription, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Banners {
	return predicate.Banners(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Banners {
	return predicate.Banners(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Banners {
	return predicate.Banners(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Banners {
	return predicate.Banners(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Banners {
	return predicate.Banners(sql.FieldContainsFold(FieldImage, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Banners {
	return predicate.Banners(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.Users) predicate.Banners {
	return predicate.Banners(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Banners) predicate.Banners {
	return predicate.Banners(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Banners) predicate.Banners {
	return predicate.Banners(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Banners) predicate.Banners {
	return predicate.Banners(sql.NotPredicates(p))
}
