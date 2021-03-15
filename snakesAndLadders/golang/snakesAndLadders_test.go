package snakesAndLadders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func snakesAndLadders(board [][]int) int {

func TestSnakesAndLadders(t *testing.T) {
	testCases := []struct {
		board  [][]int
		expect int
	}{
		{
			[][]int{
				[]int{-1, -1, -1, -1, -1, -1},
				[]int{-1, -1, -1, -1, -1, -1},
				[]int{-1, -1, -1, -1, -1, -1},
				[]int{-1, 35, -1, -1, 13, -1},
				[]int{-1, -1, -1, -1, -1, -1},
				[]int{-1, 15, -1, -1, -1, -1},
			},
			4,
		},
	}

	for _, testCase := range testCases {
		actual := snakesAndLadders(testCase.board)
		assert.Equal(t, testCase.expect, actual)
	}
}
