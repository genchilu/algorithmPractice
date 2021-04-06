package sortColors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//sortColors
func TestSortColors(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 0},
			[]int{0, 1, 2},
		},
	}

	for _, testCase := range testCases {
		sortColors(testCase.nums)
		assert.Equal(t, testCase.expect, testCase.nums)
	}
}
