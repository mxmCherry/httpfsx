package main

import (
	flag "flag"
	html "html"
	io "io"
	ioutil "io/ioutil"
	log "log"
	mime "mime"
	net "net"
	http "net/http"
	url "net/url"
	os "os"
	path "path"
	regexp "regexp"
	strings "strings"
	time "time"
)

import (
	humanize "github.com/dustin/go-humanize"
	fasthttp "github.com/valyala/fasthttp"
)

// ----------------------------------------------------------------------------

type Config struct {
	Addr string
	Root string

	FasthttpConcurrency          uint64
	FasthttpReadBufferSize       uint64
	FasthttpWriteBufferSize      uint64
	FasthttpReadTimeout          uint64
	FasthttpWriteTimeout         uint64
	FasthttpMaxConnsPerIP        uint64
	FasthttpMaxRequestsPerConn   uint64
	FasthttpMaxKeepaliveDuration uint64
	FasthttpMaxRequestBodySize   uint64
	FasthttpReduceMemoryUsage    bool
	FasthttpGetOnly              bool
}

// ----------------------------------------------------------------------------

func main() {

	var err error

	config := Config{
		Addr: "tcp://127.0.0.1:1024",
		Root: "./",

		FasthttpConcurrency:          0,
		FasthttpReadBufferSize:       0,
		FasthttpWriteBufferSize:      0,
		FasthttpReadTimeout:          0,
		FasthttpWriteTimeout:         0,
		FasthttpMaxConnsPerIP:        0,
		FasthttpMaxRequestsPerConn:   0,
		FasthttpMaxKeepaliveDuration: 0,
		FasthttpMaxRequestBodySize:   0,
		FasthttpReduceMemoryUsage:    false,
		FasthttpGetOnly:              true,
	}

	flagSet := flag.NewFlagSet("main", flag.ExitOnError)

	flagSet.StringVar(&config.Addr, "addr", config.Addr, "Address to serve on")
	flagSet.StringVar(&config.Root, "root", config.Root, "Public root directory path")

	flagSet.Uint64Var(&config.FasthttpConcurrency, "fasthttp-concurrency", config.FasthttpConcurrency, "The maximum number of concurrent connections the server may serve")
	flagSet.Uint64Var(&config.FasthttpReadBufferSize, "fasthttp-read-buffer-size", config.FasthttpReadBufferSize, "Per-connection buffer size for requests' reading. This also limits the maximum header size; bytes")
	flagSet.Uint64Var(&config.FasthttpWriteBufferSize, "fasthttp-write-buffer-size", config.FasthttpWriteBufferSize, "Per-connection buffer size for responses' writing; bytes")
	flagSet.Uint64Var(&config.FasthttpReadTimeout, "fasthttp-read-timeout", config.FasthttpReadTimeout, "Maximum duration for full request reading (including body); milliseconds")
	flagSet.Uint64Var(&config.FasthttpWriteTimeout, "fasthttp-write-timeout", config.FasthttpWriteTimeout, "Maximum duration for full response writing (including body); milliseconds")
	flagSet.Uint64Var(&config.FasthttpMaxConnsPerIP, "fasthttp-max-conns-per-ip", config.FasthttpMaxConnsPerIP, "Maximum number of concurrent client connections allowed per IP")
	flagSet.Uint64Var(&config.FasthttpMaxRequestsPerConn, "fasthttp-max-requests-per-conn", config.FasthttpMaxRequestsPerConn, "Maximum number of requests served per connection")
	flagSet.Uint64Var(&config.FasthttpMaxKeepaliveDuration, "fasthttp-max-keepalive-duration", config.FasthttpMaxKeepaliveDuration, "Maximum keep-alive connection lifetime; milliseconds")
	flagSet.Uint64Var(&config.FasthttpMaxRequestBodySize, "fasthttp-max-request-body-size", config.FasthttpMaxRequestBodySize, "Maximum request body size; bytes")
	flagSet.BoolVar(&config.FasthttpReduceMemoryUsage, "fasthttp-reduce-memory-usage", config.FasthttpReduceMemoryUsage, "Aggressively reduces memory usage at the cost of higher CPU usage if set to true")
	flagSet.BoolVar(&config.FasthttpGetOnly, "fasthttp-get-only", config.FasthttpGetOnly, "Rejects all non-GET requests if set to true")

	flagSet.Parse(os.Args[1:])

	if !regexp.MustCompile("^(?:(?:tcp[46]?)://(?:.*?):\\d{1,})|(?:unix(?:packet)?://.+?)$").MatchString(config.Addr) {
		os.Stderr.WriteString("Error: --addr should be provided in a form of PROTO://ADDR, where PROTO is tcp, tcp4, tcp6, unix or unixpacket and ADDR is a HOST:PORT combination or /path/to/socket.unix")
		flag.Usage()
		os.Exit(1)
		return
	}

	config.Root = path.Clean(config.Root)
	if !path.IsAbs(config.Root) {
		wd, err := os.Getwd()
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
			return
		}
		config.Root = path.Join(wd, config.Root)
	}

	addrURL, err := url.Parse(config.Addr)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}

	listener, err := net.Listen(addrURL.Scheme, addrURL.Host)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)

	handler := makeHandler(config.Root)

	server := fasthttp.Server{
		Handler:              handler,
		Name:                 "httpfsx v0.0.1",
		Concurrency:          int(config.FasthttpConcurrency),
		ReadBufferSize:       int(config.FasthttpReadBufferSize),
		WriteBufferSize:      int(config.FasthttpWriteBufferSize),
		ReadTimeout:          time.Millisecond * time.Duration(config.FasthttpReadTimeout),
		WriteTimeout:         time.Millisecond * time.Duration(config.FasthttpWriteTimeout),
		MaxConnsPerIP:        int(config.FasthttpMaxConnsPerIP),
		MaxRequestsPerConn:   int(config.FasthttpMaxRequestsPerConn),
		MaxKeepaliveDuration: time.Millisecond * time.Duration(config.FasthttpMaxKeepaliveDuration),
		MaxRequestBodySize:   int(config.FasthttpMaxRequestBodySize),
		ReduceMemoryUsage:    config.FasthttpReduceMemoryUsage,
		GetOnly:              config.FasthttpGetOnly,
		Logger:               logger,
	}

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

