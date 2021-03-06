// Code generated by entc, DO NOT EDIT.

package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/masseelch/go-api-skeleton/ent"
)

// The AccountHandler.
type AccountHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new AccountHandler
func NewAccountHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *AccountHandler {
	return &AccountHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *AccountHandler) EnableAllEndpoints() *AccountHandler {
	h.EnableCreateEndpoint()
	h.EnableReadEndpoint()
	h.EnableListEndpoint()
	return h
}

// Enable the create operation.
func (h *AccountHandler) EnableCreateEndpoint() *AccountHandler {
	h.r.Post("/", h.Create)
	return h
}

// Enable the read operation.
func (h *AccountHandler) EnableReadEndpoint() *AccountHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// Enable the list operation.
func (h *AccountHandler) EnableListEndpoint() *AccountHandler {
	h.r.Get("/", h.List)
	return h
}

// The TagHandler.
type TagHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new TagHandler
func NewTagHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *TagHandler {
	return &TagHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h TagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *TagHandler) EnableAllEndpoints() *TagHandler {
	h.EnableCreateEndpoint()
	h.EnableReadEndpoint()
	h.EnableListEndpoint()
	return h
}

// Enable the create operation.
func (h *TagHandler) EnableCreateEndpoint() *TagHandler {
	h.r.Post("/", h.Create)
	return h
}

// Enable the read operation.
func (h *TagHandler) EnableReadEndpoint() *TagHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// Enable the list operation.
func (h *TagHandler) EnableListEndpoint() *TagHandler {
	h.r.Get("/", h.List)
	return h
}

// The TransactionHandler.
type TransactionHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// Create a new TransactionHandler
func NewTransactionHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *TransactionHandler {
	return &TransactionHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}
}

// Implement the net/http Handler interface.
func (h TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// Enable all endpoints.
func (h *TransactionHandler) EnableAllEndpoints() *TransactionHandler {
	h.EnableCreateEndpoint()
	h.EnableReadEndpoint()
	h.EnableListEndpoint()
	return h
}

// Enable the create operation.
func (h *TransactionHandler) EnableCreateEndpoint() *TransactionHandler {
	h.r.Post("/", h.Create)
	return h
}

// Enable the read operation.
func (h *TransactionHandler) EnableReadEndpoint() *TransactionHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// Enable the list operation.
func (h *TransactionHandler) EnableListEndpoint() *TransactionHandler {
	h.r.Get("/", h.List)
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
	h.EnableCreateEndpoint()
	h.EnableReadEndpoint()
	h.EnableListEndpoint()
	return h
}

// Enable the create operation.
func (h *UserHandler) EnableCreateEndpoint() *UserHandler {
	h.r.Post("/", h.Create)
	return h
}

// Enable the read operation.
func (h *UserHandler) EnableReadEndpoint() *UserHandler {
	h.r.Get("/{id:\\d+}", h.Read)
	return h
}

// Enable the list operation.
func (h *UserHandler) EnableListEndpoint() *UserHandler {
	h.r.Get("/", h.List)
	return h
}
