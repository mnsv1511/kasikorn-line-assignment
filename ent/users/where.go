// Code generated by ent, DO NOT EDIT.

package users

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldName, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldImageURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldName, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldImageURL, v))
}

// HasAccounts applies the HasEdge predicate on the "accounts" edge.
func HasAccounts() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountsWith applies the HasEdge predicate on the "accounts" edge with a given conditions (other predicates).
func HasAccountsWith(preds ...predicate.Accounts) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newAccountsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBanners applies the HasEdge predicate on the "banners" edge.
func HasBanners() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BannersTable, BannersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBannersWith applies the HasEdge predicate on the "banners" edge with a given conditions (other predicates).
func HasBannersWith(preds ...predicate.Banners) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newBannersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDebitCards applies the HasEdge predicate on the "debit_cards" edge.
func HasDebitCards() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DebitCardsTable, DebitCardsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDebitCardsWith applies the HasEdge predicate on the "debit_cards" edge with a given conditions (other predicates).
func HasDebitCardsWith(preds ...predicate.DebitCards) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newDebitCardsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.Transactions) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newTransactionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserGreetings applies the HasEdge predicate on the "user_greetings" edge.
func HasUserGreetings() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserGreetingsTable, UserGreetingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserGreetingsWith applies the HasEdge predicate on the "user_greetings" edge with a given conditions (other predicates).
func HasUserGreetingsWith(preds ...predicate.UserGreetings) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newUserGreetingsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountBalances applies the HasEdge predicate on the "account_balances" edge.
func HasAccountBalances() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountBalancesTable, AccountBalancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountBalancesWith applies the HasEdge predicate on the "account_balances" edge with a given conditions (other predicates).
func HasAccountBalancesWith(preds ...predicate.AccountBalances) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newAccountBalancesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountDetails applies the HasEdge predicate on the "account_details" edge.
func HasAccountDetails() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountDetailsTable, AccountDetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountDetailsWith applies the HasEdge predicate on the "account_details" edge with a given conditions (other predicates).
func HasAccountDetailsWith(preds ...predicate.AccountDetails) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newAccountDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountFlags applies the HasEdge predicate on the "account_flags" edge.
func HasAccountFlags() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountFlagsTable, AccountFlagsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountFlagsWith applies the HasEdge predicate on the "account_flags" edge with a given conditions (other predicates).
func HasAccountFlagsWith(preds ...predicate.AccountFlags) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newAccountFlagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDebitCardDesign applies the HasEdge predicate on the "debit_card_design" edge.
func HasDebitCardDesign() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DebitCardDesignTable, DebitCardDesignColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDebitCardDesignWith applies the HasEdge predicate on the "debit_card_design" edge with a given conditions (other predicates).
func HasDebitCardDesignWith(preds ...predicate.DebitCardDesign) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newDebitCardDesignStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDebitCardDetails applies the HasEdge predicate on the "debit_card_details" edge.
func HasDebitCardDetails() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DebitCardDetailsTable, DebitCardDetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDebitCardDetailsWith applies the HasEdge predicate on the "debit_card_details" edge with a given conditions (other predicates).
func HasDebitCardDetailsWith(preds ...predicate.DebitCardDetails) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newDebitCardDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDebitCardStatus applies the HasEdge predicate on the "debit_card_status" edge.
func HasDebitCardStatus() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DebitCardStatusTable, DebitCardStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDebitCardStatusWith applies the HasEdge predicate on the "debit_card_status" edge with a given conditions (other predicates).
func HasDebitCardStatusWith(preds ...predicate.DebitCardStatus) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newDebitCardStatusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(sql.NotPredicates(p))
}
