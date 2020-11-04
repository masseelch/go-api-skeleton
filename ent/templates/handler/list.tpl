{{ define "handler/list" }}

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
        {{- range $n := $.Nodes}}
            "{{ $.Config.Package }}/{{ $n.Package }}"
        {{- end }}
    )

    {{ range $n := $.Nodes }}
        // This function queries for {{ $n.Name }} models. Can be filtered by query parameters.
        func(h {{ $n.Name }}Handler) List(w http.ResponseWriter, r *http.Request) {
            q := h.client.{{ $n.Name }}.Query()
            {{- range $e := $n.Edges }}
                {{- range $l := $n.Annotations.HandlerGen.ListEager}}
                    {{- if eq $l $e.Name }}.With{{ pascal $e.Name }}(){{ end -}}
                {{ end -}}
            {{ end }}

            // Pagination. Default is 30 items per page.
            page, itemsPerPage, err := pagination(w, r, h.logger)
            if err != nil {
                return
            }
            q = q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

            // Use the query parameters to filter the query.
{{/*            {{ range $f := $n.Fields }}*/}}
{{/*                if f := r.URL.Query().Get("{{ tagLookup $f.StructTag "json"}}"); f != "" {*/}}
{{/*                    // {{ $f.BasicType "string" }}*/}}
{{/*                    q = q.Where({{ $n.Package }}.{{$f.StructField}}(f))*/}}
{{/*                }*/}}
{{/*            {{ end }}*/}}

            es, err := q.All(r.Context())
            if err != nil {
                h.logger.WithError(err).Error("unexpected") // todo - better error
                render.InternalServerError(w, r, "logic")
                return
{{/*                switch err.(type) {*/}}
{{/*                    case *ent.NotFoundError:*/}}
{{/*                        h.logger.WithError(err).Debug("job not found")*/}}
{{/*                        render.NotFound(w, r, err)*/}}
{{/*                        return*/}}
{{/*                    case *ent.NotSingularError:*/}}
{{/*                        h.logger.WithError(err).Error("unexpected")                  // todo - better error*/}}
{{/*                        render.InternalServerError(w, r, "unexpected error occurred") // todo - better error*/}}
{{/*                        return*/}}
{{/*                    default:*/}}
{{/*                        h.logger.WithError(err).Error("logic") // todo - better stuff here pls*/}}
{{/*                        render.InternalServerError(w, r, "logic")*/}}
{{/*                        return*/}}
{{/*                }*/}}
            }

            h.logger.WithField("amount", len(es)).Info("jobs rendered") // todo - better stuff here pls
            render.OK(w, r, es)
        }

    {{ end }}

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

{{end}}