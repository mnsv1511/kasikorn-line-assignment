package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountDetails holds the schema definition for the AccountDetails entity.
type AccountDetails struct {
	ent.Schema
}

// Fields of the AccountDetails.
func (AccountDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Optional(),
		field.Int("user_id").Optional(),
		field.String("color"),
		field.Bool("is_main_account"),
		field.Int("progress"),
		field.String("name"),
	}
}

// Edges of the AccountDetails.
func (AccountDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("account_details").Field("user_id").Unique(),
		edge.From("accounts", Accounts.Type).Ref("account_details").Field("account_id").Unique(),
	}
}
