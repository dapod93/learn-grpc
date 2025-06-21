package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
