package rawhandler

import (
	"net/http"

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
	http.ServeFile(w, r, h.fs.Abs(r.URL.Path))
}
