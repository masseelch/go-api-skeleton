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
		field.Int("id"),
		field.Time("date").
			Optional(),
		field.String("task").
			Optional(),
		field.String("state").
			Default(JobStateOpen),
		field.Text("report").
			Optional(),
		field.Text("rest").
			Optional(),
		field.Text("note").
			Optional(),
		field.String("customerName").
			Optional(),
		field.Bool("riskAssessmentRequired").
			Default(false),
		field.Bool("maintenanceRequired").
			Default(false),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			StructTag(`json:"users,omitempty"`),
	}
}

// Annotations of the Job.
func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"edges"`,
		},
		HandlerAnnotation{
			ReadEager: []string{"users"},
		},
	}
}
