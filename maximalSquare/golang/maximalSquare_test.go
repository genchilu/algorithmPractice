package maximalSquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {
	testCases := []struct {
		matrix [][]byte
		expect int
	}{
		{
			[][]byte{
				[]byte{'1', '0', '1', '0', '0'},
				[]byte{'1', '0', '1', '1', '1'},
				[]byte{'1', '1', '1', '1', '1'},
				[]byte{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				[]byte{'0', '0', '1', '0'},
				[]byte{'1', '1', '1', '1'},
				[]byte{'1', '1', '1', '1'},
				[]byte{'1', '1', '1', '0'},
				[]byte{'1', '1', '0', '0'},
				[]byte{'1', '1', '1', '1'},
				[]byte{'1', '1', '1', '0'},
			},
			9,
		},
	}
	// 0 0 1 0
	// 4 3 2 1
	// 4 3 2 1
	// 3 2 1 0
	// 2 1 0 0
	// 4 3 2 1
	// 3 2 1 0

	for _, testCase := range testCases {
		actual := maximalSquare(testCase.matrix)
		assert.Equal(t, testCase.expect, actual)
	}
}
