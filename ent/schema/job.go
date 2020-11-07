package schema

import (
	"fmt"
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
			StructTag(`groups:"job:list,job:read"`),
		field.Time("date").
			Optional().
			StructTag(`groups:"job:list,job:read"`),
		field.Text("task").
			StructTag(`groups:"job:list,job:read" validate:"required"`),
		field.String("state").
			Default(JobStateOpen).
			StructTag(fmt.Sprintf(`groups:"job:list,job:read" validate:"required,oneof=%s %s %s"`, JobStateOpen, JobStateClosed, JobStateBilled)),
		field.Text("report").
			Optional().
			StructTag(`groups:"job:list,job:read"`),
		field.Text("rest").
			Optional().
			StructTag(`groups:"job:list,job:read"`),
		field.Text("note").
			Optional().
			StructTag(`groups:"job:list,job:read"`),
		field.String("customerName").
			Optional().
			StructTag(`groups:"job:list,job:read"`),
		field.Bool("riskAssessmentRequired").
			Default(false).
			StructTag(`groups:"job:list,job:read"`),
		field.Bool("maintenanceRequired").
			Default(false).
			StructTag(`groups:"job:list,job:read"`),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			StructTag(`json:"users,omitempty" groups:"job:list,job:read"`),
	}
}

// Annotations of the Job.
func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"edges" groups:"job:list,job:read"`,
		},
		HandlerAnnotation{
			ReadGroups: []string{"job:list", "user:list"},
			CreateGroups: []string{"job:read", "user:list"},

			ReadEager: []string{"users"},
			ListEager: []string{"users"},
		},
	}
}
