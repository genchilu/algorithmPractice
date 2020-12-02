package maxAreaOfIsland

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//maxAreaOfIsland(grid [][]int) int {

func TestMaxAreaOfIsland(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expect int
	}{
		{
			[][]int{
				[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
		{
			[][]int{
				[]int{0, 0, 0, 0, 0, 0, 0, 0},
			},
			0,
		},
		{
			[][]int{
				[]int{0, 1, 1, 1, 1, 1, 1, 1},
				[]int{0, 0, 0, 0, 1, 0, 0, 1},
				[]int{1, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 1, 1, 0, 1, 0},
				[]int{0, 0, 0, 1, 0, 0, 0, 1},
				[]int{0, 1, 1, 1, 0, 0, 0, 1},
				[]int{0, 0, 1, 0, 1, 1, 0, 1},
				[]int{0, 1, 0, 1, 0, 1, 1, 1},
				[]int{1, 0, 1, 0, 1, 1, 1, 0},
				[]int{0, 1, 1, 0, 1, 0, 0, 0},
			},
			12,
		},
	}

	for _, testCase := range testCases {
		actual := maxAreaOfIsland(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}
}
