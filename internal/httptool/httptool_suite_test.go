package httptool_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHttputil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httptool Suite")
}
