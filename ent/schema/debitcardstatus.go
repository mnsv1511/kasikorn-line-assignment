package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DebitCardStatus holds the schema definition for the DebitCardStatus entity.
type DebitCardStatus struct {
	ent.Schema
}

// Fields of the DebitCardStatus.
func (DebitCardStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("card_id").Optional(),
		field.Int("user_id").Optional(),
		field.String("status"),
	}
}

// Edges of the DebitCardStatus.
func (DebitCardStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("debit_card_status").Field("user_id").Unique(),
		edge.From("debit_cards", DebitCards.Type).Ref("debit_card_status").Field("card_id").Unique(),
	}
}
