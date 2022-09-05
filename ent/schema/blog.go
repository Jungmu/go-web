package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {

	return []ent.Field{
		field.Int64("id"),

		field.String("title").
			Comment("title").
			Unique(),

		field.String("sub_title").
			Default("").
			Comment("sub_title"),

		field.String("tags").
			Default("").
			Comment("tags"),

		field.String("content").
			Comment("content"),

		field.Time("update_datetime").
			Comment("update_datetime").
			UpdateDefault(time.Now).
			Default(time.Now),

		field.Time("create_datetime").
			Comment("update_datetime").
			Default(time.Now),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}
