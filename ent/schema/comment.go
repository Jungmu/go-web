package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {

	return []ent.Field{
		field.Int64("id"),

		field.String("name").
			Comment("name"),

		field.String("password").
			Comment("password"),

		field.String("content").
			Comment("content"),

		field.Int64("comment_id").
			Comment("comment_id").
			Default(0),

		field.Time("update_datetime").
			Comment("update_datetime").
			UpdateDefault(time.Now).
			Default(time.Now),

		field.Time("create_datetime").
			Comment("update_datetime").
			Default(time.Now),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Blog.Type).
			Ref("comments").
			Unique(),
	}
}
