package longestIncreasingSubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func lengthOfLIS(nums []int) int {
func TestDev(t *testing.T) {
	testCsses := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
	}

	for _, testCase := range testCsses {
		actual := lengthOfLIS(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
