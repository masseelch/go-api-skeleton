// Code generated by entc, DO NOT EDIT.

package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/liip/sheriff"
	"github.com/masseelch/render"

	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/ent/job"
	"github.com/masseelch/go-api-skeleton/ent/session"
	"github.com/masseelch/go-api-skeleton/ent/user"

	go_token "github.com/masseelch/go-token"
)

// This function fetches the Job model identified by a give url-parameter from
// database and returns it to the client.
func (h JobHandler) Read(w http.ResponseWriter, r *http.Request) {
	idp := chi.URLParam(r, "id")
	if idp == "" {
		h.logger.WithField("id", idp).Info("empty 'id' url param")
		render.BadRequest(w, r, "id cannot be ''")
		return
	}
	id, err := strconv.Atoi(idp)
	if err != nil {
		h.logger.WithField("id", idp).Info("error parsing url parameter 'id'")
		render.BadRequest(w, r, "id must be a positive integer greater zero")
		return
	}

	e, err := h.client.Job.Query().Where(job.ID(id)).WithUsers().Only(r.Context())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("Job.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("Job.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("Job.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"job:list", "user:list"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Job.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("job", e.ID).Info("job rendered")
	render.OK(w, r, d)
}

// This function fetches the Session model identified by a give url-parameter from
// database and returns it to the client.
func (h SessionHandler) Read(w http.ResponseWriter, r *http.Request) {
	idp := chi.URLParam(r, "id")
	if idp == "" {
		h.logger.WithField("id", idp).Info("empty 'id' url param")
		render.BadRequest(w, r, "id cannot be ''")
		return
	}
	id := go_token.Token(idp)

	e, err := h.client.Session.Query().Where(session.ID(id)).Only(r.Context())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("Session.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("Session.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("Session.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"session:list"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Session.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("session", e.ID).Info("job rendered")
	render.OK(w, r, d)
}

// This function fetches the User model identified by a give url-parameter from
// database and returns it to the client.
func (h UserHandler) Read(w http.ResponseWriter, r *http.Request) {
	idp := chi.URLParam(r, "id")
	if idp == "" {
		h.logger.WithField("id", idp).Info("empty 'id' url param")
		render.BadRequest(w, r, "id cannot be ''")
		return
	}
	id, err := strconv.Atoi(idp)
	if err != nil {
		h.logger.WithField("id", idp).Info("error parsing url parameter 'id'")
		render.BadRequest(w, r, "id must be a positive integer greater zero")
		return
	}

	e, err := h.client.User.Query().Where(user.ID(id)).Only(r.Context())
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("User.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("User.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("User.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"user:list"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("User.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("user", e.ID).Info("job rendered")
	render.OK(w, r, d)
}
