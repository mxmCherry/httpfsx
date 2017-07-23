package image

import (
	"image"
	"image/jpeg"
	"io"
	"os"

	"github.com/nfnt/resize"

	_ "image/gif"
	_ "image/png"
)

type ThumbnailOptions struct {
	MaxWidth, MaxHeight uint
	Quality             float64
}

func Thumbnail(w io.Writer, filename string, opt *ThumbnailOptions) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	_ = file.Close()

	maxWidth, maxHeight := opt.MaxWidth, opt.MaxHeight
	size := img.Bounds().Max
	if maxWidth == 0 {
		maxWidth = uint(size.X)
	}
	if maxHeight == 0 {
		maxHeight = uint(size.Y)
	}

	thumb := resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)
	return jpeg.Encode(w, thumb, &jpeg.Options{
		Quality: int(opt.Quality * 100),
	})
}
