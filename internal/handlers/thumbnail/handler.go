package thumbnail

import (
	"net/http"
	"os"

	"github.com/gorilla/schema"
	"github.com/mxmCherry/httpfsx/internal/filesystem"
	"github.com/mxmCherry/httpfsx/internal/httptool"
	"github.com/mxmCherry/httpfsx/internal/thumbnail"
)

var schemaDecoder = schema.NewDecoder()

type options struct {
	MaxWidth     uint    `schema:"w"`
	MaxHeight    uint    `schema:"h"`
	ImageQuality float64 `schema:"q"`
	VideoOffset  float64 `schema:"offset"`
}

func New(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		abs := filesystem.Abs(root, r.URL.Path)

		var opt options
		if err := schemaDecoder.Decode(&opt, r.URL.Query()); err != nil {
			sendError(w, err, http.StatusBadRequest)
			return
		}
		if opt.ImageQuality == 0 {
			opt.ImageQuality = 0.75
		}

		stats, err := os.Lstat(abs)
		if err != nil {
			sendError(w, err, http.StatusInternalServerError)
			return
		}

		if httptool.NotModified(w, r, stats.ModTime()) {
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")
		err = thumbnail.Thumbnail(r.Context(), w, abs, &thumbnail.Options{
			MaxWidth:     opt.MaxWidth,
			MaxHeight:    opt.MaxHeight,
			ImageQuality: opt.ImageQuality,
			VideoOffset:  opt.VideoOffset,
		})
		if err != nil {
			sendError(w, err, http.StatusInternalServerError)
			return
		}
	})
}

func sendError(w http.ResponseWriter, err error, code int) {
	switch {
	case os.IsNotExist(err):
		code = http.StatusNotFound
	case os.IsPermission(err):
		code = http.StatusForbidden
	case code == 0:
		code = http.StatusInternalServerError
	}
	http.Error(w, err.Error(), code)
}
