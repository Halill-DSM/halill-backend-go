package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	regExp := "/^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i"
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("email").Match(regexp.MustCompile(regExp)),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
	}
}
