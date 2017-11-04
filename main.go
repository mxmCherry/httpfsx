package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mxmCherry/httpfsx/internal/httpfsx"
)

var flags struct {
	addr string
	root string
}

func init() {
	flag.StringVar(&flags.addr, "addr", ":1024", "listen addr")
	flag.StringVar(&flags.root, "root", ".", "public root dir")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err.Error())
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

	log.Println("listening on", flags.addr)
	return http.ListenAndServe(flags.addr, httpfsx.FileServer(http.Dir(flags.root)))
}
