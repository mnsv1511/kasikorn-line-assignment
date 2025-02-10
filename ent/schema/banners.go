package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Banners holds the schema definition for the Banners entity.
type Banners struct {
	ent.Schema
}

// Fields of the Banners.
func (Banners) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("title"),
		field.Text("description"),
		field.String("image"),
	}
}

// Edges of the Banners.
func (Banners) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("banners").Field("user_id").Unique(),
	}
}
