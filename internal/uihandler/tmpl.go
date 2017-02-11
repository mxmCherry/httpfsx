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
		<style>{{ template "style" }}</style>
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
		<script type="text/javascript">{{ template "script" }}</script>
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

{{ define "style" }}
.httpfsx {
	font-family: sans-serif;
}

.httpfsx a {
	text-decoration: none;
}

.httpfsx .header {
	margin: 0;
	padding: 0.5cm 0;
	font-size: 1cm;
}

.httpfsx .list {
	margin: 0;
	padding: 0;
	list-style: none;
}

.httpfsx .list .item .star {
	display: inline-block;
	width: 1cm;
	height: 1cm;
	font-size: 1cm;
	line-height: 1cm;
	text-align: center;
	vertical-align: middle;
	color: #808080;
}

.httpfsx .list .item .star.active {
	color: #8B0000;
}

.httpfsx .list .item .link {
	display: inline-block;
	width: calc(100% - 1cm - 0.5cm);
	min-height: 1cm;
	padding: 0;
	vertical-align: top;
}

.httpfsx .list .item .link .name {
	font-size: 0.5cm;
	line-height: 0.6cm;
	color: #101010;
}

.httpfsx .list .item .link .name::after {
	color: #E0E0E0;
}

.httpfsx .list .item.dir .link .name::after {
	content: " 📁";
}

.httpfsx .list .item .link .meta {
	display: block;
	font-size: 0.3cm;
	line-height: 0.6cm;
	color: #808080;
}

.httpfsx .footer {
	font-size: 0.3cm;
	line-height: 0.3cm;
	text-align: right;
	padding-top: 1cm;
}

.httpfsx .footer .clear-storage {
	color: #808080;
}
{{ end }}

{{ define "script" }}
(function() {
	'use strict'

	// root element:
	var httpfsx = document.querySelector('.httpfsx')

	// file-system item nodes:
	var items = httpfsx.querySelectorAll('.httpfsx .list .item')

	// paths, that are listed on current page:
	var existingPaths = []

	for( var i = 0; i < items.length; i++ ) {

		// item wrapper node:
		var item = items[i]

		var star = item.querySelector('.star') // starring element
		var link = item.querySelector('.link') // link element (for item's path detection)

		// item's path:
		var path = link.getAttribute('href')

		// what key is used for storing current item's starred state:
		var starKey = 'httpfsx:star:' + path

		// remember "starring" key to simplify "toggle starring" click handler:
		star.setAttribute('data-httpfsx-star-key', starKey)

		// change star's view, if item is starred:
		if( localStorage.getItem(starKey) ) {
			star.classList.add('active')
		}

		// remember this item's path to remove deleted items from localStorage down the code:
		existingPaths.push(path.replace(/\/{2,}|\/$/g, ''))

	}

	// current request (location) path:
	var currentPath = location.pathname.replace(/\/{2,}|\/$/g, '')

	// traversing localStorage items to clean up deleted ones:
	for( var key in localStorage ) {

		// ignoring any foreign keys:
		if( key.indexOf('httpfsx:') == -1 ) {
			continue
		}

		// extracting stored item path from key:
		var storedPath = key.replace(/httpfsx:[^:]+?:/, '')

		// got item from other path, cannot touch it:
		if( storedPath.indexOf(currentPath) != 0 ) {
			continue
		}

		// does current localStorage item exists (not deleted)?
		var exists = false

		// checking, if current localStorage item is present on current location (page):
		for( var i = 0; i < existingPaths.length; i++ ) {
			var existingPath = existingPaths[i]
			if( storedPath.indexOf(existingPath) == 0 ) {
				exists = true
				break
			}
		}

		// removing deleted file-system items from localStorage:
		if( !exists ) {
			localStorage.removeItem(key)
		}

	}

	// capturing "star" and "clear-storage" clicks:
	httpfsx.addEventListener('click', function(event) {

		if( event.target.classList.contains('star') ) {

			var star = event.target

			var starKey = star.getAttribute('data-httpfsx-star-key')

			// toggle starring status:
			if( localStorage.getItem(starKey) ) {
				localStorage.removeItem(starKey)
				star.classList.remove('active')
			} else {
				localStorage.setItem(starKey, 'T')
				star.classList.add('active')
			}

		} else if( event.target.classList.contains('clear-storage') ) {

			// confirm and clear localStorage:
			if( confirm('Clear storage?') ) {

				localStorage.clear()

				// apply loosing stars to UI:
				var stars = httpfsx.querySelectorAll('.star')
				for( var i = 0; i < stars.length; i++ ) {
					var star = stars[i]
					star.classList.remove('active')
				}

				alert('Storage cleared')
			}

		}

	})

})()
{{ end }}
`))
