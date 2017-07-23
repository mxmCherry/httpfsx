/*

Command httpfsx launches mobile-friendly HTTP file-system explorer (readonly)

Basic usage:
	httpfsx --addr=:1024 --root=$HOME/share

*/
package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mxmCherry/httpfsx/internal/handlers/static"
	"github.com/mxmCherry/httpfsx/internal/handlers/thumbnail"
	"github.com/mxmCherry/httpfsx/internal/handlers/ui"
)

var flags struct {
	addr string
	root string
}

func init() {
	flag.StringVar(&flags.addr, "addr", ":1024", "listen addr")
	flag.StringVar(&flags.root, "root", ":1024", "public dir")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	if !filepath.IsAbs(flags.root) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		flags.root = filepath.Join(wd, flags.root)
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", static.New()))
	mux.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir(flags.root))))
	mux.Handle("/thumb/", http.StripPrefix("/thumb", thumbnail.New(flags.root)))

	mux.Handle("/index/", ui.New(flags.root, ui.Config{
		MountPath:  "/index/",
		RawPath:    "/files/",
		StaticPath: "/static/",
	}))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index/", http.StatusTemporaryRedirect)
	}))

	return http.ListenAndServe(flags.addr, mux)
}
