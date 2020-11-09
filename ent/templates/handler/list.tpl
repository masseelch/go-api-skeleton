{{ define "handler/list" }}

    {{ $pkg := base "handler" }}
    {{- with extend $ "Package" "handler" -}}
        {{ template "header" . }}
    {{ end }}

    import (
        "net/http"
        "strconv"

        "github.com/go-chi/chi"
        "github.com/liip/sheriff"
        "github.com/masseelch/render"
        "github.com/sirupsen/logrus"

        "{{ $.Config.Package }}"
        {{- range $n := $.Nodes}}
            "{{ $.Config.Package }}/{{ $n.Package }}"
        {{- end }}
    )

    {{ range $n := $.Nodes }}
        {{ if not $n.Annotations.HandlerGen.SkipGeneration }}
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

                // Use the query parameters to filter the query. todo - nested filter?
                {{- range $f := $n.Fields }}
                    {{- $jsonName := index (split (tagLookup $f.StructTag "json") ",") 0 }}
                    if f := r.URL.Query().Get("{{ $jsonName }}"); f != "" {
                        {{- if $f.IsBool }}
                            var b bool
                            if f == "true" {
                                b = true
                            } else if f == "false" {
                                b = false
                            } else {
                                h.logger.WithError(err).WithField("{{ $jsonName }}", f).Debug("could not parse query parameter")
                                render.BadRequest(w, r, "'{{ $jsonName }}' must be 'true' or 'false'")
                                return
                            }
                            q = q.Where({{ $n.Package }}.{{$f.StructField}}(b))
                        {{ else if $f.IsInt }}
                            i, err := strconv.Atoi(f)
                            if err != nil {
                                h.logger.WithError(err).WithField("{{ $jsonName }}", f).Debug("could not parse query parameter")
                                render.BadRequest(w, r, "'{{ $jsonName }}' must be an integer")
                                return
                            }
                            q = q.Where({{ $n.Package }}.{{$f.StructField}}(i))
                        {{ else if $f.IsString }}
                            q = q.Where({{ $n.Package }}.{{$f.StructField}}(f))
                        {{ else if $f.IsTime }}
                            // todo
                        {{ end -}}
                    }
                {{ end }}

                es, err := q.All(r.Context())
                if err != nil {
                    h.logger.WithError(err).Error("error querying database") // todo - better error
                    render.InternalServerError(w, r, nil)
                    return
                }

                {{ $groups := $n.Annotations.HandlerGen.ReadGroups }}
                d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{
                    {{- if $groups }}
                        {{- range $g := $groups}}"{{$g}}",{{ end -}}
                    {{ else -}}
                        "{{ $n.Name | snake }}:read"
                    {{- end -}}
                }}, es)
                if err != nil {
                    h.logger.WithError(err).Error("serialization error")
                    render.InternalServerError(w, r, nil)
                    return
                }

                h.logger.WithField("amount", len(es)).Info("jobs rendered")
                render.OK(w, r, d)
            }
        {{ end }}
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