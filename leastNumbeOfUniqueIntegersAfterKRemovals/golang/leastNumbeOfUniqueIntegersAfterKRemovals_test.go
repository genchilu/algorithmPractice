package leastNumbeOfUniqueIntegersAfterKRemovals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findLeastNumOfUniqueInts(arr []int, k int) int
func TestFindLeastNumOfUniqueInts(t *testing.T) {
	testCases := []struct {
		arr    []int
		k      int
		expect int
	}{
		{
			[]int{5, 5, 4},
			1,
			1,
		},
		{
			[]int{4, 3, 1, 1, 3, 3, 2},
			3,
			2,
		},
	}

	for _, testCase := range testCases {
		actual := findLeastNumOfUniqueInts(testCase.arr, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
