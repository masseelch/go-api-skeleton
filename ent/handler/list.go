// Code generated by entc, DO NOT EDIT.

package handler

import (
	"net/http"
	"strconv"

	"github.com/liip/sheriff"
	"github.com/masseelch/render"
	"github.com/sirupsen/logrus"

	"github.com/masseelch/go-api-skeleton/ent/group"
	"github.com/masseelch/go-api-skeleton/ent/job"
	"github.com/masseelch/go-api-skeleton/ent/user"
)

// This function queries for Group models. Can be filtered by query parameters.
func (h GroupHandler) List(w http.ResponseWriter, r *http.Request) {
	q := h.client.Group.Query()

	// Pagination. Default is 30 items per page.
	page, itemsPerPage, err := pagination(w, r, h.logger)
	if err != nil {
		return
	}
	q = q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query.
	if f := r.URL.Query().Get("title"); f != "" {
		q = q.Where(group.Title(f))
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"group:read"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("jobs rendered")
	render.OK(w, r, d)
}

// This function queries for Job models. Can be filtered by query parameters.
func (h JobHandler) List(w http.ResponseWriter, r *http.Request) {
	q := h.client.Job.Query().WithUsers()

	// Pagination. Default is 30 items per page.
	page, itemsPerPage, err := pagination(w, r, h.logger)
	if err != nil {
		return
	}
	q = q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query.
	if f := r.URL.Query().Get("date"); f != "" {
		// todo
	}

	if f := r.URL.Query().Get("task"); f != "" {
		q = q.Where(job.Task(f))
	}

	if f := r.URL.Query().Get("state"); f != "" {
		q = q.Where(job.State(f))
	}

	if f := r.URL.Query().Get("report"); f != "" {
		q = q.Where(job.Report(f))
	}

	if f := r.URL.Query().Get("rest"); f != "" {
		q = q.Where(job.Rest(f))
	}

	if f := r.URL.Query().Get("note"); f != "" {
		q = q.Where(job.Note(f))
	}

	if f := r.URL.Query().Get("customerName"); f != "" {
		q = q.Where(job.CustomerName(f))
	}

	if f := r.URL.Query().Get("riskAssessmentRequired"); f != "" {
		var b bool
		if f == "true" {
			b = true
		} else if f == "false" {
			b = false
		} else {
			h.logger.WithError(err).WithField("riskAssessmentRequired", f).Debug("could not parse query parameter")
			render.BadRequest(w, r, "'riskAssessmentRequired' must be 'true' or 'false'")
			return
		}
		q = q.Where(job.RiskAssessmentRequired(b))
	}

	if f := r.URL.Query().Get("maintenanceRequired"); f != "" {
		var b bool
		if f == "true" {
			b = true
		} else if f == "false" {
			b = false
		} else {
			h.logger.WithError(err).WithField("maintenanceRequired", f).Debug("could not parse query parameter")
			render.BadRequest(w, r, "'maintenanceRequired' must be 'true' or 'false'")
			return
		}
		q = q.Where(job.MaintenanceRequired(b))
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"job:list", "user:list"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("jobs rendered")
	render.OK(w, r, d)
}

// This function queries for User models. Can be filtered by query parameters.
func (h UserHandler) List(w http.ResponseWriter, r *http.Request) {
	q := h.client.User.Query()

	// Pagination. Default is 30 items per page.
	page, itemsPerPage, err := pagination(w, r, h.logger)
	if err != nil {
		return
	}
	q = q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query.
	if f := r.URL.Query().Get("email"); f != "" {
		q = q.Where(user.Email(f))
	}

	if f := r.URL.Query().Get("password"); f != "" {
		q = q.Where(user.Password(f))
	}

	if f := r.URL.Query().Get("enabled"); f != "" {
		var b bool
		if f == "true" {
			b = true
		} else if f == "false" {
			b = false
		} else {
			h.logger.WithError(err).WithField("enabled", f).Debug("could not parse query parameter")
			render.BadRequest(w, r, "'enabled' must be 'true' or 'false'")
			return
		}
		q = q.Where(user.Enabled(b))
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"user:read"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("jobs rendered")
	render.OK(w, r, d)
}

func pagination(w http.ResponseWriter, r *http.Request, l *logrus.Logger) (page int, itemsPerPage int, err error) {
	page = 1
	itemsPerPage = 30

	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.WithField("itemsPerPage", d).Info("error parsing query parameter 'itemsPerPage'")
			render.BadRequest(w, r, "itemsPerPage must be a positive integer greater zero")
			return
		}
	}

	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.WithField("page", d).Info("error parsing query parameter 'page'")
			render.BadRequest(w, r, "page must be a positive integer greater zero")
			return
		}
	}

	return
}
