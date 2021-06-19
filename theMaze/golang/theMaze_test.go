package theMaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPath(t *testing.T) {
	testCases := []struct {
		maze        [][]int
		start       []int
		destination []int
		expect      bool
	}{
		{
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			[]int{0, 4},
			[]int{4, 4},
			true,
		},
		{
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			[]int{0, 4},
			[]int{3, 2},
			false,
		},
		{
			[][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 0, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1},
				{0, 1, 0, 0, 0},
			},
			[]int{4, 3},
			[]int{0, 1},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := hasPath(testCase.maze, testCase.start, testCase.destination)
		assert.Equal(t, testCase.expect, actual)
	}
}
