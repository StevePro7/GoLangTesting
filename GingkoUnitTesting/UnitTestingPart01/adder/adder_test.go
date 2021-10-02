package adder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"unittesting/adder"
)

var _ = Describe("Adder", func() {

	Describe("Add", func() {

		It("adds two numbers", func() {

			sum := adder.Add(2, 3)
			Expect(sum).To(Equal(5))
		})
	})

})
