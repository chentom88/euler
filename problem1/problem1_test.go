package problem1_test

import (
	"github.com/chentom88/euler/problem1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo/extensions/table"
)

var _ = Describe("Problem1", func() {
	DescribeTable("Should sum all multiples of 3 and 5 less than given number", func(max, expectedResult int, throwsError bool) {
		 sum, err := problem1.Solve(max)

		if throwsError {
			Expect(err).To(HaveOccurred())
			return
		}

		Expect(sum).To(Equal(expectedResult))
	},
	Entry("Finds result for 10", 10, 23, false),
	Entry("Finds result for 20", 20, 78, false),
	Entry( "Find result for zero", 0, 0, false),
	Entry( "Find result for negative 1", -1, 0, true))
})
