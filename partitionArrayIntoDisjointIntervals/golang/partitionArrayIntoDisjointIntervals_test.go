package partitionArrayIntoDisjointIntervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{5, 0, 3, 8, 6},
			3,
		},
		{
			[]int{1, 1, 1, 0, 6, 12},
			4,
		},
	}

	for _, testCase := range testCases {
		actual := partitionDisjoint(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
