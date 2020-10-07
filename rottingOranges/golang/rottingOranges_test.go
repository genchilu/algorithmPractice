package rottingOranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	testCases := []struct {
		grids  [][]int
		expect int
	}{
		{
			[][]int{
				[]int{2, 1, 1},
				[]int{1, 1, 0},
				[]int{0, 1, 1},
			},
			4,
		},
		{
			[][]int{
				[]int{2, 1, 1},
				[]int{0, 1, 1},
				[]int{1, 0, 1},
			},
			-1,
		},
		{
			[][]int{
				[]int{0, 2},
			},
			0,
		},
		{
			[][]int{
				[]int{0},
			},
			0,
		},
		{
			[][]int{
				[]int{2, 2},
				[]int{1, 1},
				[]int{0, 0},
				[]int{2, 0},
			},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := orangesRotting(testCase.grids)
		assert.Equal(t, testCase.expect, actual)
	}
}
