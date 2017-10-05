package problem2_test

import (
	"github.com/chentom88/euler/problem2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Birthday candles", func() {
	DescribeTable("should blow out correct number of candles", func(input []int, expected int){
		candlesBlownOut := problem2.BlowOutCandles(input)

		Expect(candlesBlownOut).To(Equal(expected))
	},
	Entry("Basic", []int {3,2,1,3}, 2),
	Entry("Medium", []int {3, 2, 3, 1, 3, 3}, 4),
	Entry("James", []int{1, 1,1,3, 3, 1, 2, 3}, 3),
	Entry( "Negatives", []int{-3, -2, -3, -2, -2}, 3),
	Entry( "Negatives", []int{-1, -2, -3, -2, -2}, 1),
	Entry("Empty array", []int{}, 0),
	Entry("all the same",[]int{1,1,1,1,1,1,1,1,1}, 9),
	Entry("all the same",[]int{0,0,0,0,0}, 5))

	It("Counts the lengths", func() {
		result := problem2.CountCandles([]int{1, 1,1,3, 3, 1, 2, 3})

		Expect(len(result)).ToNot(Equal(0))
		Expect(result[1]).To(Equal(4))
		Expect(result[2]).To(Equal(1))
		Expect(result[3]).To(Equal(3))
	})
})
