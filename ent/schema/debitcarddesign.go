package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DebitCardDesign holds the schema definition for the DebitCardDesign entity.
type DebitCardDesign struct {
	ent.Schema
}

// Fields of the DebitCardDesign.
func (DebitCardDesign) Fields() []ent.Field {
	return []ent.Field{
		field.Int("card_id").Optional(),
		field.Int("user_id").Optional(),
		field.String("color"),
		field.String("border_color"),
	}
}

// Edges of the DebitCardDesign.
func (DebitCardDesign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("debit_card_design").Field("user_id").Unique(),
		edge.From("debit_cards", DebitCards.Type).Ref("debit_card_design").Field("card_id").Unique(),
	}
}
