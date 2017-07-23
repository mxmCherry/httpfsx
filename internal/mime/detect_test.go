package mime_test

import (
	. "github.com/mxmCherry/httpfsx/internal/mime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Detect", func() {
	It("should detect file MIME", func() {
		Expect(Detect("testdata/lenna.png")).To(Equal("image/png"))
	})
})
