package temp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoTemp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTemp Suite")
}
