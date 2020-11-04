{{ define "sheriff" }}

    {{ $pkg := base $.Config.Package }}
    {{ template "header" $ }}

    import (
        "error"

        "github.com/liip/sheriff"
    )

    {{ range $n := $.Nodes }}
        {{ $receiver := $n.Receiver }}
        func ({{ $receiver }} *{{ $n.Name }}) Marshal(options *sheriff.Options) (interface{}, error) {
            return sheriff.Marshal(options, {{ $receiver }})
        }
    {{ end }}
{{ end }}