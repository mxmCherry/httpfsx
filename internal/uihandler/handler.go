package uihandler

import (
	"net/http"
	"os"

	"github.com/mxmCherry/httpfsx/internal/filesystem"
)

type Handler struct {
	fs *filesystem.FS
}

func New(fs *filesystem.FS) *Handler {
	return &Handler{
		fs: fs,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type templateData struct {
		List *filesystem.List
	}

	list, err := h.fs.List(r.URL.Path)
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

	err = tmpl.Execute(w, &templateData{
		List: list,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
