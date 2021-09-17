package splitArrayLargestSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		m      int
		expect int
	}{
		{
			[]int{7, 2, 5, 10, 8},
			2,
			18,
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			9,
		},
		{
			[]int{1, 4, 4},
			3,
			4,
		},
		{
			[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
			8,
			25,
		},
	}

	for _, testCase := range testCases {
		actual := splitArray(testCase.nums, testCase.m)
		assert.Equal(t, testCase.expect, actual)
	}
}
