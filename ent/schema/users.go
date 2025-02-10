package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("image_url"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Accounts.Type),
		edge.To("banners", Banners.Type),
		edge.To("debit_cards", DebitCards.Type),
		edge.To("transactions", Transactions.Type),
		edge.To("user_greetings", UserGreetings.Type),
		edge.To("account_balances", AccountBalances.Type),
		edge.To("account_details", AccountDetails.Type),
		edge.To("account_flags", AccountFlags.Type),
		edge.To("debit_card_design", DebitCardDesign.Type),
		edge.To("debit_card_details", DebitCardDetails.Type),
		edge.To("debit_card_status", DebitCardStatus.Type),
	}
}
