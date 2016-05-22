/*

Command httpfsx launches mobile-friendly HTTP file-system explorer (readonly)

Basic usage:
	httpfsx --addr=tcp://:1024 --root=$HOME/share

*/
package main

import (
	"flag"
	"html"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/valyala/fasthttp"
)

// ----------------------------------------------------------------------------

// Config holds all the data, needed to bootstrap an application.
type Config struct {
	Addr string // address to serve on
	Root string // public root directory path
}

// ----------------------------------------------------------------------------

func main() {

	var err error

	config := Config{
		Addr: "tcp://127.0.0.1:1024",
		Root: "./",
	}

	flagSet := flag.NewFlagSet("main", flag.ExitOnError)
	flagSet.StringVar(&config.Addr, "addr", config.Addr, "Address to serve on")
	flagSet.StringVar(&config.Root, "root", config.Root, "Public root directory path")
	flagSet.Parse(os.Args[1:])

	// validating addr to provide more informative error message (http.Listen is not so detailed):
	if !regexp.MustCompile("^(?:(?:tcp[46]?)://(?:.*?):\\d{1,})|(?:unix(?:packet)?://.+?)$").MatchString(config.Addr) {
		os.Stderr.WriteString("Error: --addr should be provided in a form of PROTO://ADDR, where PROTO is tcp, tcp4, tcp6, unix or unixpacket and ADDR is a HOST:PORT combination or /path/to/socket.unix")
		flag.Usage()
		os.Exit(1)
		return
	}

	// determining absolute public root path (to be informative; also, this allows to provide non-empty header for root location):
	config.Root = filepath.Clean(config.Root)
	if !filepath.IsAbs(config.Root) {
		wd, err := os.Getwd()
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
			return
		}
		config.Root = filepath.Join(wd, config.Root)
	}

	// parsing address string (decompose it into schema + host):
	addrURL, err := url.Parse(config.Addr)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}

	// create listener (doing it right now helps to detect errors like "address already in use"):
	listener, err := net.Listen(addrURL.Scheme, addrURL.Host)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)

	handler := makeHandler(config.Root)

	server := fasthttp.Server{
		Handler: handler,
		Name:    "httpfsx v0.0.6", // TODO: don't forget to change this before creating release!
		GetOnly: true,
		Logger:  logger,
	}

	// notify user about current settings:
	logger.Printf("Addr: %s", addrURL.String())
	logger.Printf("Root: %s", config.Root)

	err = server.Serve(listener)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}
}

// ----------------------------------------------------------------------------

// decomposeOsError tries to determine HTTP status, suitable for passed os.* error.
func decomposeOsError(err error) (string, int) {

	if os.IsNotExist(err) {
		return err.Error(), fasthttp.StatusNotFound
	}

	if os.IsPermission(err) {
		return err.Error(), fasthttp.StatusForbidden
	}

	return err.Error(), fasthttp.StatusInternalServerError
}

// ----------------------------------------------------------------------------

// makeHandler generates new fasthttp handler closure for given root path.
func makeHandler(root string) fasthttp.RequestHandler {
	fs := &fasthttp.FS{
		Root:            root,
		Compress:        false,
		AcceptByteRange: true,
	}
	fsHandler := fs.NewRequestHandler()

	return func(ctx *fasthttp.RequestCtx) {

		relPath := path.Join("/", string(ctx.Path())) // relative request item path
		absPath := filepath.Join(root, relPath)       // absolute file-system item path

		// check if requested item exists and what is it's type:
		stats, err := os.Stat(absPath)
		if err != nil {
			ctx.Error(decomposeOsError(err))
			return
		}

		// serve dir's listing:
		if stats.IsDir() {

			// get dir contents:
			fis, err := ioutil.ReadDir(absPath)
			if err != nil {
				ctx.Error(decomposeOsError(err))
				return
			}

			escapedDirName := html.EscapeString(stats.Name())

			page := `
				<!DOCTYPE html>
				<html class="httpfsx">
					<head>
						<meta name="robots" content="none">
						<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, target-densityDpi=medium-dpi">
						<title>` + escapedDirName + `</title>
						<style>` + style + `</style>
					</head>
					<body>
						<h1 class="header">` + escapedDirName + `</h1>
						<ul class="list">
			`

			for _, fi := range fis {

				name := fi.Name()

				// skip (don't list) dotted (hidden) files:
				if strings.HasPrefix(name, ".") {
					continue
				}

				itemType := "file" // dir child type (file/dir)
				meta := ""         // some human-readable metadata about current child

				if fi.IsDir() {
					itemType = "dir"
					// for dirs only human-readable mod time is available:
					meta = "Modified " + humanize.Time(fi.ModTime())
				} else {
					// for files both human-readable size and mod time are available:
					meta = humanize.Bytes(uint64(fi.Size())) + ", modified " + humanize.Time(fi.ModTime())
				}

				page += `
					<li class="item ` + itemType + `">
						<a class="star" href="javascript://">★</a>
						<a class="link" href="` + html.EscapeString(path.Join(relPath, name)) + `">
							<span class="name">` + html.EscapeString(name) + `</span>
							<small class="meta">` + html.EscapeString(meta) + `</small>
						</a>
					</li>
				`
			}

			page += `
						</ul>
						<footer class="footer">
							<a class="clear-storage" href="javascript://">Clear storage</a>
						</footer>
						<script type="text/javascript">` + script + `</script>
					<body>
				</html>
			`

			ctx.SetContentType("text/html; charset=utf-8")
			ctx.SetBodyString(page)

			return
		}

		// serve file:
		fsHandler(ctx)
	}
}

// ----------------------------------------------------------------------------

// httpfsx CSS:
var style string = `
	
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
	
`

// ----------------------------------------------------------------------------

// httpfsx JS:
var script string = `

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

`
