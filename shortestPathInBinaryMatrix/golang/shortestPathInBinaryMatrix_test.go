package shortestPathInBinaryMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expect int
	}{
		{
			[][]int{
				[]int{0, 1},
				[]int{1, 0},
			},
			2,
		},
		{
			[][]int{
				[]int{0, 0, 0},
				[]int{1, 1, 0},
				[]int{1, 1, 0},
			},
			4,
		},
		{
			[][]int{
				[]int{1, 0, 0},
				[]int{1, 1, 0},
				[]int{1, 1, 0},
			},
			-1,
		},
		{
			[][]int{
				[]int{0, 0, 1, 0, 0, 0, 0},
				[]int{0, 1, 0, 0, 0, 0, 1},
				[]int{0, 0, 1, 0, 1, 0, 0},
				[]int{0, 0, 0, 1, 1, 1, 0},
				[]int{1, 0, 0, 1, 1, 0, 0},
				[]int{1, 1, 1, 1, 1, 0, 1},
				[]int{0, 0, 1, 0, 0, 0, 0},
			},
			10,
		},
	}

	for _, testCase := range testCases {
		actual := shortestPathBinaryMatrix(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}
}
