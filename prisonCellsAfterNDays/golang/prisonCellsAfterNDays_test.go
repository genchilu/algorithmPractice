package prisonCellsAfterNDays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrisonAfterNDays(t *testing.T) {
	testCases := []struct {
		cells  []int
		N      int
		expect []int
	}{
		// {
		// 	[]int{0, 1, 0, 1, 1, 0, 0, 1},
		// 	7,
		// 	[]int{0, 0, 1, 1, 0, 0, 0, 0},
		// },
		// {
		// 	[]int{1, 0, 0, 1, 0, 0, 1, 0},
		// 	1000000000,
		// 	[]int{0, 0, 1, 1, 1, 1, 1, 0},
		// },
		{
			[]int{1, 0, 0, 0, 1, 0, 0, 1},
			99,
			[]int{0, 0, 1, 0, 1, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := prisonAfterNDays(testCase.cells, testCase.N)
		assert.Equal(t, testCase.expect, actual)
	}
}
