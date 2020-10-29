package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/sirupsen/logrus"
	"net/http"
)

type jobHandler struct {
	r *chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

func NewJobHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *jobHandler {
	h := &jobHandler{
		r:         chi.NewRouter(),
		client:    c,
		validator: v,
		logger:    log,
	}

	h.r.Get("/", h.read)

	return h
}

func (h jobHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h jobHandler) read(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello There"))
}

// func JobReadHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id, err := strconv.Atoi(chi.URLParam(r, "id"))
// 		if err != nil {
// 			render.BadRequest(w, r, "id must be a positive integer greater zero")
// 			return
// 		}
//
// 		e, err := h.services.User.Find(uint(id))
// 		if err != nil {
// 			render.NotFound(w, r, "User not found.")
// 			return
// 		}
//
// 		// Check if the logged in user has permission to read this resource.
// 		if !h.voters.IsGranted("read", e, h.services.User.UserFromContext(r.Context())) {
// 			render.Forbidden(w, r, nil)
// 			return
// 		}
//
// 		d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"read"}}, e.View())
// 		if err != nil {
// 			h.log.WithError(err).WithFields(logrus.Fields{
// 				"caller": "userHandler",
// 				"func":   "read()",
// 			}).Error("serialization error")
// 			render.InternalServerError(w, r, nil)
// 			return
// 		}
//
// 		render.OK(w, r, d)
// 	}
// }
