package temp_test

import (
	. "github.com/cryptojuice/go-temp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temp", func() {
	Describe("Create", func() {
		It("", func() {
			f, err := Create()
			Expect(err).To(BeNil())
			Expect(f.Name()).To(Equal("temp123"))
		})
	})
})
