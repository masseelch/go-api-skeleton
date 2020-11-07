// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/masseelch/go-api-skeleton/ent/migrate"
	go_token "github.com/masseelch/go-token"

	"github.com/masseelch/go-api-skeleton/ent/group"
	"github.com/masseelch/go-api-skeleton/ent/job"
	"github.com/masseelch/go-api-skeleton/ent/session"
	"github.com/masseelch/go-api-skeleton/ent/user"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Group = NewGroupClient(c.config)
	c.Job = NewJobClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Group:   NewGroupClient(cfg),
		Job:     NewJobClient(cfg),
		Session: NewSessionClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:  cfg,
		Group:   NewGroupClient(cfg),
		Job:     NewJobClient(cfg),
		Session: NewSessionClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Group.Use(hooks...)
	c.Job.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.UsersTable, group.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `job.Hooks(f(g(h())))`.
func (c *JobClient) Use(hooks ...Hook) {
	c.hooks.Job = append(c.hooks.Job, hooks...)
}

// Create returns a create builder for Job.
func (c *JobClient) Create() *JobCreate {
	mutation := newJobMutation(c.config, OpCreate)
	return &JobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Job entities.
func (c *JobClient) CreateBulk(builders ...*JobCreate) *JobCreateBulk {
	return &JobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	mutation := newJobMutation(c.config, OpUpdate)
	return &JobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJob(j))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id int) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJobID(id))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	mutation := newJobMutation(c.config, OpDelete)
	return &JobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *JobClient) DeleteOneID(id int) *JobDeleteOne {
	builder := c.Delete().Where(job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JobDeleteOne{builder}
}

// Query returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{config: c.config}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id int) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id int) *Job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Job.
func (c *JobClient) QueryUsers(j *Job) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, job.UsersTable, job.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *JobClient) Hooks() []Hook {
	return c.hooks.Job
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id go_token.Token) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id go_token.Token) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id go_token.Token) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id go_token.Token) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Session.
func (c *SessionClient) QueryUser(s *Session) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, session.UserTable, session.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySessions queries the sessions edge of a User.
func (c *UserClient) QuerySessions(u *User) *SessionQuery {
	query := &SessionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SessionsTable, user.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryJobs queries the jobs edge of a User.
func (c *UserClient) QueryJobs(u *User) *JobQuery {
	query := &JobQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.JobsTable, user.JobsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a User.
func (c *UserClient) QueryGroup(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.GroupTable, user.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
