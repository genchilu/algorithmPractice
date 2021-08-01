package findPeakElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findPeakElement(nums []int) int {
func TestFindPeakElement(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{1, 2, 3, 1},
			[]int{2},
		},
		{
			[]int{1, 2, 1, 3, 5, 6, 4},
			[]int{1, 5},
		},
	}

	for _, testCase := range testCases {
		actual := findPeakElement(testCase.nums)
		assert.Contains(t, testCase.expect, actual)
	}
}
