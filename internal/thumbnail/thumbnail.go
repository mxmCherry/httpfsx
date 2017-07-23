package thumbnail

import (
	"fmt"
	"io"
	"path"
	"time"

	"github.com/mxmCherry/httpfsx/internal/mime"
	"github.com/mxmCherry/httpfsx/internal/thumbnail/image"
	"github.com/mxmCherry/httpfsx/internal/thumbnail/video"
)

func Thumbnail(w io.Writer, width, height uint, jpegQuality int, offset time.Duration, filename string) error {
	mime, err := mime.Detect(filename)
	if err != nil {
		return err
	}

	switch path.Dir(mime) {
	case "image":
		return image.Thumbnail(w, width, height, jpegQuality, filename)
	case "video":
		return video.Thumbnail(w, width, height, jpegQuality, offset, filename)
	}
	// TODO: make this configurable - don't fail (render empty image?) for unsupported types.
	return fmt.Errorf("thumbnail: unsupported file %s type: %s", filename, mime)
}
