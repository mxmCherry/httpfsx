package video

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parseDuration", func() {
	It("should parse duration", func() {
		Expect(parseDuration("1:02:03.456")).To(Equal(
			1*time.Hour +
				2*time.Minute +
				3*time.Second +
				456*time.Millisecond,
		))
	})
})

var _ = Describe("formatDuration", func() {
	It("should format duration", func() {
		d := 1*time.Hour +
			2*time.Minute +
			3*time.Second +
			456*time.Millisecond

		Expect(formatDuration(d)).To(Equal("01:02:03.456"))
	})
})
