package thumbnail

import (
	"fmt"
	"io"
	"path"

	"github.com/mxmCherry/httpfsx/internal/image"
	"github.com/mxmCherry/httpfsx/internal/mime"
	"github.com/mxmCherry/httpfsx/internal/video"
)

type Options struct {
	MaxWidth, MaxHeight uint
	ImageQuality        float64
	VideoOffset         float64
	Strict              bool
}

func Thumbnail(w io.Writer, file string, opt *Options) error {
	mime, err := mime.Detect(file)
	if err != nil {
		return err
	}

	switch path.Dir(mime) {

	case "image":
		return image.Thumbnail(w, file, &image.ThumbnailOptions{
			MaxWidth:  opt.MaxWidth,
			MaxHeight: opt.MaxHeight,
			Quality:   opt.ImageQuality,
		})

	case "video":
		if !opt.Strict && !video.Supported() {
			_, err := w.Write(pixel)
			return err
		}
		return video.Thumbnail(w, file, &video.ThumbnailOptions{
			MaxWidth:  opt.MaxWidth,
			MaxHeight: opt.MaxHeight,
			Offset:    opt.VideoOffset,
		})

	}

	if opt.Strict {
		return fmt.Errorf("thumbnail: unsupported file %s type: %s", file, mime)
	}
	_, err = w.Write(pixel)
	return err
}
