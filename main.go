package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/mxmCherry/httpfsx/internal/filesystem"
	"github.com/mxmCherry/httpfsx/internal/rawhandler"
	"github.com/mxmCherry/httpfsx/internal/statichandler"
	"github.com/mxmCherry/httpfsx/internal/uihandler"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	addr := flag.String("addr", ":1024", "listen addr")
	root := flag.String("root", wd, "root dir")

	flag.Parse()

	fs := filesystem.New(*root)

	mux := http.NewServeMux()
	mux.Handle("/fs/static/", http.StripPrefix("/fs/static/", statichandler.New()))
	mux.Handle("/fs/raw/", http.StripPrefix("/fs/raw", rawhandler.New(fs)))
	mux.Handle("/fs/explore/", http.StripPrefix("/fs/explore", uihandler.New(fs)))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/fs/explore/", http.StatusTemporaryRedirect)
	}))

	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		panic(err.Error())
	}
}
