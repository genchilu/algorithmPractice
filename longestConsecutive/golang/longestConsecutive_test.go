package longestConsecutive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//longestConsecutive(nums []int) int
func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			[]int{0, -1},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := longestConsecutive(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
