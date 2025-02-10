package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserGreetings holds the schema definition for the UserGreetings entity.
type UserGreetings struct {
	ent.Schema
}

// Fields of the UserGreetings.
func (UserGreetings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Text("greeting"),
	}
}

// Edges of the UserGreetings.
func (UserGreetings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("user_greetings").Field("user_id").Unique(),
	}
}
