package video_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVideo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Video Suite")
}
