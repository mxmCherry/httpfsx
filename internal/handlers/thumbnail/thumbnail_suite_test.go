package thumbnail_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestThumbnail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Thumbnail Suite")
}
