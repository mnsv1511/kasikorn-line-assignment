package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountBalances holds the schema definition for the AccountBalances entity.
type AccountBalances struct {
	ent.Schema
}

// Fields of the AccountBalances.
func (AccountBalances) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Optional(),
		field.Int("user_id").Optional(),
		field.Float("amount"),
	}
}

// Edges of the AccountBalances.
func (AccountBalances) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("account_balances").Field("user_id").Unique(),
		edge.From("accounts", Accounts.Type).Ref("account_balances").Field("account_id").Unique(),
	}
}
