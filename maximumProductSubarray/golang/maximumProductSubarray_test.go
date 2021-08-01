package maximumProductSubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxProduct(nums []int) int
func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{-100},
			-100,
		},
		{
			[]int{2, 3},
			6,
		},
		{
			[]int{2, 3, 0, 1, 8},
			8,
		},
		{
			[]int{2, 3, 0, 2, 8, 8, 0, 5, 6, 3},
			128,
		},
		{
			[]int{2, 3, 0, 2, 8, 8, 0, 5, 6, 3},
			128,
		},
		{
			[]int{2, 3, 0, 2, -1, 8, 8, 0, 5, 6, 3},
			90,
		},
		{
			[]int{2, 3, 0, 2, -1, 8, 8, -1, 0, 5, 6, 3},
			128,
		},
		{
			[]int{0, 0, 0},
			0,
		},
		{
			[]int{-1, -2, -3},
			6,
		},
		{
			[]int{-2, 0, -1},
			0,
		},
		{
			[]int{3, -1, 4},
			4,
		},
	}

	for _, testCase := range testCases {
		actual := maxProduct(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
