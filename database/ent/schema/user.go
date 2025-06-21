package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("email").NotEmpty().Unique(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.Time("created_at").Default(time.Now().UTC()).Immutable(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "email"),
	}
}
