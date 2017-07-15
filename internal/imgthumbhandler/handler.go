package imgthumbhandler

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/nfnt/resize"

	_ "image/gif"
	_ "image/png"
)

const (
	jpegMime           = "image/jpeg"
	defaultJpegQuality = 75
)

func New(resolvePath func(requestPath string) string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := resolvePath(path.Join("/", r.URL.Path))

		var maxWidth, maxHeight uint
		if maxWidthStr := r.FormValue("max_width"); maxWidthStr != "" {
			val, err := strconv.ParseUint(maxWidthStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			maxWidth = uint(val)
		}
		if maxHeightStr := r.FormValue("max_height"); maxHeightStr != "" {
			val, err := strconv.ParseUint(maxHeightStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			maxHeight = uint(val)
		}
		if maxWidth == 0 && maxHeight == 0 {
			http.Error(w, "Neither max_width nor max_height provided", http.StatusBadRequest)
			return
		}
		if maxWidth == 0 {
			maxWidth = maxHeight
		} else if maxHeight == 0 {
			maxHeight = maxWidth
		}

		var jpegQuality int
		if jpegQualityStr := r.FormValue("jpeg_quality"); jpegQualityStr != "" {
			val, err := strconv.ParseInt(jpegQualityStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jpegQuality = int(val)
		}
		if jpegQuality == 0 {
			jpegQuality = defaultJpegQuality
		} else if jpegQuality < 0 || jpegQuality > 100 {
			http.Error(w, "jpeg_quality must be positive and less than or equal to 100", http.StatusBadRequest)
			return
		}

		stats, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if ifMod := r.Header.Get("If-Modified-Since"); ifMod != "" {
			t, err := time.Parse(http.TimeFormat, ifMod)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if stats.ModTime().Before(t) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// decode jpeg into image.Image
		img, _, err := image.Decode(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = file.Close()

		thumb := resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)

		w.Header().Set("Content-Type", jpegMime)
		w.Header().Set("Last-Modified", stats.ModTime().Format(http.TimeFormat))
		err = jpeg.Encode(w, thumb, &jpeg.Options{
			Quality: jpegQuality,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
