package singleElementInASortedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//singleNonDuplicate(nums []int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			2,
		},
		{
			[]int{3, 3, 7, 7, 10, 11, 11},
			10,
		},
		{
			[]int{1, 2, 2, 3, 3},
			1,
		},
	}

	for _, testCaase := range testCases {
		actual := singleNonDuplicate(testCaase.nums)
		assert.Equal(t, testCaase.expect, actual)
	}
}
