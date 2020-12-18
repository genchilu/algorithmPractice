package minimumEffortPath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumEffortPath(t *testing.T) {
	testCases := []struct {
		heights [][]int
		expect  int
	}{
		{
			[][]int{
				[]int{1, 2, 1, 1, 1},
				[]int{1, 2, 1, 2, 1},
				[]int{1, 2, 1, 2, 1},
				[]int{1, 2, 1, 2, 1},
				[]int{1, 1, 1, 2, 1},
			},
			0,
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{3, 8, 4},
				[]int{5, 3, 5},
			},
			1,
		},
		{
			[][]int{
				[]int{1, 2, 2},
				[]int{3, 8, 2},
				[]int{5, 3, 5},
			},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := minimumEffortPath(testCase.heights)
		assert.Equal(t, testCase.expect, actual)
	}
}
