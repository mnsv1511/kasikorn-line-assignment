package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DebitCardDetails holds the schema definition for the DebitCardDetails entity.
type DebitCardDetails struct {
	ent.Schema
}

// Fields of the DebitCardDetails.
func (DebitCardDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int("card_id").Optional(),
		field.Int("user_id").Optional(),
		field.String("issuer"),
		field.String("number"),
	}
}

// Edges of the DebitCardDetails.
func (DebitCardDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("debit_card_details").Field("user_id").Unique(),
		edge.From("debit_cards", DebitCards.Type).Ref("debit_card_details").Field("card_id").Unique(),
	}
}
