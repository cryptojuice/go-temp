package temp_test

import (
	. "github.com/cryptojuice/go-temp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temp", func() {
	Describe("Create", func() {
		It("", func() {
			file := new(File)
			file.SetDir("test1/")
			_, err := file.Create()
			Expect(err).To(BeNil())
			file.CleanUp()
		})
	})
})
