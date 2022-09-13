package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BlogLog holds the schema definition for the BlogLog entity.
type BlogLog struct {
	ent.Schema
}

// Fields of the BlogLog.
func (BlogLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.Int64("blog_id").
			Comment("blog_id"),

		field.String("url").
			Comment("request url"),

		field.String("reason").
			Comment("reason"),

		field.String("detail").
			Comment("detail"),

		field.String("client_ip").
			Comment("client_ip"),

		field.Time("create_datetime").
			Comment("update_datetime").
			Default(time.Now),
	}
}

// Indexes of the BlogLog.
func (BlogLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("blog_id"),
	}
}

// Edges of the BlogLog.
func (BlogLog) Edges() []ent.Edge {
	return nil
}
