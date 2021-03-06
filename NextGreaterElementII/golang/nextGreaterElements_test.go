package nextGreaterElements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nextGreaterElements(nums []int) []int {
func TestNextGreaterElements(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{3, 8, 4, 1, 2},
			[]int{8, -1, 8, 2, 3},
		},
		{
			[]int{1, 2, 1},
			[]int{2, -1, 2},
		},
		{
			[]int{1, 2, 3, 4, 3},
			[]int{2, 3, 4, -1, 4},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{-1, 5, 5, 5, 5},
		},
	}

	for _, teseCase := range testCases {
		actual := nextGreaterElements(teseCase.nums)
		assert.Equal(t, teseCase.expect, actual)
	}
}
