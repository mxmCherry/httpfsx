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

func Thumbnail(w io.Writer, width, height uint, jpegQuality int, filename string) error {
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

	thumb := resize.Thumbnail(width, height, img, resize.Lanczos3)
	return jpeg.Encode(w, thumb, &jpeg.Options{
		Quality: jpegQuality,
	})
}
