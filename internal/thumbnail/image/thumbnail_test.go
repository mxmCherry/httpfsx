package image_test

import (
	"bytes"
	"image"

	. "github.com/mxmCherry/httpfsx/internal/thumbnail/image"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Thumbnail", func() {

	It("should generate thumbnail", func() {
		buf := bytes.NewBuffer(nil)

		err := Thumbnail(buf, 100, 100, 50, "testdata/lenna.png")
		Expect(err).NotTo(HaveOccurred())

		img, format, err := image.Decode(buf)
		Expect(err).NotTo(HaveOccurred())
		Expect(format).To(Equal("jpeg"))

		size := img.Bounds().Max
		Expect(size.X).To(BeNumerically("<=", 100))
		Expect(size.Y).To(BeNumerically("<=", 100))
	})

})
