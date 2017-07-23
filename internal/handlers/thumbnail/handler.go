package thumbnail

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mxmCherry/httpfsx/internal/thumbnail"
)

const (
	jpegMime           = "image/jpeg"
	defaultJpegQuality = 75
)

func New(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		abs := filepath.Join(root, filepath.Join("/", r.URL.Path)) // TODO: extract into filesystem as standalone func

		var maxWidth, maxHeight, jpegQuality, offset uint
		jpegQuality = defaultJpegQuality
		offset = 1
		if !parseUint(w, r, "max_width", &maxWidth) {
			return
		}
		maxHeight = maxWidth
		if !parseUint(w, r, "max_height", &maxHeight) {
			return
		}
		if !parseUint(w, r, "jpeg_quality", &jpegQuality) {
			return
		}
		if !parseUint(w, r, "offset", &offset) {
			return
		}

		stats, err := os.Stat(abs)
		if err != nil {
			sendError(w, err, http.StatusInternalServerError)
			return
		}

		if ifMod := r.Header.Get("If-Modified-Since"); ifMod != "" {
			t, err := time.Parse(http.TimeFormat, ifMod)
			if err != nil {
				sendError(w, err, http.StatusBadRequest)
				return
			}
			if stats.ModTime().Before(t) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		w.Header().Set("Content-Type", jpegMime)
		w.Header().Set("Last-Modified", stats.ModTime().Format(http.TimeFormat))
		if err := thumbnail.Thumbnail(w, maxWidth, maxHeight, int(jpegQuality), time.Duration(offset)*time.Second, abs); err != nil {
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

func parseUint(w http.ResponseWriter, r *http.Request, name string, u *uint) bool {
	str := r.FormValue(name)
	if str == "" {
		if *u != 0 {
			return true
		}
		sendError(w, fmt.Errorf("missing required param '%s'", name), http.StatusBadRequest)
		return false
	}

	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		sendError(w, fmt.Errorf("failed to parse uint '%s' param (value: '%s'): %s", name, str, err.Error()), http.StatusBadRequest)
		return false
	}
	*u = uint(val)
	return true
}
