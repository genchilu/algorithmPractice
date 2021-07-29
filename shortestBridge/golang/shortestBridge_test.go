package shortestBridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func shortestBridge(grid [][]int) int

func TestShortestBridge(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expcet int
	}{
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			1,
		},
		{
			[][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			2,
		},
		{
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			1,
		},
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := shortestBridge(testCase.grid)
		assert.Equal(t, testCase.expcet, actual)
	}
}
