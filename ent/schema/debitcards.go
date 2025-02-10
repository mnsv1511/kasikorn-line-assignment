package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DebitCards holds the schema definition for the DebitCards entity.
type DebitCards struct {
	ent.Schema
}

// Fields of the DebitCards.
func (DebitCards) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("name"),
	}
}

// Edges of the DebitCards.
func (DebitCards) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("debit_card_design", DebitCardDesign.Type),
		edge.To("debit_card_details", DebitCardDetails.Type),
		edge.To("debit_card_status", DebitCardStatus.Type),
		edge.From("users", Users.Type).Ref("debit_cards").Field("user_id").Unique(),
	}
}
