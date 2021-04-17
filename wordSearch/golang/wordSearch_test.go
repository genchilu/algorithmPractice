package wordSearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func exist(board [][]byte, word string) bool {
func TestExist(t *testing.T) {
	testCases := []struct {
		board  [][]byte
		word   string
		expect bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
			false,
		},
	}

	for _, testCase := range testCases {
		actual := exist(testCase.board, testCase.word)
		assert.Equal(t, testCase.expect, actual)
	}
}
