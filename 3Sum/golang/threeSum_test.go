package threeSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{0},
			[][]int{},
		},
	}

	for _, testCase := range testCases {
		actual := threeSum(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
