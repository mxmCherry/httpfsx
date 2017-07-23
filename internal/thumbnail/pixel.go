package thumbnail

import (
	"bytes"
	"image"
	"image/jpeg"
)

var pixel []byte

func init() {
	buf := bytes.NewBuffer(nil)
	err := jpeg.Encode(buf, image.NewAlpha(image.Rect(0, 0, 1, 1)), &jpeg.Options{
		Quality: 0,
	})
	if err != nil {
		panic("thumbnail: init pixel failed: " + err.Error()) // impossible
	}
	pixel = buf.Bytes()
}
