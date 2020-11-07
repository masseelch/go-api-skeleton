{{ define "handler/read" }}

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

        {{/* Import a possible custom GoType for the ID. */}}
        {{- range $n := $.Nodes}}
            {{- if $n.ID.HasGoType }}
                "{{ $n.ID.Type.PkgPath }}"
            {{- end }}
        {{- end}}
    )

    {{ range $n := $.Nodes }}
        {{ if not $n.Annotations.HandlerGen.SkipGeneration }}
            // This function fetches the {{ $n.Name }} model identified by a give url-parameter from
            // database and returns it to the client.
            func(h {{ $n.Name }}Handler) Read(w http.ResponseWriter, r *http.Request) {
                idp := chi.URLParam(r, "id")
                if idp == "" {
                    h.logger.WithField("id", idp).Info("empty 'id' url param")
                    render.BadRequest(w, r, "id cannot be ''")
                    return
                }
                {{ if $n.ID.HasGoType -}}
                    id := {{ $n.ID.Type.String }}(idp)
                {{ else if $n.ID.IsString -}}
                    id := idp
                {{ else if $n.ID.IsInt -}}
                    id, err := strconv.Atoi(idp)
                    if err != nil {
                        h.logger.WithField("id", idp).Info("error parsing url parameter 'id'")
                        render.BadRequest(w, r, "id must be a positive integer greater zero")
                        return
                    }
                {{- end}}

                {{/* If one of the given handler groups is set on the edge eager join it.*/}}
                // todo - nested eager loading?
                e, err := h.client.{{ $n.Name }}.Query().Where({{ $n.Name | snake }}.ID(id)).
                {{- range $e := $n.Edges }}
                    {{- range $g := $n.Annotations.HandlerGen.ReadGroups }}
                        {{- range $eg := split (tagLookup $e.StructTag "groups") "," }}
                            {{- if eq $g $eg }}With{{ pascal $e.Name }}().{{ end -}}
                        {{- end }}
                    {{- end }}
                {{- end -}}
                Only(r.Context())
                if err != nil {
                    switch err.(type) {
                        case *ent.NotFoundError:
                            h.logger.WithError(err).WithField("{{ $n.Name }}.{{ $n.ID.Name }}", id).Debug("job not found")
                            render.NotFound(w, r, err)
                            return
                        case *ent.NotSingularError:
                            h.logger.WithError(err).WithField("{{ $n.Name }}.{{ $n.ID.Name }}", id).Error("duplicate entry for id")
                            render.InternalServerError(w, r, nil)
                            return
                        default:
                            h.logger.WithError(err).WithField("{{ $n.Name }}.{{ $n.ID.Name }}", id).Error("error fetching node from db")
                            render.InternalServerError(w, r, nil)
                            return
                    }
                }

                {{ $groups := $n.Annotations.HandlerGen.ReadGroups }}
                d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{
                    {{- if $groups }}
                        {{- range $g := $groups}}"{{$g}}",{{ end -}}
                    {{ else -}}
                        "{{ $n.Name | snake }}:list"
                    {{- end -}}
                }}, e)
                if err != nil {
                    h.logger.WithError(err).WithField("{{ $n.Name }}.{{ $n.ID.Name }}", id).Error("serialization error")
                    render.InternalServerError(w, r, nil)
                    return
                }

                h.logger.WithField("{{ $n.Name | snake }}", e.ID).Info("{{ $n.Name | snake }} rendered")
                render.OK(w, r, d)
            }

        {{ end }}
    {{ end }}
{{end}}