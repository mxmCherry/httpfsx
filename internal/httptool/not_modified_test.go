package httptool_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/mxmCherry/httpfsx/internal/httptool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NotModified", func() {
	var res *httptest.ResponseRecorder
	var req *http.Request

	before := time.Now()
	after := before.Add(time.Second)

	BeforeEach(func() {
		res = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/", nil)
	})

	It("should handle malformed If-Modified-Since headers", func() {
		req.Header.Set("If-Modified-Since", "INVALID")

		Expect(NotModified(res, req, time.Time{})).To(BeTrue())
		Expect(res.Code).To(Equal(http.StatusBadRequest))
		Expect(res.Header()).NotTo(HaveKey("Last-Modified"))
	})

	It("should return Not Modified if last modification time is before requested in If-Modified-Since", func() {
		req.Header.Set("If-Modified-Since", after.Format(http.TimeFormat))

		Expect(NotModified(res, req, before)).To(BeTrue())
		Expect(res.Code).To(Equal(http.StatusNotModified))
		Expect(res.Header()).NotTo(HaveKey("Last-Modified"))
		Expect(res.Body.Len()).To(BeZero())
	})

	It("should not serve request, but set Last-Modified if last modification time is after requested in If-Modified-Since", func() {
		req.Header.Set("If-Modified-Since", before.Format(http.TimeFormat))

		Expect(NotModified(res, req, after)).To(BeFalse())
		Expect(res.Header().Get("Last-Modified")).To(Equal(after.Format(http.TimeFormat)))
	})
})
