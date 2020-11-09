package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`groups:"tag:list"`),
		field.String("title").
			StructTag(`groups:"tag:list"`),
		field.Text("description").
			StructTag(`groups:"tag:list"`),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
