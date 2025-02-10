package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Accounts holds the schema definition for the Accounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the Accounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("type"),
		field.String("currency"),
		field.String("account_number"),
		field.String("issuer"),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account_balances", AccountBalances.Type),
		edge.To("account_details", AccountDetails.Type),
		edge.To("account_flags", AccountFlags.Type),
		edge.From("users", Users.Type).Ref("accounts").Field("user_id").Unique(),
	}
}
