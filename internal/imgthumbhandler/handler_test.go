package imgthumbhandler_test

import (
	"image/jpeg"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"

	. "github.com/mxmCherry/httpfsx/internal/imgthumbhandler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var subject http.Handler
	var req *http.Request
	var resp *httptest.ResponseRecorder

	BeforeEach(func() {
		subject = New(func(requestPath string) string {
			return filepath.Join("testdata", requestPath)
		})
		req = httptest.NewRequest("GET", "/lenna.png?max_width=500&max_height=500", nil)
		resp = httptest.NewRecorder()
	})

	It("should return error if bad max_width given", func() {
		subject.ServeHTTP(resp, httptest.NewRequest("GET", "/lenna.png?max_width=INVALID&max_height=500", nil))
		Expect(resp.Code).To(Equal(http.StatusBadRequest))
	})

	It("should return error if bad max_height given", func() {
		subject.ServeHTTP(resp, httptest.NewRequest("GET", "/lenna.png?max_width=500&max_height=INVALID", nil))
		Expect(resp.Code).To(Equal(http.StatusBadRequest))
	})

	It("should return error if no max dimensions given", func() {
		subject.ServeHTTP(resp, httptest.NewRequest("GET", "/lenna.png", nil))
		Expect(resp.Code).To(Equal(http.StatusBadRequest))
	})

	DescribeTable("return error if invalid jpeg_quality given",
		func(quality string) {
			subject.ServeHTTP(resp, httptest.NewRequest("GET", "/lenna.png?max_width=500&max_height=500&jpeg_quality="+quality, nil))
			Expect(resp.Code).To(Equal(http.StatusBadRequest))
		},
		Entry("non-integer", "INVALID"),
		Entry("negative", "-1"),
		Entry("greater than 100", "101"),
	)

	It("should return error if file does not exist", func() {
		subject.ServeHTTP(resp, httptest.NewRequest("GET", "/non-existing.png?max_width=500&max_height=500", nil))
		Expect(resp.Code).To(Equal(http.StatusNotFound))
	})

	It("should set Content-Type header", func() {
		subject.ServeHTTP(resp, req)
		Expect(resp.Code).To(Equal(http.StatusOK))
		Expect(resp.Header().Get("Content-Type")).To(Equal("image/jpeg"))
	})

	It("should set Last-Modified header", func() {
		subject.ServeHTTP(resp, req)
		Expect(resp.Code).To(Equal(http.StatusOK))

		lastModStr := resp.Header().Get("Last-Modified")
		Expect(lastModStr).NotTo(BeEmpty())
		Expect(lastModStr).To(Equal(fileLastMod("testdata/lenna.png").Format(http.TimeFormat)))
	})

	It("should return JPEG image with width/height up to to specified size", func() {
		subject.ServeHTTP(resp, req)
		Expect(resp.Code).To(Equal(http.StatusOK))

		img, err := jpeg.Decode(resp.Body)
		Expect(err).NotTo(HaveOccurred())

		size := img.Bounds().Max
		Expect(size.X).To(Equal(500))
		Expect(size.Y).To(Equal(500))
	})

	// TODO: would be nice to test if it resizes image proportionally etc.
	// TODO: would be nice to test if jpeg_quality actually applied

})

func fileLastMod(name string) time.Time {
	s, err := os.Stat(name)
	Expect(err).NotTo(HaveOccurred())
	return s.ModTime()
}
