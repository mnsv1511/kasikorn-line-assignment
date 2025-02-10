package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountFlags holds the schema definition for the AccountFlags entity.
type AccountFlags struct {
	ent.Schema
}

// Fields of the AccountFlags.
func (AccountFlags) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").Optional(),
		field.Int("user_id").Optional(),
		field.String("flag_type"),
		field.String("flag_value"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the AccountFlags.
func (AccountFlags) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("account_flags").Field("user_id").Unique(),
		edge.From("accounts", Accounts.Type).Ref("account_flags").Field("account_id").Unique(),
	}
}
