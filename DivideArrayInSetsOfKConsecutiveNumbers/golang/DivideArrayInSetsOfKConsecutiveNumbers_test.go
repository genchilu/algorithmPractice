package DivideArrayInSetsOfKConsecutiveNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossibleDivide(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		expect bool
	}{
		{
			[]int{1, 2, 3, 3, 4, 4, 5, 6},
			4,
			true,
		},
		{
			[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11},
			3,
			true,
		},
		{
			[]int{3, 3, 2, 2, 1, 1},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 4},
			3,
			false,
		},
		{
			[]int{1, 2, 3, 3, 4, 4, 5, 6},
			4,
			true,
		},
		{
			[]int{16, 21, 26, 35},
			4,
			false,
		},
	}

	for _, testCase := range testCases {
		actual := isPossibleDivide(testCase.nums, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
