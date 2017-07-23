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

	"github.com/mxmCherry/httpfsx/internal/filesystem"
	"github.com/mxmCherry/httpfsx/internal/handlers/thumbnail"
	"github.com/mxmCherry/httpfsx/internal/rawhandler"
	"github.com/mxmCherry/httpfsx/internal/statichandler"
	"github.com/mxmCherry/httpfsx/internal/uihandler"
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

	fs := filesystem.New(flags.root)

	mux := http.NewServeMux()

	mux.Handle("/fs/static/", http.StripPrefix("/fs/static/", statichandler.New()))
	mux.Handle("/fs/raw/", http.StripPrefix("/fs/raw", rawhandler.New(fs)))
	mux.Handle("/thumb/", http.StripPrefix("/thumb", thumbnail.New(flags.root)))

	mux.Handle("/fs/explore/", uihandler.New(fs, uihandler.Config{
		MountPath:  "/fs/explore/",
		RawPath:    "/fs/raw/",
		StaticPath: "/fs/static/",
	}))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/fs/explore/", http.StatusTemporaryRedirect)
	}))

	return http.ListenAndServe(flags.addr, mux)
}
