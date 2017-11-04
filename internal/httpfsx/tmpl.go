package httpfsx

import (
	"html/template"
	"path"
)

var tmpl = template.Must(template.New("dir").Funcs(template.FuncMap{
	"joinPath": func(components ...string) string {
		return path.Join(components...)
	},
}).Parse(`
	{{ $ctx := . }}
	<!DOCTYPE html>
	<html>
		<head>
			<title>{{ $ctx.Self.Name }}</title>
			<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
			<style>
				body {
					margin: 0;
					color: #111;
				}
				h1 {
					padding: 0 2mm;
					margin: 2mm 0;
				}
				ol {
					list-style: none;
					padding: 0;
					margin: 0;
				}
				li:nth-child(odd) {
					background: #EEE;
				}
				a {
					display: inline-block;
					box-sizing: border-box;
					width: 100%;
					line-height: 6mm;
					color: #111;
					padding: 2mm;
				}
				a:visited {
					color: #666;
				}
				a span {
					display: inline-block;
					vertical-align: middle;
					line-height: normal;
					word-wrap: break-word;
					max-width: 100%;
				}
			</style>
		</head>
		<body>
			<h1>{{ $ctx.Self.Name }}</h1>
			<ol>
				{{ range $dir := $ctx.Dirs }}
					<li><a href="{{ joinPath $ctx.Base $dir.Name }}"><span>{{ $dir.Name }}/</span></a></li>
				{{ end }}
				{{ range $file := $ctx.Files }}
					<li><a href="{{ joinPath $ctx.Base $file.Name }}"><span>{{ $file.Name }}</span></a></li>
				{{ end }}
			</ol>
		</body>
	</html>
`))
