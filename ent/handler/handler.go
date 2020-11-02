// Code generated by entc, DO NOT EDIT.

package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/masseelch/go-api-skeleton/ent"
)

// The JobHandler.
type JobHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new JobHandler
func NewJobHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *JobHandler {
	return &JobHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h JobHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *JobHandler) EnableAllEndpoints() *JobHandler {
	h.EnableReadEndpoint()
	return h
}

// Enable the read operation.
func (h *JobHandler) EnableReadEndpoint() *JobHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// The SessionHandler.
type SessionHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new SessionHandler
func NewSessionHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *SessionHandler {
	return &SessionHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *SessionHandler) EnableAllEndpoints() *SessionHandler {
	h.EnableReadEndpoint()
	return h
}

// Enable the read operation.
func (h *SessionHandler) EnableReadEndpoint() *SessionHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// The UserHandler.
type UserHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new UserHandler
func NewUserHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *UserHandler {
	return &UserHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *UserHandler) EnableAllEndpoints() *UserHandler {
	h.EnableReadEndpoint()
	return h
}

// Enable the read operation.
func (h *UserHandler) EnableReadEndpoint() *UserHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}