package gameOfLife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//gameOfLife(board [][]int) {
func TestDev(t *testing.T) {
	testCases := []struct {
		board  [][]int
		expect [][]int
	}{
		{
			[][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			[][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
		{
			[][]int{{1, 1}, {1, 0}},
			[][]int{{1, 1}, {1, 1}},
		},
	}

	for _, testCase := range testCases {
		gameOfLife(testCase.board)
		assert.Equal(t, testCase.expect, testCase.board)
	}
}