func makeHandler(root string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		relPath := path.Join("/", string(ctx.Path()))
		absPath := path.Join(root, relPath)

		stats, err := os.Stat(absPath)
		if err != nil {
			ctx.Error(decomposeOsError(err))
			return
		}

		if stats.IsDir() {

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
				if strings.HasPrefix(name, ".") {
					continue
				}
				itemType := "file"
				meta := ""
				if fi.IsDir() {
					itemType = " dir"
					meta = "Modified " + humanize.Time(fi.ModTime())
				} else {
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

		mimeType := mime.TypeByExtension(path.Ext(absPath))

		if mimeType != "" {
			ctx.SetContentType(mimeType)
			ctx.SendFile(absPath)
			return
		}

		file, err := os.Open(absPath)
		if err != nil {
			ctx.Error(decomposeOsError(err))
			return
		}

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil && err != io.EOF {
			ctx.Error(decomposeOsError(err))
			return
		}

		_, err = file.Seek(0, os.SEEK_SET)
		if err != nil {
			ctx.Error(decomposeOsError(err))
			return
		}

		mimeType = http.DetectContentType(buf)

		ctx.SetContentType(mimeType)
		ctx.SetBodyStream(file, int(stats.Size()))

	}

}

// ----------------------------------------------------------------------------

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

var script string = `

'use strict'

var httpfsx = document.querySelector('.httpfsx')

var items = httpfsx.querySelectorAll('.httpfsx .list .item')

var existingPaths = []

for( var i = 0; i < items.length; i++ ) {
	
	var item = items[i]
	
	var star = item.querySelector('.star')
	var link = item.querySelector('.link')
	
	var path = link.getAttribute('href')
	
	var starKey = 'httpfsx:star:' + path
	
	star.setAttribute('data-httpfsx-star-key', starKey)
	
	if( localStorage.getItem(starKey) ) {
		star.classList.add('active')
	}
	
	existingPaths.push(path.replace(/\/{2,}|\/$/g, ''))
	
}

var currentPath = location.pathname.replace(/\/{2,}|\/$/g, '')

for( var key in localStorage ) {
	
	if( key.indexOf('httpfsx:') == -1 ) {
		continue
	}
	
	var storedPath = key.replace(/httpfsx:[^:]+?:/, '')
	
	if( storedPath.indexOf(currentPath) != 0 ) {
		continue
	}
	
	var exists = false
	
	for( var i = 0; i < existingPaths.length; i++ ) {
		var existingPath = existingPaths[i]
		if( storedPath.indexOf(existingPath) == 0 ) {
			exists = true
			break
		}
	}
	
	if( !exists ) {
		localStorage.removeItem(key)
	}
	
}

httpfsx.addEventListener('click', function(event) {
	
	if( event.target.classList.contains('star') ) {
		
		var star = event.target
		
		var starKey = star.getAttribute('data-httpfsx-star-key')
		
		if( localStorage.getItem(starKey) ) {
			localStorage.removeItem(starKey)
			star.classList.remove('active')
		} else {
			localStorage.setItem(starKey, 'T')
			star.classList.add('active')
		}
		
	} else if( event.target.classList.contains('clear-storage') ) {
		if( confirm('Clear storage?') ) {
			localStorage.clear()
			var stars = httpfsx.querySelectorAll('.star')
			for( var i = 0; i < stars.length; i++ ) {
				var star = stars[i]
				star.classList.remove('active')
			}
			alert('Storage cleared')
		}
	}
	
})

`
