package topKFrequentElements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//topKFrequent(nums []int, k int) []int {

func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		expect []int
	}{
		// {
		// 	[]int{1, 1, 1, 2, 2, 3},
		// 	2,
		// 	[]int{1, 2},
		// },
		// {
		// 	[]int{1},
		// 	1,
		// 	[]int{1},
		// },
		{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			10,
			[]int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
		},
	}

	for _, testCase := range testCases {
		actual := topKFrequent(testCase.nums, testCase.k)

		assert.Equal(t, len(testCase.expect), len(actual))
		for _, v := range actual {
			assert.Contains(t, testCase.expect, v)
		}
		//assert.Equal(t, testCase.expect, actual)
	}
}
