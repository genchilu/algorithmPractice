package searchInRotatedSortedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func search(nums []int, target int) int {
func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		expect int
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			-1,
		},
		{
			[]int{1},
			1,
			0,
		},
		{
			[]int{5, 1, 3},
			1,
			1,
		},
		{
			[]int{7, 8, 1, 2, 3, 4, 5, 6},
			2,
			3,
		},
		{
			[]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
			1,
			6,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.nums, testCase.target)
		assert.Equal(t, testCase.expect, actual)
	}
}
