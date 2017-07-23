package image_test

import (
	"bytes"
	"image"

	. "github.com/mxmCherry/httpfsx/internal/image"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Thumbnail", func() {

	It("should generate thumbnail", func() {
		buf := bytes.NewBuffer(nil)

		err := Thumbnail(buf, "testdata/lenna.png", &ThumbnailOptions{
			MaxWidth:  100,
			MaxHeight: 100,
			Quality:   0.50,
		})
		Expect(err).NotTo(HaveOccurred())

		img, format, err := image.Decode(buf)
		Expect(err).NotTo(HaveOccurred())
		Expect(format).To(Equal("jpeg"))

		size := img.Bounds().Max
		Expect(size.X).To(BeNumerically("<=", 100))
		Expect(size.Y).To(BeNumerically("<=", 100))
	})

})
