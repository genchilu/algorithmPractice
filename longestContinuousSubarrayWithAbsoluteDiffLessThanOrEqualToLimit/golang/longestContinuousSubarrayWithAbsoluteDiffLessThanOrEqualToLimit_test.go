package longestContinuousSubarrayWithAbsoluteDiffLessThanOrEqualToLimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {

	testCases := []struct {
		nums   []int
		limit  int
		expect int
	}{
		// {
		// 	[]int{8, 2, 4, 7},
		// 	4,
		// 	2,
		// },
		// {
		// 	[]int{10, 1, 2, 4, 7, 2},
		// 	5,
		// 	4,
		// },
		{
			[]int{4, 2, 2, 2, 4, 4, 2, 2},
			0,
			3,
		},
	}

	for _, testCase := range testCases {
		actual := longestSubarray(testCase.nums, testCase.limit)
		assert.Equal(t, testCase.expect, actual)
	}

}
