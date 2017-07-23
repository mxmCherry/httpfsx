package thumbnail

import (
	"fmt"
	"io"
	"path"

	"github.com/mxmCherry/httpfsx/internal/mime"
	"github.com/mxmCherry/httpfsx/internal/thumbnail/image"
	"github.com/mxmCherry/httpfsx/internal/video"
)

type Options struct {
	MaxWidth, MaxHeight uint
	ImageQuality        float64
	VideoOffset         float64
}

func Thumbnail(w io.Writer, file string, opt *Options) error {
	mime, err := mime.Detect(file)
	if err != nil {
		return err
	}

	switch path.Dir(mime) {
	case "image":
		return image.Thumbnail(w, opt.MaxWidth, opt.MaxHeight, int(opt.ImageQuality*100), file)
	case "video":
		return video.Thumbnail(w, file, &video.ThumbnailOptions{
			MaxWidth:  opt.MaxWidth,
			MaxHeight: opt.MaxHeight,
			Offset:    opt.VideoOffset,
		})
	}
	// TODO: make this configurable - don't fail (render empty image?) for unsupported types.
	return fmt.Errorf("thumbnail: unsupported file %s type: %s", file, mime)
}
