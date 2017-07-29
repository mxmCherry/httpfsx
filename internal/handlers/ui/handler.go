package ui

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/mxmCherry/httpfsx/internal/filesystem"
	"github.com/mxmCherry/httpfsx/internal/httptool"
)

type Handler struct {
	config Config
	root   string
	tmpl   *template.Template
}

type Config struct {
	MountPath  string
	RawPath    string
	StaticPath string
}

func New(root string, config Config) *Handler {
	config.MountPath = path.Clean(config.MountPath)
	config.RawPath = path.Clean(config.RawPath)
	config.StaticPath = path.Clean(config.StaticPath)

	tmpl := template.New("uihandler").Funcs(template.FuncMap{
		"staticLink": func(staticPath string) string {
			return path.Join(config.StaticPath, staticPath)
		},
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
				return path.Join(config.MountPath, x.Path)
			case filesystem.File:
				return path.Join(config.RawPath, x.Path)
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
		"metaKind": func(v interface{}) string {
			switch x := v.(type) {
			case filesystem.File:
				return path.Dir(x.Mime)
			default:
				return ""
			}
		},
	})

	return &Handler{
		config: config,
		root:   root,
		tmpl:   template.Must(tmpl.Parse(tmplCode)),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type templateData struct {
		List *filesystem.List
	}

	list, err := filesystem.Ls(h.root, strings.TrimPrefix(r.URL.Path, h.config.MountPath))
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if os.IsPermission(err) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if httptool.NotModified(w, r, list.LastMod) {
		return
	}

	err = h.tmpl.Execute(w, &templateData{
		List: list,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
