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
        {{ if not $n.Annotations.HandlerGen.SkipGeneration }}

            // The {{ $n.Name }}Handler.
            type {{ $n.Name }}Handler struct {
                r *chi.Mux

                client    *ent.Client
                validator *validator.Validate
                logger    *logrus.Logger
            }

            // Create a new {{ $n.Name }}Handler
            func New{{ $n.Name }}Handler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *{{ $n.Name }}Handler {
                return &{{ $n.Name }}Handler{
                    r:         chi.NewRouter(),
                    client:    c,
                    validator: v,
                    logger:    log,
                }
            }

            // Implement the net/http Handler interface.
            func (h {{ $n.Name }}Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
                h.r.ServeHTTP(w, r)
            }

            // Enable all endpoints.
            func (h *{{ $n.Name }}Handler) EnableAllEndpoints() *{{ $n.Name }}Handler {
                h.EnableCreateEndpoint()
                h.EnableReadEndpoint()
                h.EnableListEndpoint()
                return h
            }

            // Enable the create operation.
            func (h *{{ $n.Name }}Handler) EnableCreateEndpoint() *{{ $n.Name }}Handler {
                h.r.Post("/", h.Create)
                return h
            }

            // Enable the read operation.
            func (h *{{ $n.Name }}Handler) EnableReadEndpoint() *{{ $n.Name }}Handler {
                h.r.Get("/{id:\\d+}", h.Read)
                return h
            }

            // Enable the list operation.
            func (h *{{ $n.Name }}Handler) EnableListEndpoint() *{{ $n.Name }}Handler {
                h.r.Get("/", h.List)
                return h
            }

        {{ end }}
    {{ end }}
{{end}}