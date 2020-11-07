{{ define "handler/create" }}

    {{ $pkg := base "handler" }}
    {{- with extend $ "Package" "handler" -}}
        {{ template "header" . }}
    {{ end }}

    import (
        "net/http"
        "strconv"

        "github.com/go-chi/chi"
        "github.com/go-playground/validator/v10"
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
        // struct to bind the post body to.
        type {{ $n.Name | camel }}CreateRequest struct {
            {{ range $f := $n.Fields -}}
                {{- $a := $f.Annotations.FieldGen }}
                {{- if or (not $a) $a.Create }}
                    {{ $f.StructField }} {{ $f.Type.String }} `{{ $f.StructTag }}`
                {{- end -}}
            {{ end }}
        }

        // This function creates a new {{ $n.Name }} model and stores it in the database.
        func(h {{ $n.Name }}Handler) Create(w http.ResponseWriter, r *http.Request) {
            // Get the post data.
            d := {{ $n.Name | snake }}CreateRequest{} // todo - allow form-url-encdoded/xml/protobuf data.
            if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                h.logger.WithError(err).Error("error decoding json")
                render.BadRequest(w, r, "invalid json string")
                return
            }

            // Validate the data.
            if err := h.validator.Struct(d); err != nil {
                if err, ok := err.(*validator.InvalidValidationError); ok {
                    h.logger.WithError(err).Error("error validating request data")
                    render.InternalServerError(w, r, nil)
                    return
                }

                h.logger.WithError(err).Info("validation failed")
                render.BadRequest(w, r, err)
                return
            }

            // Save the data.
            b := h.client.{{ $n.Name }}.Create()
            {{ range $f := $n.Fields -}}
                {{- $a := $f.Annotations.FieldGen }}
                {{- if or (not $a) $a.Create }}
                    b = b.Set{{ $f.StructField }}(d.{{ $f.StructField }})
                {{- end -}}
            {{ end }}

            // Store in database.
            e, err := b.Save(r.Context())
            if err != nil {
                h.logger.WithError(err).Error("error saving {{ $n.Name }}")
                render.InternalServerError(w, r, nil)
                return
            }

            {{ $groups := $n.Annotations.HandlerGen.CreateGroups }}
            j, err := sheriff.Marshal(&sheriff.Options{Groups: []string{
                {{- if $groups }}
                    {{- range $g := $groups}}"{{$g}}",{{ end -}}
                {{ else -}}
                    "{{ $n.Name | snake }}:create"
                {{- end -}}
            }}, e)
            if err != nil {
                h.logger.WithError(err).WithField("{{ $n.Name }}.{{ $n.ID.Name }}", e.ID).Error("serialization error")
                render.InternalServerError(w, r, nil)
                return
            }

            h.logger.WithField("{{ $n.Name | snake }}", e.ID).Info("{{ $n.Name | snake }} rendered")
            render.OK(w, r, j)
        }

    {{ end }}

{{ end }}