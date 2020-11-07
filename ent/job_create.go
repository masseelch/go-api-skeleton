// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/masseelch/go-api-skeleton/ent/job"
	"github.com/masseelch/go-api-skeleton/ent/user"
)

// JobCreate is the builder for creating a Job entity.
type JobCreate struct {
	config
	mutation *JobMutation
	hooks    []Hook
}

// SetDate sets the date field.
func (jc *JobCreate) SetDate(t time.Time) *JobCreate {
	jc.mutation.SetDate(t)
	return jc
}

// SetNillableDate sets the date field if the given value is not nil.
func (jc *JobCreate) SetNillableDate(t *time.Time) *JobCreate {
	if t != nil {
		jc.SetDate(*t)
	}
	return jc
}

// SetTask sets the task field.
func (jc *JobCreate) SetTask(s string) *JobCreate {
	jc.mutation.SetTask(s)
	return jc
}

// SetState sets the state field.
func (jc *JobCreate) SetState(s string) *JobCreate {
	jc.mutation.SetState(s)
	return jc
}

// SetNillableState sets the state field if the given value is not nil.
func (jc *JobCreate) SetNillableState(s *string) *JobCreate {
	if s != nil {
		jc.SetState(*s)
	}
	return jc
}

// SetReport sets the report field.
func (jc *JobCreate) SetReport(s string) *JobCreate {
	jc.mutation.SetReport(s)
	return jc
}

// SetNillableReport sets the report field if the given value is not nil.
func (jc *JobCreate) SetNillableReport(s *string) *JobCreate {
	if s != nil {
		jc.SetReport(*s)
	}
	return jc
}

// SetRest sets the rest field.
func (jc *JobCreate) SetRest(s string) *JobCreate {
	jc.mutation.SetRest(s)
	return jc
}

// SetNillableRest sets the rest field if the given value is not nil.
func (jc *JobCreate) SetNillableRest(s *string) *JobCreate {
	if s != nil {
		jc.SetRest(*s)
	}
	return jc
}

// SetNote sets the note field.
func (jc *JobCreate) SetNote(s string) *JobCreate {
	jc.mutation.SetNote(s)
	return jc
}

// SetNillableNote sets the note field if the given value is not nil.
func (jc *JobCreate) SetNillableNote(s *string) *JobCreate {
	if s != nil {
		jc.SetNote(*s)
	}
	return jc
}

// SetCustomerName sets the customerName field.
func (jc *JobCreate) SetCustomerName(s string) *JobCreate {
	jc.mutation.SetCustomerName(s)
	return jc
}

// SetNillableCustomerName sets the customerName field if the given value is not nil.
func (jc *JobCreate) SetNillableCustomerName(s *string) *JobCreate {
	if s != nil {
		jc.SetCustomerName(*s)
	}
	return jc
}

// SetRiskAssessmentRequired sets the riskAssessmentRequired field.
func (jc *JobCreate) SetRiskAssessmentRequired(b bool) *JobCreate {
	jc.mutation.SetRiskAssessmentRequired(b)
	return jc
}

// SetNillableRiskAssessmentRequired sets the riskAssessmentRequired field if the given value is not nil.
func (jc *JobCreate) SetNillableRiskAssessmentRequired(b *bool) *JobCreate {
	if b != nil {
		jc.SetRiskAssessmentRequired(*b)
	}
	return jc
}

// SetMaintenanceRequired sets the maintenanceRequired field.
func (jc *JobCreate) SetMaintenanceRequired(b bool) *JobCreate {
	jc.mutation.SetMaintenanceRequired(b)
	return jc
}

// SetNillableMaintenanceRequired sets the maintenanceRequired field if the given value is not nil.
func (jc *JobCreate) SetNillableMaintenanceRequired(b *bool) *JobCreate {
	if b != nil {
		jc.SetMaintenanceRequired(*b)
	}
	return jc
}

// SetID sets the id field.
func (jc *JobCreate) SetID(i int) *JobCreate {
	jc.mutation.SetID(i)
	return jc
}

// AddUserIDs adds the users edge to User by ids.
func (jc *JobCreate) AddUserIDs(ids ...int) *JobCreate {
	jc.mutation.AddUserIDs(ids...)
	return jc
}

// AddUsers adds the users edges to User.
func (jc *JobCreate) AddUsers(u ...*User) *JobCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jc.AddUserIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (jc *JobCreate) Mutation() *JobMutation {
	return jc.mutation
}

// Save creates the Job in the database.
func (jc *JobCreate) Save(ctx context.Context) (*Job, error) {
	var (
		err  error
		node *Job
	)
	jc.defaults()
	if len(jc.hooks) == 0 {
		if err = jc.check(); err != nil {
			return nil, err
		}
		node, err = jc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jc.check(); err != nil {
				return nil, err
			}
			jc.mutation = mutation
			node, err = jc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jc.hooks) - 1; i >= 0; i-- {
			mut = jc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobCreate) SaveX(ctx context.Context) *Job {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (jc *JobCreate) defaults() {
	if _, ok := jc.mutation.State(); !ok {
		v := job.DefaultState
		jc.mutation.SetState(v)
	}
	if _, ok := jc.mutation.RiskAssessmentRequired(); !ok {
		v := job.DefaultRiskAssessmentRequired
		jc.mutation.SetRiskAssessmentRequired(v)
	}
	if _, ok := jc.mutation.MaintenanceRequired(); !ok {
		v := job.DefaultMaintenanceRequired
		jc.mutation.SetMaintenanceRequired(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JobCreate) check() error {
	if _, ok := jc.mutation.Task(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required field \"task\"")}
	}
	if _, ok := jc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required field \"state\"")}
	}
	if _, ok := jc.mutation.RiskAssessmentRequired(); !ok {
		return &ValidationError{Name: "riskAssessmentRequired", err: errors.New("ent: missing required field \"riskAssessmentRequired\"")}
	}
	if _, ok := jc.mutation.MaintenanceRequired(); !ok {
		return &ValidationError{Name: "maintenanceRequired", err: errors.New("ent: missing required field \"maintenanceRequired\"")}
	}
	return nil
}

func (jc *JobCreate) sqlSave(ctx context.Context) (*Job, error) {
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (jc *JobCreate) createSpec() (*Job, *sqlgraph.CreateSpec) {
	var (
		_node = &Job{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: job.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		}
	)
	if id, ok := jc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := jc.mutation.Task(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldTask,
		})
		_node.Task = value
	}
	if value, ok := jc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldState,
		})
		_node.State = value
	}
	if value, ok := jc.mutation.Report(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldReport,
		})
		_node.Report = value
	}
	if value, ok := jc.mutation.Rest(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldRest,
		})
		_node.Rest = value
	}
	if value, ok := jc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldNote,
		})
		_node.Note = value
	}
	if value, ok := jc.mutation.CustomerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldCustomerName,
		})
		_node.CustomerName = value
	}
	if value, ok := jc.mutation.RiskAssessmentRequired(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: job.FieldRiskAssessmentRequired,
		})
		_node.RiskAssessmentRequired = value
	}
	if value, ok := jc.mutation.MaintenanceRequired(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: job.FieldMaintenanceRequired,
		})
		_node.MaintenanceRequired = value
	}
	if nodes := jc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   job.UsersTable,
			Columns: job.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobCreateBulk is the builder for creating a bulk of Job entities.
type JobCreateBulk struct {
	config
	builders []*JobCreate
}

// Save creates the Job entities in the database.
func (jcb *JobCreateBulk) Save(ctx context.Context) ([]*Job, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Job, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (jcb *JobCreateBulk) SaveX(ctx context.Context) []*Job {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
