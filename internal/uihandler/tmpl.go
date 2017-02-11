package uihandler

import (
	"html/template"
	"path"

	"github.com/dustin/go-humanize"
	"github.com/mxmCherry/httpfsx/internal/filesystem"
)

var tmpl = template.Must(template.New("ui").Funcs(template.FuncMap{
	"fsType": func(v interface{}) string {
		switch v.(type) {
		case filesystem.Dir:
			return "dir"
		case filesystem.File:
			return "file"
		default:
			return ""
		}
	},
	"fsLink": func(v interface{}) string {
		switch x := v.(type) {
		case filesystem.Dir:
			return path.Join("/fs/explore", x.Path)
		case filesystem.File:
			return path.Join("/fs/raw", x.Path)
		default:
			return ""
		}
	},
	"fsMeta": func(v interface{}) string {
		switch x := v.(type) {
		case filesystem.Dir:
			return "Modified " + humanize.Time(x.LastMod)
		case filesystem.File:
			return humanize.Bytes(uint64(x.Size)) + ", modified " + humanize.Time(x.LastMod)
		default:
			return ""
		}
	},
}).Parse(`
<!DOCTYPE html>
<html class="httpfsx">
	<head>
		<meta name="robots" content="none">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, target-densityDpi=medium-dpi">
		<title>{{ .List.Parent.Name }}</title>
		<link rel="stylesheet" type="text/css" href="/fs/static/style.css">
	</head>
	<body>
		<h1 class="header">{{ .List.Parent.Name }}</h1>
		<ul class="list">

			{{ range .List.Dirs }}
				{{ template "fsItemView" . }}
			{{ end }}

			{{ range .List.Files }}
				{{ template "fsItemView" . }}
			{{ end }}

		</ul>
		<footer class="footer">
			<a class="clear-storage" href="javascript://">Clear storage</a>
		</footer>
		<script type="text/javascript" src="/fs/static/script.js"></script>
	<body>
</html>

{{ define "fsItemView" }}
	<li class="item {{ fsType . }}">
		<a class="star" href="javascript://">★</a>
		<a class="link" href="{{ fsLink . }}">
			<span class="name">{{ .Name }}</span>
			<small class="meta">{{ fsMeta . }}</small>
		</a>
	</li>
{{ end }}
`))
