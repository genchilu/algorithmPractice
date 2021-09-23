package maxChunksToSorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		arr    []int
		expect int
	}{
		{
			[]int{4, 3, 2, 1, 0},
			1,
		},
		{
			[]int{1, 0, 2, 3, 4},
			4,
		},
		{
			[]int{5, 4, 3, 9, 1, 2},
			1,
		},
		{
			[]int{8, 7, 6, 12, 11, 10, 13, 5, 4},
			1,
		},
		{
			[]int{1, 4, 3, 6, 0, 7, 8, 2, 5},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := maxChunksToSorted(testCase.arr)
		assert.Equal(t, testCase.expect, actual)
	}
}
