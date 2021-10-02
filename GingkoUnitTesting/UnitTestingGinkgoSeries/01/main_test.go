package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "unittesting"
)

func TestUnitTestingGinkgoSeries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingDemo Suite")
}

var _ = Describe("Sum", func() {

	var (
		p, q, m, n, sum1, sum2 int
	)

	BeforeEach(func() {
		p, q, sum1 = 5, 6, 11
		// putting wrong value of sum2 intentionally
		m, n, sum2 = 8, 7, 16
	})

	Context("Addition of two digits", func() {
		It("should return sum of the two digits", func() {
			additionOfTwoDigits := Sum(p, q)
			Expect(additionOfTwoDigits).Should(Equal(sum1))
		})
		It("should not return the sum provided", func() {
			additionOfTwoDigits := Sum(m, n)
			Expect(additionOfTwoDigits).ShouldNot(Equal(sum2))
		})
	})

})
