{{ define "handler/read" }}

    {{ $pkg := base "handler" }}
    {{- with extend $ "Package" "handler" -}}
        {{ template "header" . }}
    {{ end }}

    import (
    "github.com/go-chi/chi"
    "github.com/go-playground/validator/v10"
    "github.com/masseelch/render"
    "github.com/sirupsen/logrus"
    "net/http"
    "strconv"

    "{{ $.Config.Package }}"
    )

    {{ range $n := $.Nodes }}
        {{ if not (eq $n.Name "Session") }}
            func(h {{ $n.Name }}Handler) read(w http.ResponseWriter, r *http.Request) {
            id, err := strconv.Atoi(chi.URLParam(r, "id"))
            if err != nil {
            h.logger.WithField("id", chi.URLParam(r, "id")).Info("error parsing url parameter 'id'")
            render.BadRequest(w, r, "id must be a positive integer greater zero")
            return
            }

            e, err := h.client.{{ $n.Name }}.Get(r.Context(), id)
            if err != nil {
            switch err.(type) {
            case *ent.NotFoundError:
            h.logger.WithError(err).Debug("job not found")
            render.NotFound(w, r, err)
            return
            case *ent.NotSingularError:
            h.logger.WithError(err).Error("unexpected")                  // todo - better error
            render.InternalServerError(w, r, "unexpected error occured") // todo - better error
            return
            default:
            h.logger.WithError(err).Error("logic") // todo - better stuff here pls
            render.InternalServerError(w, r, "logic")
            return
            }
            }

            d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"job:list"}}, e)
            if err != nil {
            h.logger.WithError(err).Error("sheriff") // todo - better stuff here pls
            render.InternalServerError(w, r, "sheriff")
            return
            }

            h.logger.WithField("job", e.ID).Info("job rendered") // todo - better stuff here pls
            render.OK(w, r, d)
            }
        {{end}}
    {{ end }}

{{end}}