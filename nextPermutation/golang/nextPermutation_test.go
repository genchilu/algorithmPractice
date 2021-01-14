package nextPermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func nextPermutation(nums []int) {
func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 5},
			[]int{1, 5, 1},
		},
		{
			[]int{1},
			[]int{1},
		},
	}

	for _, testCase := range testCases {
		nextPermutation(testCase.nums)
		assert.Equal(t, testCase.expect, testCase.nums)
	}
}
