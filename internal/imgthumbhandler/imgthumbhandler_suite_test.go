package imgthumbhandler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImgthumbhandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imgthumbhandler Suite")
}
