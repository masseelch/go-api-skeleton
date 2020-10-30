package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

const (
	JobStateOpen   = "open"
	JobStateClosed = "closed"
	JobStateBilled = "billed"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`groups:"job:list"`),
		field.Time("date").
			Optional().
			StructTag(`groups:"job:list"`),
		field.String("task").
			Optional().
			StructTag(`groups:"job:list"`),
		field.String("state").
			Default(JobStateOpen).
			StructTag(`groups:"job:list"`),
		field.Text("report").
			Optional().
			StructTag(`groups:"job:list"`),
		field.Text("rest").
			Optional().
			StructTag(`groups:"job:list"`),
		field.Text("note").
			Optional().
			StructTag(`groups:"job:list"`),
		field.String("customerName").
			Optional().
			StructTag(`groups:"job:list"`),
		field.Bool("riskAssessmentRequired").
			Default(false).
			StructTag(`groups:"job:list"`),
		field.Bool("maintenanceRequired").
			Default(false).
			StructTag(`groups:"job:list"`),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			StructTag(`groups:"job:list"`),
	}
}

// Annotations of the Job.
func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"edges" groups:"job:list"`,
		},
	}
}
