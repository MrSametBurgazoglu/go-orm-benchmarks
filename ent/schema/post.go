package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("title").
			NotEmpty(),
		field.String("content").
			NotEmpty(),
		field.Uint("user_id"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Unique().
			Required().
			Ref("posts").
			Field("user_id"),
		edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
