package surroundedRegions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func solve(board [][]byte) {

func TestSolve(t *testing.T) {
	testCases := []struct {
		board  [][]byte
		expect [][]byte
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			[][]byte{
				{'X'},
			},
			[][]byte{
				{'X'},
			},
		},
	}

	for _, testCase := range testCases {
		solve(testCase.board)
		assert.Equal(t, testCase.expect, testCase.board)
	}
}
