package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transactions holds the schema definition for the Transactions entity.
type Transactions struct {
	ent.Schema
}

// Fields of the Transactions.
func (Transactions) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("name"),
		field.String("image"),
		field.Bool("isBank"),
	}
}

// Edges of the Transactions.
func (Transactions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("transactions").Field("user_id").Unique(),
	}
}
