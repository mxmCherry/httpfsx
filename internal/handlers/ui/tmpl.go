package ui

const tmplCode = `
<!DOCTYPE html>
<html class="httpfsx">
	<head>
		<meta name="robots" content="none">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, target-densityDpi=medium-dpi">
		<title>{{ .List.Parent.Name }}</title>
		<link rel="stylesheet" type="text/css" href="{{ staticLink "/style.css" }}">
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
			<button class="clear-storage">Clear storage</button>
		</footer>
		<script type="text/javascript" src="{{ staticLink "/script.js" }}"></script>
	<body>
</html>

{{ define "fsItemView" }}
	<li class="item {{ fsType . }} {{ metaKind . }}">
		<button class="star">★</button>
		<a class="link" href="{{ fsLink . }}">
			<span class="name">{{ .Name }}</span>
			<small class="meta">{{ fsMeta . }}</small>
		</a>
	</li>
{{ end }}
`
