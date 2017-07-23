package video_test

import (
	"time"

	. "github.com/mxmCherry/httpfsx/internal/video"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseDuration", func() {
	It("should parse duration", func() {
		Expect(ParseDuration("1:02:03.456")).To(Equal(
			1*time.Hour +
				2*time.Minute +
				3*time.Second +
				456*time.Millisecond,
		))
	})
})

var _ = Describe("FormatDuration", func() {
	It("should format duration", func() {
		d := 1*time.Hour +
			2*time.Minute +
			3*time.Second +
			456*time.Millisecond

		Expect(FormatDuration(d)).To(Equal("01:02:03.456"))
	})
})
