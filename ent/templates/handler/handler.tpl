{{ define "handler/handler" }}

    {{ $pkg := base "handler" }}
    {{- with extend $ "Package" "handler" -}}
        {{ template "header" . }}
    {{ end }}

    import (
    "github.com/go-chi/chi"
    "github.com/go-playground/validator/v10"
    "github.com/sirupsen/logrus"
    "net/http"

    "{{ $.Config.Package }}"
    )

    {{ range $n := $.Nodes }}
        {{ if not (eq $n.Name "Session") }}
            {{/* The handler struct */}}
            type {{ $n.Name }}Handler struct {
            r *chi.Mux

            client    *ent.Client
            validator *validator.Validate
            logger    *logrus.Logger
            }

            {{/* Create a new handler */}}
            func New{{ $n.Name }}Handler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *{{ $n.Name }}Handler {
            h := &{{ $n.Name }}Handler{
            r:         chi.NewRouter(),
            client:    c,
            validator: v,
            logger:    log,
            }

            h.r.Get("/{id:\\d+}", h.read)

            return h
            }

            {{/* Implement the net/http Handler interface */}}
            func (h {{ $n.Name }}Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
            h.r.ServeHTTP(w, r)
            }
        {{end}}
    {{ end }}

{{end}}