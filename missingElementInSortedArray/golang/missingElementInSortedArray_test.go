package missingElementInSortedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//missingElement(nums []int, k int) int {
func TestMissingElement(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		expect int
	}{
		{
			[]int{4, 7, 9, 10},
			1,
			5,
		},
		{
			[]int{4, 7, 9, 10},
			3,
			8,
		},
		{
			[]int{1, 2, 4},
			3,
			6,
		},
	}

	for _, testCase := range testCases {
		actual := missingElement(testCase.nums, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
