package numberOfIslands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {

	testCases := []struct {
		grids  [][]byte
		expect int
	}{
		{
			[][]byte{
				[]byte{'1', '1', '1', '1', '0'},
				[]byte{'1', '1', '0', '1', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '1', '0', '0'},
				[]byte{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := numIslands(testCase.grids)
		assert.Equal(t, testCase.expect, actual)
	}
}
