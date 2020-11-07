package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`groups:"user:list"`),
		field.String("email").
			Unique().
			Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")).
			StructTag(`groups:"user:list"`),
		field.String("password").
			Sensitive(),
		field.Bool("enabled").
			Default(false).
			StructTag(`groups:"user:list"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", Session.Type).
			StructTag(`json:"-"`),
		edge.From("jobs", Job.Type).
			Ref("users").
			StructTag(`json:"jobs,omitempty" groups:"user:read"`).
			Annotations(FieldAnnotation{Create: false}),
		edge.From("group", Group.Type).
			Ref("users").Unique(),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"edges" groups:"user:read"`,
		},
	}
}
