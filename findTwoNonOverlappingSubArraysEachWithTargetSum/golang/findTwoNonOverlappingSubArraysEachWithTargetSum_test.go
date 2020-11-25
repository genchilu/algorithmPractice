package findTwoNonOverlappingSubArraysEachWithTargetSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSumOfLengths(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		expect int
	}{
		{
			[]int{3, 2, 2, 4, 3},
			3,
			2,
		},
		{
			[]int{7, 3, 4, 7},
			7,
			2,
		},
		{
			[]int{4, 3, 2, 6, 2, 3, 4},
			6,
			-1,
		},
		{
			[]int{5, 5, 4, 4, 5},
			3,
			-1,
		},
		{
			[]int{3, 1, 1, 1, 5, 1, 2, 1},
			3,
			3,
		},
		{
			[]int{1, 6, 1},
			7,
			-1,
		},
		{
			[]int{1, 1, 1, 2, 2, 2, 4, 4},
			6,
			6,
		},
	}

	for _, testCase := range testCases {
		actual := minSumOfLengths(testCase.arr, testCase.target)

		assert.Equal(t, testCase.expect, actual)
	}
}
